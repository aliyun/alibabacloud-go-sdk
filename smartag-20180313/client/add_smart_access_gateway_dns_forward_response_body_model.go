// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmartAccessGatewayDnsForwardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddSmartAccessGatewayDnsForwardResponseBody
	GetCode() *string
	SetData(v *AddSmartAccessGatewayDnsForwardResponseBodyData) *AddSmartAccessGatewayDnsForwardResponseBody
	GetData() *AddSmartAccessGatewayDnsForwardResponseBodyData
	SetHttpStatusCode(v int32) *AddSmartAccessGatewayDnsForwardResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddSmartAccessGatewayDnsForwardResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSmartAccessGatewayDnsForwardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSmartAccessGatewayDnsForwardResponseBody
	GetSuccess() *bool
}

type AddSmartAccessGatewayDnsForwardResponseBody struct {
	// The error code. A value of 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information returned for the request.
	Data *AddSmartAccessGatewayDnsForwardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E93884AC-6C21-4FEA-8E3A-7377D33B194F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSmartAccessGatewayDnsForwardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSmartAccessGatewayDnsForwardResponseBody) GoString() string {
	return s.String()
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) GetData() *AddSmartAccessGatewayDnsForwardResponseBodyData {
	return s.Data
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) SetCode(v string) *AddSmartAccessGatewayDnsForwardResponseBody {
	s.Code = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) SetData(v *AddSmartAccessGatewayDnsForwardResponseBodyData) *AddSmartAccessGatewayDnsForwardResponseBody {
	s.Data = v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) SetHttpStatusCode(v int32) *AddSmartAccessGatewayDnsForwardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) SetMessage(v string) *AddSmartAccessGatewayDnsForwardResponseBody {
	s.Message = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) SetRequestId(v string) *AddSmartAccessGatewayDnsForwardResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) SetSuccess(v bool) *AddSmartAccessGatewayDnsForwardResponseBody {
	s.Success = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddSmartAccessGatewayDnsForwardResponseBodyData struct {
	// The domain name.
	//
	// example:
	//
	// yfiy.cn
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// sagv3dnsforward-nc7qabskj17werc7su
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The primary DNS server.
	//
	// example:
	//
	// 14.104.81.13
	MasterIp *string `json:"MasterIp,omitempty" xml:"MasterIp,omitempty"`
	// The forwarding mode.
	//
	// example:
	//
	// first
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The number of the egress port.
	//
	// example:
	//
	// 0
	OutboundPortIndex *int32 `json:"OutboundPortIndex,omitempty" xml:"OutboundPortIndex,omitempty"`
	// The egress port.
	//
	// example:
	//
	// eth0
	OutboundPortName *string `json:"OutboundPortName,omitempty" xml:"OutboundPortName,omitempty"`
	// The type of the egress port.
	//
	// example:
	//
	// PhysicalPort
	OutboundPortType *string `json:"OutboundPortType,omitempty" xml:"OutboundPortType,omitempty"`
	// The secondary DNS server.
	//
	// example:
	//
	// 172.16.0.14
	SlaveIp *string `json:"SlaveIp,omitempty" xml:"SlaveIp,omitempty"`
}

func (s AddSmartAccessGatewayDnsForwardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddSmartAccessGatewayDnsForwardResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetMasterIp() *string {
	return s.MasterIp
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetMode() *string {
	return s.Mode
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetOutboundPortIndex() *int32 {
	return s.OutboundPortIndex
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetOutboundPortName() *string {
	return s.OutboundPortName
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetOutboundPortType() *string {
	return s.OutboundPortType
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) GetSlaveIp() *string {
	return s.SlaveIp
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetDomain(v string) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.Domain = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetInstanceId(v string) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetMasterIp(v string) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.MasterIp = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetMode(v string) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.Mode = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetOutboundPortIndex(v int32) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.OutboundPortIndex = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetOutboundPortName(v string) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.OutboundPortName = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetOutboundPortType(v string) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.OutboundPortType = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) SetSlaveIp(v string) *AddSmartAccessGatewayDnsForwardResponseBodyData {
	s.SlaveIp = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponseBodyData) Validate() error {
	return dara.Validate(s)
}
