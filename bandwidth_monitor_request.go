package homehub

import (
	"fmt"
	"time"
)

type bandwidthMonitorRequest struct {
	genericRequest
}

func newBandwidthMonitorRequest(authData *authData, xpath string) (req *bandwidthMonitorRequest) {
	authData.requestCount++

	// TODO: Enable dates to be configurable
	now := time.Now()
	date := fmt.Sprintf("%d%02d%02d", now.Year(), now.Month(), now.Day())
	params := &parameters{
		StartDate: date,
		EndDate:   date,
	}

	a := action{
		ID:         0,
		Method:     methodUploadStatistics,
		XPath:      xpath,
		Parameters: params,
	}

	var actions []action
	actions = append(actions, a)
	requestBody := newRequestBody(authData, actions)

	return &bandwidthMonitorRequest{
		genericRequest: genericRequest{
			*requestBody,
			*authData,
		},
	}
}
