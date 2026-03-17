// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayDnsForwardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ViewSmartAccessGatewayDnsForwardsResponseBody
	GetCode() *string
	SetCount(v int32) *ViewSmartAccessGatewayDnsForwardsResponseBody
	GetCount() *int32
	SetData(v []*ViewSmartAccessGatewayDnsForwardsResponseBodyData) *ViewSmartAccessGatewayDnsForwardsResponseBody
	GetData() []*ViewSmartAccessGatewayDnsForwardsResponseBodyData
	SetHttpStatusCode(v int32) *ViewSmartAccessGatewayDnsForwardsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ViewSmartAccessGatewayDnsForwardsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ViewSmartAccessGatewayDnsForwardsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ViewSmartAccessGatewayDnsForwardsResponseBody
	GetSuccess() *bool
}

type ViewSmartAccessGatewayDnsForwardsResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1914
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A DNS forwarding list.
	Data []*ViewSmartAccessGatewayDnsForwardsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// E223E535-AE11-4158-B00F-DC107887A909
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ViewSmartAccessGatewayDnsForwardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsForwardsResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) GetData() []*ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	return s.Data
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) SetCode(v string) *ViewSmartAccessGatewayDnsForwardsResponseBody {
	s.Code = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) SetCount(v int32) *ViewSmartAccessGatewayDnsForwardsResponseBody {
	s.Count = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) SetData(v []*ViewSmartAccessGatewayDnsForwardsResponseBodyData) *ViewSmartAccessGatewayDnsForwardsResponseBody {
	s.Data = v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) SetHttpStatusCode(v int32) *ViewSmartAccessGatewayDnsForwardsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) SetMessage(v string) *ViewSmartAccessGatewayDnsForwardsResponseBody {
	s.Message = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayDnsForwardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) SetSuccess(v bool) *ViewSmartAccessGatewayDnsForwardsResponseBody {
	s.Success = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ViewSmartAccessGatewayDnsForwardsResponseBodyData struct {
	// The domain name.
	//
	// example:
	//
	// www.baidu.com
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
	// 172.16.58.20
	MasterIp *string `json:"MasterIp,omitempty" xml:"MasterIp,omitempty"`
	// The forwarding mode.
	//
	// >
	//
	// 	- This parameter is not in use. Ignore this parameter.
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

func (s ViewSmartAccessGatewayDnsForwardsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsForwardsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetMasterIp() *string {
	return s.MasterIp
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetMode() *string {
	return s.Mode
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetOutboundPortIndex() *int32 {
	return s.OutboundPortIndex
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetOutboundPortName() *string {
	return s.OutboundPortName
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetOutboundPortType() *string {
	return s.OutboundPortType
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) GetSlaveIp() *string {
	return s.SlaveIp
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetDomain(v string) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetInstanceId(v string) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetMasterIp(v string) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.MasterIp = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetMode(v string) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.Mode = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetOutboundPortIndex(v int32) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.OutboundPortIndex = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetOutboundPortName(v string) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.OutboundPortName = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetOutboundPortType(v string) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.OutboundPortType = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) SetSlaveIp(v string) *ViewSmartAccessGatewayDnsForwardsResponseBodyData {
	s.SlaveIp = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
