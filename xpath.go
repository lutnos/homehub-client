package homehub

const accessControlPortForwardingUID string = "Device/NAT/PortMappings/PortMapping[@uid=#]"
const accessControlPortForwardingPortmappings string = "Device/NAT/PortMappings"
const accessControlPortForwardingUpdateExternalPort string = "Device/NAT/PortMappings/PortMapping[@uid=#]/ExternalPort"
const accessControlPortForwardingUpdateDescription string = "Device/NAT/PortMappings/PortMapping[@uid=#]/Description"
const accessControlPortForwardingUpdateRemoteHost string = "Device/NAT/PortMappings/PortMapping[@uid=#]/RemoteHost"
const accessControlPortForwardingUpdateInternalClient string = "Device/NAT/PortMappings/PortMapping[@uid=#]/InternalClient"
const accessControlPortForwardingUpdateAlias string = "Device/NAT/PortMappings/PortMapping[@uid=#]/Alias"
const accessControlPortForwardingUpdateEnable string = "Device/NAT/PortMappings/PortMapping[@uid=#]/Enable"
const accessControlPortForwardingUpdateService string = "Device/NAT/PortMappings/PortMapping[@uid=#]/Service"
const accessControlPortForwardingUpdateInternalPort string = "Device/NAT/PortMappings/PortMapping[@uid=#]/InternalPort"
const accessControlPortForwardingUpdateProtocol string = "Device/NAT/PortMappings/PortMapping[@uid=#]/Protocol"
const accessControlPortForwardingUpdateExternalPortEndRange string = "Device/NAT/PortMappings/PortMapping[@uid=#]/ExternalPortEndRange"
const accessControlPortForwardingValidationDataSubnet string = "Device/Managers/NetworkData/NetmaskLan"
const accessControlPortForwardingValidationDataTr69 string = "Device/ManagementServer/TR69InternalData/Settings/Port"
const accessControlPortForwardingValidationDataPortClamping string = "Device/Services/DeviceConfig/PortClampingEnable"
const accessControlPortForwardingValidationDataLanIP string = "Device/Managers/NetworkData/IpLan"
const accessControlFirewallSipAlgEnable string = "Device/NAT/SIPALGEnable"
const accessControlFirewallConfig string = "Device/Firewall/Config"
const advancedRemoteManagementIpv4Address string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv4Addresses/IPv4Address[Alias=\"IP_DATA_ADDRESS\"]/IPAddress"
const mymediaSambaIP string = "Device/Managers/NetworkData/IpLan"
const mymediaSambaHost string = "Device/DNS/Client/HostName"
const mainUsbPhysicalMedium string = "Device/Services/StorageServices/StorageService[@uid='1']/PhysicalMediums/PhysicalMedium[SerialNumber=\"#\"]"
const mainUsbPhysicalMediumUID string = "Device/Services/StorageServices/StorageService[@uid='1']/PhysicalMediums/PhysicalMedium[@uid='#']"
const mainWifiLayerInterfaceWifi5 string = "Device/WiFi/Radios/Radio[RADIO5G]"
const mainWifiLayerInterfaceWifi24 string = "Device/WiFi/Radios/Radio[RADIO2G4]"
const mainCacheableUsbLogicalVolume string = "Device/Services/StorageServices/StorageService[@uid='1']/LogicalVolumes"
const mainCacheableHosts string = "Device/Hosts/Hosts"
const mainCacheableHistoryMaxAgeInDays string = "Device/Hosts/HistoryMaxAgeInDays"
const mainCacheableDNS string = "Device/DNS/Client/HostName"
const mainCacheableLocalDomains string = "Device/DNS/Client/LocalDomains"
const mainCacheableGwName string = "Device/DeviceInfo/ProductClass"
const oieEnable string = "Device/Services/OnlineInstall/Enable"
const oieEnableGreyedOut string = "Device/Services/OnlineInstall/EnableGreyedOut"
const oieTimeout string = "Device/Services/OnlineInstall/EngineerModeTimeout"
const loginClearTextPassword string = "Device/UserAccounts/Users/User[Login='#login#']/ClearTextPassword"
const loginPassword string = "Device/UserAccounts/Users/User[Login='#login#']/Password"
const loginPasswordHint string = "Device/UserAccounts/Users/User[Login='#login#']/SecretQuery"
const loginPasswordOverideRPC string = "Device/Managers/Global"
const btAccessControlAddTimeSlot string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[last()]/TimeSlots"
const btAccessControlAddMacAddressList string = "Device/Services/ParentalControl/MACAddressLists"
const btAccessControlAddRule string = "Device/Services/ParentalControl/Rules"
const btAccessControlAddTimeSlotChild string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[@uid='#']/TimeSlots"
const btAccessControlAddTimeSlotBlockAll string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[Alias='BlockAllDevices']/TimeSlots"
const btAccessControlAddTimeSlotList string = "Device/Services/ParentalControl/TimeSlotLists"
const btAccessControlEnable string = "Device/Services/ParentalControl/Enable"
const btAccessControlMacAddressList string = "Device/Services/ParentalControl/MACAddressLists/MACAddressList[@uid='#']"
const btAccessControlParentalControl string = "Device/Services/ParentalControl"
const btAccessControlRule string = "Device/Services/ParentalControl/Rules/Rule[@uid='#']"
const btAccessControlUpdateStartTimeBlockAll string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[Alias='BlockAllDevices']/TimeSlots/TimeSlot[@uid='#']/StartTime"
const btAccessControlUpdateIPAddress string = "Device/Services/ParentalControl/MACAddressLists/MACAddressList[@uid='#']/IPAddress"
const btAccessControlUpdateBlockAllAlways string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[Alias='BlockAllDevices']/Always"
const btAccessControlUpdateMacAddress string = "Device/Services/ParentalControl/MACAddressLists/MACAddressList[@uid='#']/MACAddresses"
const btAccessControlUpdateWeekDaysBlockAll string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[Alias='BlockAllDevices']/TimeSlots/TimeSlot[@uid='#']/WeekDays"
const btAccessControlUpdateWanAccess string = "Device/Services/ParentalControl/Rules/Rule[Alias='BlockAllDevices']/WANAccess"
const btAccessControlUpdateBlockAllDevices string = "Device/Services/ParentalControl/MACAddressLists/MACAddressList/AllDevices"
const btAccessControlUpdateWeekDays string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[@uid='#']/TimeSlots/TimeSlot[@uid='#slotUid#']/WeekDays"
const btAccessControlUpdateEnableBlockAll string = "Device/Services/ParentalControl/Rules/Rule[Alias='BlockAllDevices']/Enable"
const btAccessControlUpdateEnable string = "Device/Services/ParentalControl/Rules/Rule[@uid='#']/Enable"
const btAccessControlUpdateBlockAll string = "Device/Services/ParentalControl/MACAddressLists/MACAddressList[Alias='BlockAllDevices']/AllDevices"
const btAccessControlUpdateEndTimeBlockAll string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[Alias='BlockAllDevices']/TimeSlots/TimeSlot[@uid='#']/EndTime"
const btAccessControlUpdateStartTime string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[@uid='#']/TimeSlots/TimeSlot[@uid='#slotUid#']/StartTime"
const btAccessControlUpdateEndTime string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[@uid='#']/TimeSlots/TimeSlot[@uid='#slotUid#']/EndTime"
const btAccessControlTimeSlotChild string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[@uid='#']/TimeSlots/TimeSlot[@uid='#slotUid#']"
const btAccessControlTimeSlotList string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[@uid='#']"
const btAccessControlRemoveTimeSlotBlockAll string = "Device/Services/ParentalControl/TimeSlotLists/TimeSlotList[Alias='BlockAllDevices']/TimeSlots/TimeSlot[@uid='#']"
const deviceDiscover string = "Device/DeviceDiscovery/DeviceIdentification/DeviceTypes"
const statsWanStats string = "Device/IP/Interfaces"
const statsDataUsageDataSent string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/Stats/BytesSent"
const statsDataUsageDataReceived string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/Stats/BytesReceived"
const ethernetDeviceDevicesList string = "Device/Hosts/Hosts/Host[@uid=\"#\"]"
const internetConnectivityLastPPPConnectionDsl string = "Device/PPP/Interfaces/Interface[Alias='PPP_DSL_DATA']/LastPPPConnections"
const internetConnectivityPppFTTHPassword string = "Device/PPP/Interfaces/Interface[Alias='PPP_FTTH_DATA']/Password"
const internetConnectivityPppFTTHEnable string = "Device/PPP/Interfaces/Interface[Alias='PPP_FTTH_DATA']/Enable"
const internetConnectivityPppFTTHStatus string = "Device/PPP/Interfaces/Interface[Alias='PPP_FTTH_DATA']/Status"
const internetConnectivityPppFTTHUsername string = "Device/PPP/Interfaces/Interface[Alias='PPP_FTTH_DATA']/Username"
const internetConnectivityPppDSLPassword string = "Device/PPP/Interfaces/Interface[Alias='PPP_DSL_DATA']/Password"
const internetConnectivityPppDSLEnable string = "Device/PPP/Interfaces/Interface[Alias='PPP_DSL_DATA']/Enable"
const internetConnectivityPppDSLStatus string = "Device/PPP/Interfaces/Interface[Alias='PPP_DSL_DATA']/Status"
const internetConnectivityPppDSLUsername string = "Device/PPP/Interfaces/Interface[Alias='PPP_DSL_DATA']/Username"
const internetConnectivityOperatingModeEncapsulation string = "Device/ATM/Links/Link[Alias='ATM_DATA']/Encapsulation"
const internetConnectivityOperatingModeDestinationAddress string = "Device/ATM/Links/Link[Alias='ATM_DATA']/DestinationAddress"
const internetConnectivityOperatingModeLinkType string = "Device/ATM/Links/Link[Alias='ATM_DATA']/LinkType"
const internetConnectivityVpnPortClamping string = "Device/Services/DeviceConfig/PortClampingEnable"
const internetConnectivityIpv6LinkLocalAddress string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv6Addresses/IPv6Address[Alias=\"LINK_LOCAL\"]/IPAddress"
const internetConnectivityIpv6StatefullLessAdvManagedFlag string = "Device/RouterAdvertisement/InterfaceSettings/InterfaceSetting[Alias=\"RA_BR_LAN\"]/AdvManagedFlag"
const internetConnectivityIpv6StatefullLess string = "Device/DHCPv6/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/IANAEnable"
const internetConnectivityIpv6FtthInterface string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_FTTH_DATA\"]/Status"
const internetConnectivityIpv6FtthRemoteLocal string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_FTTH_DATA\"]/IPv6CP/RemoteInterfaceIdentifier"
const internetConnectivityIpv6NetworkStatus string = "Device/DHCPv6/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/Status"
const internetConnectivityIpv6UlaEnablePrefix string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Prefixes/IPv6Prefix[Alias=\"ULA_LAN\"]/Enable"
const internetConnectivityIpv6UlaPrefix string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Prefixes/IPv6Prefix[Alias=\"ULA_LAN\"]/Prefix"
const internetConnectivityIpv6LinkLocalAddressLAN string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Addresses/IPv6Address[Alias=\"LINK_LOCAL\"]/IPAddress"
const internetConnectivityIpv6DhcpMode string = "Device/Managers/NetworkData/IPv6Mode"
const internetConnectivityIpv6DslRemoteLocal string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_DSL_DATA\"]/IPv6CP/RemoteInterfaceIdentifier"
const internetConnectivityIpv6PreferedIPDATAAddr string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv6Addresses/IPv6Address[IPAddressStatus=\"PREFERRED\"]"
const internetConnectivityIpv6PreferedIPBRLANAddr string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Addresses/IPv6Address[IPAddressStatus=\"PREFERRED\"]"
const internetConnectivityIpv6LocalDomains string = "Device/DNS/Relay/LocalDomains"
const internetConnectivityIpv6HubStatus string = "Device/IP/IPv6Status"
const internetConnectivityIpv6StatefullLessAdvOtherConfigFlag string = "Device/RouterAdvertisement/InterfaceSettings/InterfaceSetting[Alias=\"RA_BR_LAN\"]/AdvOtherConfigFlag"
const internetConnectivityIpv6Hostname string = "Device/DNS/Client/HostName"
const internetConnectivityIpv6UlaEnableAddress string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Addresses/IPv6Address[Alias=\"ULA_LAN\"]/Enable"
const internetConnectivityIpv6DslInterface string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_DSL_DATA\"]/Status"
const internetConnectivityIpv6PreferedIPBRLANPrefix string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Prefixes"
const internetConnectivityIpv6UlaAddress string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Addresses/IPv6Address[Alias=\"ULA_LAN\"]/IPAddress"
const internetConnectivityIpv6PreferedIPDATAPrefix string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv6Prefixes"
const internetConnectivityIpv6DnsV6 string = "Device/DNS/Relay/Forwardings/Forwarding[Status=\"ENABLED\"]/DNSServer"
const internetConnectivityLastPPPConnectionFtth string = "Device/PPP/Interfaces/Interface[Alias='PPP_FTTH_DATA']/LastPPPConnections"
const internetConnectivityLastDSLSynchronization string = "Device/DSL/Lines/Line[@uid='1']/LastDSLSynchronizations"
const eventLog string = "Device/DeviceInfo/VendorLogFiles/VendorLogFile[@uid='1']"
const mySagemcomBoxDynDNSPassword string = "Device/Services/DynamicDNS/Clients/Client[@uid=\"1\"]/Password"
const mySagemcomBoxDynDNSHostname string = "Device/Services/DynamicDNS/Clients/Client[@uid=\"1\"]/Hostnames/Hostname[@uid=\"1\"]/Name"
const mySagemcomBoxDynDNSServiceQuery string = "Device/Services/DynamicDNS/Services/Service[Name='#']"
const mySagemcomBoxDynDNSServiceReference string = "Device/Services/DynamicDNS/Clients/Client[@uid=\"1\"]/ServiceReference"
const mySagemcomBoxDynDNSEnable string = "Device/Services/DynamicDNS/Clients/Client[@uid=\"1\"]/Enable"
const mySagemcomBoxDynDNSServiceQueryIndex string = "Device/Services/DynamicDNS/Services/Service[@uid=\"#\"]"
const mySagemcomBoxDynDNSClient string = "Device/Services/DynamicDNS/Clients/Client[@uid=\"1\"]"
const mySagemcomBoxDynDNSServices string = "Device/Services/DynamicDNS/Services"
const mySagemcomBoxDynDNSStatus string = "Device/Services/DynamicDNS/Clients/Client[@uid=\"1\"]/Status"
const mySagemcomBoxDynDNSUsername string = "Device/Services/DynamicDNS/Clients/Client[@uid=\"1\"]/Username"
const mySagemcomBoxBasicStatusDownstreamSyncSpeedDsl string = "Device/DSL/Channels/Channel[@uid=\"1\"]/DownstreamCurrRate"
const mySagemcomBoxBasicStatusAccessControl string = "Device/Hosts/Hosts/Host/BlacklistedSchedule"
const mySagemcomBoxBasicStatusFirmwareUpdated string = "Device/DeviceInfo/BuildDate"
const mySagemcomBoxBasicStatusFrontLedColor string = "Device/Managers/HubLightControl/FrontLEDColor"
const mySagemcomBoxBasicStatusDataUsageSent string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/Stats/BytesSent"
const mySagemcomBoxBasicStatusDataUsageReceived string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/Stats/BytesReceived"
const mySagemcomBoxBasicStatusNetworkUptime string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/LastChange"
const mySagemcomBoxBasicStatusUpstreamSyncSpeedDsl string = "Device/DSL/Channels/Channel[@uid=\"1\"]/UpstreamCurrRate"
const mySagemcomBoxBasicStatusFirmwareVersion string = "Device/DeviceInfo/ExternalFirmwareVersion"
const mySagemcomBoxUpnpUpnpStatus string = "Device/UPnP/Device/UPnPIGD"
const mySagemcomBoxUpnpUpnpSecurityStatus string = "Device/UPnP/Settings/ExtendedUPnPSecurity"
const mySagemcomBoxMaintenanceSaveRestoreRestore string = "Device/DeviceInfo/VendorConfigFiles/VendorConfigFile[Alias='DEVICE_CONFIG']"
const mySagemcomBoxMaintenanceSaveRestoreSave string = "Device/DeviceInfo/VendorConfigFiles/VendorConfigFile[Alias=\"DEVICE_CONFIG\"]"
const mySagemcomBoxMaintenanceNtpLocalTime string = "Device/Time/CurrentLocalTime"
const mySagemcomBoxMaintenanceFirmwareVersion string = "Device/DeviceInfo/ExternalFirmwareVersion"
const mySagemcomBoxDeviceInfoWifiPasswordForbiddenWords string = "Device/Services/DeviceConfig/WifiPasswordForbiddenWords"
const mySagemcomBoxDeviceInfoHubLightBrightness string = "Device/Managers/HubLightControl/Brightness"
const mySagemcomBoxDeviceInfoDatapumpVersion string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/FirmwareVersion"
const mySagemcomBoxDeviceInfoHubLightSchedule string = "Device/Managers/HubLightControl/Schedule/Enable"
const mySagemcomBoxDeviceInfoWanMacAddress string = "Device/DeviceInfo/MACAddress"
const mySagemcomBoxDeviceInfoIPBRLANIpv6 string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Addresses/IPv6Address/IPAddress"
const mySagemcomBoxDeviceInfoInterfaceType string = "Device/Services/DeviceConfig/LastSuccesfulWanType"
const mySagemcomBoxDeviceInfoPublicIpv4 string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv4Addresses/IPv4Address[Alias=\"IP_DATA_ADDRESS\"]/IPAddress"
const mySagemcomBoxDeviceInfoLocalIpv4 string = "Device/Managers/NetworkData/IpLan"
const mySagemcomBoxDeviceInfoLocalEthernetMac string = "Device/Ethernet/Interfaces/Interface[Alias='PHY1']/MACAddress"
const mySagemcomBoxDeviceInfoHardwareVersion string = "Device/DeviceInfo/HardwareVersion"
const mySagemcomBoxDeviceInfoPublicSubnetMask string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv4Addresses/IPv4Address[Alias=\"IP_DATA_ADDRESS\"]/SubnetMask"
const mySagemcomBoxDeviceInfoWifi24UpTime string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV2G\"]/ConnectionTime"
const mySagemcomBoxDeviceInfoWifi24MacAddress string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV2G\"]/MACAddress"
const mySagemcomBoxDeviceInfoWifi24OperatingMode string = "Device/WiFi/Radios/Radio[Alias=\"RADIO2G4\"]/OperatingStandards"
const mySagemcomBoxDeviceInfoWifi24MaxBitRate string = "Device/WiFi/Radios/Radio[Alias=\"RADIO2G4\"]/MaxBitRate"
const mySagemcomBoxDeviceInfoWifi24ChannelBandwidth string = "Device/WiFi/Radios/Radio[Alias=\"RADIO2G4\"]/OperatingChannelBandwidth"
const mySagemcomBoxDeviceInfoWifi24SecurityMode string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"PRIV0\"]/Security/ModeEnabled"
const mySagemcomBoxDeviceInfoWifi24Ssid string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV2G\"]/SSID"
const mySagemcomBoxDeviceInfoWifi24Status string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV2G\"]/Status"
const mySagemcomBoxDeviceInfoProductClass string = "Device/DeviceInfo/ProductClass"
const mySagemcomBoxDeviceInfoSerialNumber string = "Device/DeviceInfo/SerialNumber"
const mySagemcomBoxDeviceInfoWanInternetStatus string = "Device/IP/Interfaces/Interface[Alias='IP_DATA']/Status"
const mySagemcomBoxDeviceInfoWanWanStatus string = "Device/PPP/Interfaces/Interface[Alias='PPP_FTTH_DATA']/Status"
const mySagemcomBoxDeviceInfoWanFtthLowerLayer string = "Device/Ethernet/Links/Link[Alias=\"FTTH_DATA\"]/LowerLayers"
const mySagemcomBoxDeviceInfoWanRealConnectionStatus string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_FTTH_DATA\"]/ConnectionStatus"
const mySagemcomBoxDeviceInfoWanConnectionStatusPhy4 string = "Device/Ethernet/Interfaces/Interface[Alias=\"PHY4\"]/Status"
const mySagemcomBoxDeviceInfoWanConnectionStatusPhy6 string = "Device/Ethernet/Interfaces/Interface[Alias=\"PHY6_WAN\"]/Status"
const mySagemcomBoxDeviceInfoDNS string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv4Addresses/IPv4Address[Alias=\"IP_DATA_ADDRESS\"]/Dns"
const mySagemcomBoxDeviceInfoLinkLocalIpv6 string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_LAN\"]/IPv6Addresses/IPv6Address[Alias=\"LINK_LOCAL\"]/IPAddress"
const mySagemcomBoxDeviceInfoDefaultGateway string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/IPv4Addresses/IPv4Address/IPGateway"
const mySagemcomBoxDeviceInfoUpTime string = "Device/DeviceInfo/UpTime"
const mySagemcomBoxDeviceInfoHubLightStatus string = "Device/Managers/HubLightControl/LedEnable"
const mySagemcomBoxDeviceInfoLocalSubnetMask string = "Device/Managers/NetworkData/NetmaskLan"
const mySagemcomBoxDeviceInfoXdslRealConnectionStatus string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_DSL_DATA\"]/ConnectionStatus"
const mySagemcomBoxDeviceInfoXdslConnectionStatus string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/Status"
const mySagemcomBoxDeviceInfoXdslNumericActualRateUp string = "Device/DSL/Channels/Channel[@uid=\"1\"]/UpstreamCurrRate"
const mySagemcomBoxDeviceInfoXdslNumericActualRateDown string = "Device/DSL/Channels/Channel[@uid=\"1\"]/DownstreamCurrRate"
const mySagemcomBoxDeviceInfoXdslConnectionTimeDsl string = "Device/DSL/Lines/Line[@uid=\"1\"]/LastChange"
const mySagemcomBoxDeviceInfoSoftwareVersion string = "Device/DeviceInfo/ExternalFirmwareVersion"
const mySagemcomBoxDeviceInfoWifi5UpTime string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV5G\"]/ConnectionTime"
const mySagemcomBoxDeviceInfoWifi5MacAddress string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV5G\"]/MACAddress"
const mySagemcomBoxDeviceInfoWifi5OperatingMode string = "Device/WiFi/Radios/Radio[Alias=\"RADIO5G\"]/OperatingStandards"
const mySagemcomBoxDeviceInfoWifi5MaxBitRate string = "Device/WiFi/Radios/Radio[Alias=\"RADIO5G\"]/MaxBitRate"
const mySagemcomBoxDeviceInfoWifi5ChannelBandwidth string = "Device/WiFi/Radios/Radio[Alias=\"RADIO5G\"]/OperatingChannelBandwidth"
const mySagemcomBoxDeviceInfoWifi5SecurityMode string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"PRIV1\"]/Security/ModeEnabled"
const mySagemcomBoxDeviceInfoWifi5Ssid string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV5G\"]/SSID"
const mySagemcomBoxDeviceInfoWifi5Status string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV5G\"]/Status"
const mySagemcomBoxDhcpSubnetMaskDefault string = "Device/Managers/NetworkData/NetmaskLan"
const mySagemcomBoxDhcpIpv4Address2 string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/IPInterface"
const mySagemcomBoxDhcpIpv4PoolEnd2 string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/MaxAddress"
const mySagemcomBoxDhcpIpv4PoolStart2 string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/MinAddress"
const mySagemcomBoxDhcpIpv4Address string = "Device/Managers/NetworkData/IpLan"
const mySagemcomBoxDhcpSubnetMask string = "Device/Managers/NetworkData/NetmaskLan"
const mySagemcomBoxDhcpDhcpAuthoritative string = "Device/DHCPv4/Server/X_SAGEMCOM_Authoritative"
const mySagemcomBoxDhcpEnable string = "Device/DHCPv4/Server/Enable"
const mySagemcomBoxDhcpIpv4PoolStart string = "Device/Managers/NetworkData/DhcpLanMinAddress"
const mySagemcomBoxDhcpHost string = "Device/DNS/Client/HostName"
const mySagemcomBoxDhcpDhcpv4LeaseTime string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/LeaseTime"
const mySagemcomBoxDhcpIPReservation string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/StaticAddresses"
const mySagemcomBoxDhcpSubnetMask2 string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"DEFAULT_POOL\"]/SubnetMask"
const mySagemcomBoxDhcpIpv4PoolEnd string = "Device/Managers/NetworkData/DhcpLanMaxAddress"
const adminAdvancedConnectionUptime string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/LastChange"
const wifiMacfilterWifiControlList string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/MACFiltering/MACAddresses"
const wifiMacfilterMacFilterMode string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/MACFiltering/Mode"
const wifiGeneralBroadcast string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/SSIDAdvertisementEnabled"
const wifiGeneralWpsEnable string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/WPS/Enable"
const wifiGeneralChannel string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/Channel"
const wifiGeneralSecurityModesSupported string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/ModesSupported"
const wifiGeneralSecurityMode string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/ModeEnabled"
const wifiGeneralPasswordWep string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/WEPKey"
const wifiGeneralSsid string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/SSID"
const wifiGeneralMac string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/MACAddress"
const wifiGeneralPasswordKey string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/KeyPassphrase"
const wifiGeneralChannelsInUse string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/ChannelsInUse"
const wifiGeneralAutoChannel string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/AutoChannelEnable"
const wifiGeneralPossibleChannels string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/PossibleChannels"
const wifiGeneralEnable string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/Enable"
const wifiGeneralOperatingStandards string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingStandards"
const wifiGeneralChannelBandwidth string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingChannelBandwidth"
const wifiGeneralPossibleSecurityMode string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/ModesSupported"
const wifiGeneralAutoChannelTrigger string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/AutoChannelTrigger"
const wifiGeneralWirelessVisible string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Enable"
const wifiGeneralExtensionChannel string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/ExtensionChannel"
const wifiGeneralSsidEnable string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/Enable"
const wifiGeneralTransmitPower string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/TransmitPower"
const wifiRPCWpsRPC string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]"
const wifiStatsMaxBitRate string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/MaxBitRate"
const wifiStatsStatus string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/Status"
const wifiAdvancedGlobalMaxClients string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/MaxAssociatedDevices"
const wifiAdvancedBandSteering string = "Device/WiFi/BandSteering/Enable"
const wifiAdvancedOperatingStandards string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingStandards"
const wifiAdvancedChannelBandwidth string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingChannelBandwidth"
const wifiAdvancedWmmEnable string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/WMMEnable"
const wifiAdvancedExtensionChannel string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/ExtensionChannel"
const wifiAdvancedApsd string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/UAPSDEnable"
const wifiAdvancedRegulatoryDomain string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/RegulatoryDomain"
const wifiAdvancedTransmitPower string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/TransmitPower"
const wifiMainPageMode string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/ModeEnabled"
const wifiMainPageOperatingMode string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingStandards"
const wifiMainPageMaxBitRate string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/MaxBitRate"
const wifiMainPageChannelBandwidth string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingChannelBandwidth"
const wifiMainPageAssociatedDevices string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/AssociatedDevices"
const wifiMainPageSsid string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]"
const wifiIndexOperatingMode string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingStandards"
const wifiIndexWpsEnable string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/WPS/Enable"
const wifiIndexChannel string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/Channel"
const wifiIndexSecurityMode string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/ModeEnabled"
const wifiIndexSsid string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/SSID"
const wifiIndexUpTime string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/ConnectionTime"
const wifiIndexMacAddress string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/MACAddress"
const wifiIndexPassword string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Security/KeyPassphrase"
const wifiIndexAutoChannel string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/AutoChannelEnable"
const wifiIndexEnable string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/Enable"
const wifiIndexMaxBitRate string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/MaxBitRate"
const wifiIndexChannelBandwidth string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/OperatingChannelBandwidth"
const wifiIndexWirelessVisible string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"#AP#\"]/Enable"
const wifiIndexStatus string = "Device/WiFi/SSIDs/SSID[Alias=\"#SSID#\"]/Status"
const wifiRecommendedChannelTestChannel string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/SiteSurvey/ChannelsToTest"
const wifiRecommendedChannelSubscribe string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/SiteSurvey/State"
const wifiRecommendedChannelSurvey string = "Device/WiFi/Radios/Radio[Alias=\"#RADIO#\"]/SiteSurvey/ChannelSurveys"
const wifiRedsideAliases string = "Device/WiFi/SSIDs/SSID/Alias"
const wifiGeneralCommonWifiAutoChannelTrigger string = "Device/WiFi/Radios/Radio/AutoChannelTrigger"
const wifiGeneralCommonWifiMode string = "Device/Services/DeviceConfig/WiFiMode"
const domainLockList string = "Device/Services/DeviceConfig/DomainLockList"
const domainLockEnable string = "Device/Services/DeviceConfig/DomainLockingEnable"
const deviceNATPortMappingsDMZInternalClient string = "Device/NAT/PortMappings/PortMapping[Service=\"DMZ\"]/InternalClient"
const deviceNATPortMappingsDMZEnable string = "Device/NAT/PortMappings/PortMapping[Service=\"DMZ\"]/Enable"
const deviceNATPortMappingsNotDMZ string = "Device/NAT/PortMappings/PortMapping[Alias!=\"API_DMZ\"]"
const pinholesSaveMacID string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/MacId"
const pinholesSaveDescription string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/Description"
const pinholesSaveDestIP string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/DestIP"
const pinholesSaveDestPort string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/DestPort"
const pinholesSaveSourcePortRangeMax string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/SourcePortRangeMax"
const pinholesSaveDestPortRangeMax string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/DestPortRangeMax"
const pinholesSaveList string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules"
const pinholesSaveProtocol string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/Protocol"
const pinholesSaveSourcePort string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]/SourcePort"
const pinholesSaveRemove string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]"
const pinholesRule string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[@uid=\"#\"]"
const pinholesList string = "Device/Firewall/Chains/Chain[Name=\"Low\"]/Rules/Rule[IPVersion=\"6\"]"
const firstConnection string = "Device/DeviceInfo/FirstConnection"
const technicalLogDataSent string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/Stats/BytesSent"
const technicalLogBtWifi string = "Device/Services/OpenWiFi/OpenWifiOptedIn"
const technicalLogInterleaveDelayUs string = "Device/DSL/Channels/Channel[Alias=\"DSL0\"]/ActualInterleavingDelayus"
const technicalLogModulation string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/StandardUsed"
const technicalLogBoardVersion string = "Device/DeviceInfo/HardwareVersion"
const technicalLogMaximumDataRateDown string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/DownstreamMaxBitRate"
const technicalLogDataReceived string = "Device/IP/Interfaces/Interface[Alias=\"IP_DATA\"]/Stats/BytesReceived"
const technicalLogWifiMode string = "Device/Services/DeviceConfig/WiFiMode"
const technicalLogProductName string = "Device/DeviceInfo/ModelName"
const technicalLogFirmwareUpdated string = "Device/DeviceInfo/BuildDate"
const technicalLogSignalAttenuationDown string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/SignalDownstreamAttenuation"
const technicalLogMaximumDataRateUp string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/UpstreamMaxBitRate"
const technicalLogWifiSecurity string = "Device/WiFi/AccessPoints/AccessPoint[Alias=\"PRIV0\"]/Security/ModeEnabled"
const technicalLogNoiseMarginUp string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/UpstreamNoiseMargin"
const technicalLogVpiVci string = "Device/ATM/Links/Link[Alias=\"ATM_DATA\"]/DestinationAddress"
const technicalLogFirmwareVersion string = "Device/DeviceInfo/ExternalFirmwareVersion"
const technicalLogFtthStatus string = "Device/PPP/Interfaces/Interface[Alias='PPP_FTTH_DATA']/Status"
const technicalLogBootLoader string = "Device/DeviceInfo/BootloaderVersion"
const technicalLogBroadbandUsernameFTTH string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_FTTH_DATA\"]/Username"
const technicalLogSerialNumber string = "Device/DeviceInfo/SerialNumber"
const technicalLogWifiSSID24 string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV2G\"]/SSID"
const technicalLogBandSteering string = "Device/WiFi/BandSteering/Enable"
const technicalLogInterleaveDelay string = "Device/DSL/Channels/Channel[Alias=\"DSL0\"]/ActualInterleavingDelay"
const technicalLogWifiChannel24 string = "Device/WiFi/Radios/Radio[Alias=\"RADIO2G4\"]/Channel"
const technicalLogLineAttenuation string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/DownstreamAttenuation"
const technicalLogNoiseMarginDown string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/DownstreamNoiseMargin"
const technicalLogWifiSmart5 string = "Device/WiFi/Radios/Radio[Alias=\"RADIO5G\"]/AutoChannelEnable"
const technicalLogBroadbandUsernameDSL string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_DSL_DATA\"]/Username"
const technicalLogDataRateDown string = "Device/DSL/Channels/Channel[Alias=\"DSL0\"]/DownstreamCurrRate"
const technicalLogSignalAttenuationUp string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/SignalUpstreamAttenuation"
const technicalLogWifiSSID5 string = "Device/WiFi/SSIDs/SSID[Alias=\"WL_PRIV5G\"]/SSID"
const technicalLogMacAddress string = "Device/DeviceInfo/MACAddress"
const technicalLogWifiChannel5 string = "Device/WiFi/Radios/Radio[Alias=\"RADIO5G\"]/Channel"
const technicalLogDslUptime string = "Device/DSL/Lines/Line[Alias=\"DSL0\"]/LastChange"
const technicalLogFirewall string = "Device/Firewall/Enable"
const technicalLogWifiSmart24 string = "Device/WiFi/Radios/Radio[Alias=\"RADIO2G4\"]/AutoChannelEnable"
const technicalLogDataRateUp string = "Device/DSL/Channels/Channel[Alias=\"DSL0\"]/UpstreamCurrRate"
const technicalLogDslStatus string = "Device/PPP/Interfaces/Interface[Alias='PPP_DSL_DATA']/Status"
const managementTr69InternalDataPort string = "Device/ManagementServer/TR69InternalData/Settings/Port"
const managementLanguage string = "Device/UserInterface/CurrentLanguage"
const fingerPrintDatabaseEntry string = "Device/Services/OnlineInstall/DHCPFingerprintDatabase/Entries"
const hurlProductVariant string = "Device/DeviceInfo/ProductVariant"
const hurlEnable string = "Device/Services/DeviceConfig/ConnectionHURLPageEnable"
const hubLightControlTurnLightOn string = "Device/Managers/HubLightControl/Schedule/TurnLightOn"
const hubLightControlTurnLightOff string = "Device/Managers/HubLightControl/Schedule/TurnLightOff"
const hubLightControlLedStatus string = "Device/Managers/HubLightControl/LedEnable"
const hubLightControlScheduleEnable string = "Device/Managers/HubLightControl/Schedule/Enable"
const hubLightControlBrightness string = "Device/Managers/HubLightControl/Brightness"
const forbiddenIpsIpsQuantennaBootIPLocal string = "Device/WiFi/Quantenna/BootIPLocal"
const forbiddenIpsIpsPppFtthDataRemoteIP string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_FTTH_DATA\"]/IPCP/RemoteIPAddress"
const forbiddenIpsIpsDataModels string = "Device/DeviceInfo/SupportedDataModels/SupportedDataModel/URL"
const forbiddenIpsIpsQuantennaMngtIPRemote string = "Device/WiFi/Quantenna/MngtIPRemote"
const forbiddenIpsIpsOpenWifiIPInterface string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_OPENWIFI\"]/IPv4Addresses/IPv4Address[Alias=\"IP_BR_OPENWIFI_ADDRESS\"]/IPAddress"
const forbiddenIpsIpsQuantennaMngtIPLocal string = "Device/WiFi/Quantenna/MngtIPLocal"
const forbiddenIpsIpsOpenWifiMinAddress string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"OPENWIFI_POOL\"]/MinAddress"
const forbiddenIpsIpsOpenWifiMaxAddress string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"OPENWIFI_POOL\"]/MaxAddress"
const forbiddenIpsIpsOpenWifiSubnetMask string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"OPENWIFI_POOL\"]/SubnetMask"
const forbiddenIpsIpsPppDslDataLocalIP string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_DSL_DATA\"]/IPCP/LocalIPAddress"
const forbiddenIpsIpsOpenWifiInterface string = "Device/DHCPv4/Server/Pools/Pool[Alias=\"OPENWIFI_POOL\"]/IPInterface"
const forbiddenIpsIpsOpenWifiIPSubnet string = "Device/IP/Interfaces/Interface[Alias=\"IP_BR_OPENWIFI\"]/IPv4Addresses/IPv4Address[Alias=\"IP_BR_OPENWIFI_ADDRESS\"]/SubnetMask"
const forbiddenIpsIpsPppDslDataRemoteIP string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_DSL_DATA\"]/IPCP/RemoteIPAddress"
const forbiddenIpsIpsPppFtthDataLocalIP string = "Device/PPP/Interfaces/Interface[Alias=\"PPP_FTTH_DATA\"]/IPCP/LocalIPAddress"
const forbiddenIpsIpsQuantennaBootIPRemote string = "Device/WiFi/Quantenna/BootIPRemote"
const gatewayConnectionStatusLocalipnetworkIpv4address string = "Device/Managers/NetworkData/IpLan"
const gatewayConnectionStatusLocalipnetworkSubnetmask string = "Device/Managers/NetworkData/NetmaskLan"
const device string = "Device"
