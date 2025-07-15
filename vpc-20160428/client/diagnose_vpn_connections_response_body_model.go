// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseVpnConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DiagnoseVpnConnectionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DiagnoseVpnConnectionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DiagnoseVpnConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DiagnoseVpnConnectionsResponseBody
	GetTotalCount() *int32
	SetVpnConnections(v []*DiagnoseVpnConnectionsResponseBodyVpnConnections) *DiagnoseVpnConnectionsResponseBody
	GetVpnConnections() []*DiagnoseVpnConnectionsResponseBodyVpnConnections
}

type DiagnoseVpnConnectionsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B8094E1E-935B-1397-96A8-4F87A5D1BF29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The diagnostic information.
	VpnConnections []*DiagnoseVpnConnectionsResponseBodyVpnConnections `json:"VpnConnections,omitempty" xml:"VpnConnections,omitempty" type:"Repeated"`
}

func (s DiagnoseVpnConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseVpnConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnoseVpnConnectionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DiagnoseVpnConnectionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DiagnoseVpnConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiagnoseVpnConnectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DiagnoseVpnConnectionsResponseBody) GetVpnConnections() []*DiagnoseVpnConnectionsResponseBodyVpnConnections {
	return s.VpnConnections
}

func (s *DiagnoseVpnConnectionsResponseBody) SetPageNumber(v int32) *DiagnoseVpnConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBody) SetPageSize(v int32) *DiagnoseVpnConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBody) SetRequestId(v string) *DiagnoseVpnConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBody) SetTotalCount(v int32) *DiagnoseVpnConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBody) SetVpnConnections(v []*DiagnoseVpnConnectionsResponseBodyVpnConnections) *DiagnoseVpnConnectionsResponseBody {
	s.VpnConnections = v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DiagnoseVpnConnectionsResponseBodyVpnConnections struct {
	// The cause of the error.
	//
	// example:
	//
	// Phase1 negotiation timeout
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The error code.
	//
	// example:
	//
	// Phase1NegotiationTimeout
	FailedReasonCode *string `json:"FailedReasonCode,omitempty" xml:"FailedReasonCode,omitempty"`
	// The timestamp when the current error occurred on the IPsec-VPN connection. Unit: millisecond.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1673581161000
	FailedTime *int64 `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
	// If the values of the parameters configured for the IPsec-VPN connection and the peer gateway device do not match, this parameter indicates the value of the parameters configured for the IPsec-VPN connection.
	//
	// example:
	//
	// SHA256
	MismatchLocalParam *string `json:"MismatchLocalParam,omitempty" xml:"MismatchLocalParam,omitempty"`
	// If the parameter values configured for the IPsec-VPN connection and the peer gateway device do not match, this parameter indicates the value of the parameter configured for the peer gateway device.
	//
	// example:
	//
	// SHA
	MismatchRemoteParam *string `json:"MismatchRemoteParam,omitempty" xml:"MismatchRemoteParam,omitempty"`
	// The error level. Valid values:
	//
	// 	- **Critical**
	//
	// 	- **Warn**
	//
	// 	- **Normal**
	//
	// example:
	//
	// Warn
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The log information about the error.
	//
	// example:
	//
	// 2023-01-13 11:39:21 vco-bp1spxu8hlcvpd7ry***	- [PROTO_ERR]: ikev1.c:1433:isakmp_ph1resend(): phase1 negotiation failed due to time up. [{remote id:4}{ph1: 172.16.0.88[500] <=> 192.168.0.206[500], 172.16.0.88 <=> 192.168.0.206}]
	SourceLog *string `json:"SourceLog,omitempty" xml:"SourceLog,omitempty"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-64n1sr9dig64k6****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-bp1spxu8hlcvpd7ry****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DiagnoseVpnConnectionsResponseBodyVpnConnections) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseVpnConnectionsResponseBodyVpnConnections) GoString() string {
	return s.String()
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetFailedReasonCode() *string {
	return s.FailedReasonCode
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetFailedTime() *int64 {
	return s.FailedTime
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetMismatchLocalParam() *string {
	return s.MismatchLocalParam
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetMismatchRemoteParam() *string {
	return s.MismatchRemoteParam
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetSeverity() *string {
	return s.Severity
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetSourceLog() *string {
	return s.SourceLog
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetFailedReason(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.FailedReason = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetFailedReasonCode(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.FailedReasonCode = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetFailedTime(v int64) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.FailedTime = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetMismatchLocalParam(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.MismatchLocalParam = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetMismatchRemoteParam(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.MismatchRemoteParam = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetSeverity(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.Severity = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetSourceLog(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.SourceLog = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetTunnelId(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.TunnelId = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) SetVpnConnectionId(v string) *DiagnoseVpnConnectionsResponseBodyVpnConnections {
	s.VpnConnectionId = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponseBodyVpnConnections) Validate() error {
	return dara.Validate(s)
}
