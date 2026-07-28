// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpnGatewayDiagnoseResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *GetVpnGatewayDiagnoseResultResponseBody
	GetBeginTime() *string
	SetDiagnoseId(v string) *GetVpnGatewayDiagnoseResultResponseBody
	GetDiagnoseId() *string
	SetDiagnoseResult(v []*GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) *GetVpnGatewayDiagnoseResultResponseBody
	GetDiagnoseResult() []*GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult
	SetFinishTime(v string) *GetVpnGatewayDiagnoseResultResponseBody
	GetFinishTime() *string
	SetFinishedCount(v int32) *GetVpnGatewayDiagnoseResultResponseBody
	GetFinishedCount() *int32
	SetRequestId(v string) *GetVpnGatewayDiagnoseResultResponseBody
	GetRequestId() *string
	SetResourceInstanceId(v string) *GetVpnGatewayDiagnoseResultResponseBody
	GetResourceInstanceId() *string
	SetResourceType(v string) *GetVpnGatewayDiagnoseResultResponseBody
	GetResourceType() *string
	SetTotalCount(v int32) *GetVpnGatewayDiagnoseResultResponseBody
	GetTotalCount() *int32
	SetVpnGatewayId(v string) *GetVpnGatewayDiagnoseResultResponseBody
	GetVpnGatewayId() *string
}

type GetVpnGatewayDiagnoseResultResponseBody struct {
	// The time when the diagnostic started.
	//
	// The time is displayed in UTC in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-12-15T05:28:57Z
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The diagnostic ID.
	//
	// example:
	//
	// vpndgn-uf6sgneym02lxyuv4****
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// The list of diagnostic items.
	DiagnoseResult []*GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult `json:"DiagnoseResult,omitempty" xml:"DiagnoseResult,omitempty" type:"Repeated"`
	// The time when the diagnostic ended.
	//
	// The time is displayed in UTC in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-12-15T05:29:08Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The number of diagnostic items that have been completed.
	//
	// example:
	//
	// 7
	FinishedCount *int32 `json:"FinishedCount,omitempty" xml:"FinishedCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 312C4D5A-6563-5FC6-8C6E-A43A5A316FEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the diagnosed resource.
	//
	// example:
	//
	// vco-uf6huqsu63azl7mdp****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The type of the diagnosed resource.
	//
	// Valid values: **IPsec**, which indicates an IPsec-VPN connection.
	//
	// example:
	//
	// IPsec
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The total number of diagnostic items.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The VPN gateway instance ID.
	//
	// example:
	//
	// vpn-uf6fzwp0ck3frwtbk****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s GetVpnGatewayDiagnoseResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpnGatewayDiagnoseResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetDiagnoseResult() []*GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult {
	return s.DiagnoseResult
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetFinishedCount() *int32 {
	return s.FinishedCount
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetBeginTime(v string) *GetVpnGatewayDiagnoseResultResponseBody {
	s.BeginTime = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetDiagnoseId(v string) *GetVpnGatewayDiagnoseResultResponseBody {
	s.DiagnoseId = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetDiagnoseResult(v []*GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) *GetVpnGatewayDiagnoseResultResponseBody {
	s.DiagnoseResult = v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetFinishTime(v string) *GetVpnGatewayDiagnoseResultResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetFinishedCount(v int32) *GetVpnGatewayDiagnoseResultResponseBody {
	s.FinishedCount = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetRequestId(v string) *GetVpnGatewayDiagnoseResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetResourceInstanceId(v string) *GetVpnGatewayDiagnoseResultResponseBody {
	s.ResourceInstanceId = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetResourceType(v string) *GetVpnGatewayDiagnoseResultResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetTotalCount(v int32) *GetVpnGatewayDiagnoseResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) SetVpnGatewayId(v string) *GetVpnGatewayDiagnoseResultResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBody) Validate() error {
	if s.DiagnoseResult != nil {
		for _, item := range s.DiagnoseResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult struct {
	// The diagnostic item.
	//
	// - **RouteEntryConflict**: route conflict.
	//
	// - **VpnRouteQuota**: VPN gateway destination route quota.
	//
	// - **VpnIPsecQuota**: VPN gateway IPsec-VPN connection quota.
	//
	// - **VpnPbrRouteQuota**: VPN gateway policy-based route quota.
	//
	// - **VcoConfigConsistency**: IPsec configuration consistency.
	//
	// - **VcoUserInternetIpConnectivity**: public network connectivity of the customer gateway.
	//
	// - **VcoPrivateConnectivity**: private network connectivity.
	//
	// For more information about each diagnostic item, see [One-click diagnostics background information](https://help.aliyun.com/document_detail/190330.html).
	//
	// example:
	//
	// RouteEntryConflict
	DiagnoseName *string `json:"DiagnoseName,omitempty" xml:"DiagnoseName,omitempty"`
	// The diagnostic result of the diagnostic item.
	//
	// The operation returns different information for each diagnostic item:
	//
	// - **RouteEntryConflict**: The system returns information about the route conflict.
	//
	// - **VpnRouteQuota**:
	//
	//     - **quotaName**: The destination route quota ID.
	//
	//     - **quantity**: The number of destination routes that the current VPN gateway instance supports.
	//
	//     - **used**: The number of destination routes that have been created for the current VPN gateway instance.
	//
	// - **VpnIPsecQuota**:
	//
	//     - **quotaName**: The IPsec-VPN connection quota ID.
	//
	//     - **quantity**: The number of IPsec-VPN connections that the current VPN gateway instance supports.
	//
	//     - **used**: The number of IPsec-VPN connections that have been created for the current VPN gateway instance.
	//
	// - **VpnPbrRouteQuota**:
	//
	//     - **quotaName**: The policy-based route quota ID.
	//
	//     - **quantity**: The number of policy-based routes that the current VPN gateway instance supports.
	//
	//     - **used**: The number of policy-based routes that have been created for the current VPN gateway instance.
	//
	// - **VcoConfigConsistency**:
	//
	//     - **vcoLackConf**: The system cannot obtain the configuration of the peer end of the IPsec-VPN connection.
	//
	//     - **vcoRunningConf**: The configuration that has been added to the peer end of the IPsec-VPN connection.
	//
	//     - **vcoDiffConf**: The list of configurations that are inconsistent between the local end and the peer end of the IPsec-VPN connection.
	//
	//     - **vcoConf**: The configuration that has been added to the local end of the IPsec-VPN connection.
	//
	// - **VcoUserInternetIpConnectivity**:
	//
	//     - **targetIp**: The public IP address of the customer gateway.
	//
	//     - **rtt**: The latency when the system accesses the public IP address of the customer gateway. Unit: ms.
	//
	//     - **lossRate**: The packet loss rate when the system accesses the public IP address of the customer gateway.
	//
	// - **VcoPrivateConnectivity**:
	//
	//     - **targetIp**: The source IP address.
	//
	//     - **srcIp**: The destination IP address.
	//
	//     - **rtt**: The latency when the source IP address accesses the destination IP address. Unit: ms.
	//
	//     - **lossRate**: The packet loss rate when the source IP address accesses the destination IP address.
	//
	// example:
	//
	// {\\"targetIp\\":\\"192.168.0.1\\",\\"srcIp\\":\\"192.168.1.1\\",\\"rtt\\":-1.0,\\"lossRate\\":100.0}
	DiagnoseResultDescription *string `json:"DiagnoseResultDescription,omitempty" xml:"DiagnoseResultDescription,omitempty"`
	// The diagnostic result level of the diagnostic item.
	//
	// - **normal**: Normal.
	//
	// - **warning**: Warning.
	//
	// - **error**: Error.
	//
	// For more information about the diagnostic result levels of each diagnostic item, see [One-click diagnostics background information](https://help.aliyun.com/document_detail/190330.html).
	//
	// example:
	//
	// normal
	DiagnoseResultLevel *string `json:"DiagnoseResultLevel,omitempty" xml:"DiagnoseResultLevel,omitempty"`
}

func (s GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) String() string {
	return dara.Prettify(s)
}

func (s GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) GoString() string {
	return s.String()
}

func (s *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) GetDiagnoseName() *string {
	return s.DiagnoseName
}

func (s *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) GetDiagnoseResultDescription() *string {
	return s.DiagnoseResultDescription
}

func (s *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) GetDiagnoseResultLevel() *string {
	return s.DiagnoseResultLevel
}

func (s *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) SetDiagnoseName(v string) *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult {
	s.DiagnoseName = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) SetDiagnoseResultDescription(v string) *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult {
	s.DiagnoseResultDescription = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) SetDiagnoseResultLevel(v string) *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult {
	s.DiagnoseResultLevel = &v
	return s
}

func (s *GetVpnGatewayDiagnoseResultResponseBodyDiagnoseResult) Validate() error {
	return dara.Validate(s)
}
