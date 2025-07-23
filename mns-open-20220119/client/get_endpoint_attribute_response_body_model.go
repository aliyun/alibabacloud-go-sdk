// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEndpointAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetEndpointAttributeResponseBody
	GetCode() *int64
	SetData(v *GetEndpointAttributeResponseBodyData) *GetEndpointAttributeResponseBody
	GetData() *GetEndpointAttributeResponseBodyData
	SetMessage(v string) *GetEndpointAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEndpointAttributeResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetEndpointAttributeResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetEndpointAttributeResponseBody
	GetSuccess() *bool
}

type GetEndpointAttributeResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetEndpointAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEndpointAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetEndpointAttributeResponseBody) GetData() *GetEndpointAttributeResponseBodyData {
	return s.Data
}

func (s *GetEndpointAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEndpointAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEndpointAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEndpointAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEndpointAttributeResponseBody) SetCode(v int64) *GetEndpointAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetData(v *GetEndpointAttributeResponseBodyData) *GetEndpointAttributeResponseBody {
	s.Data = v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetMessage(v string) *GetEndpointAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetRequestId(v string) *GetEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetStatus(v string) *GetEndpointAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) SetSuccess(v bool) *GetEndpointAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *GetEndpointAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEndpointAttributeResponseBodyData struct {
	// The list of CIDR block.
	CidrList []*GetEndpointAttributeResponseBodyDataCidrList `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
	// Specifies whether the endpoint is enabled.
	//
	// example:
	//
	// true
	EndpointEnabled *bool `json:"EndpointEnabled,omitempty" xml:"EndpointEnabled,omitempty"`
}

func (s GetEndpointAttributeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponseBodyData) GetCidrList() []*GetEndpointAttributeResponseBodyDataCidrList {
	return s.CidrList
}

func (s *GetEndpointAttributeResponseBodyData) GetEndpointEnabled() *bool {
	return s.EndpointEnabled
}

func (s *GetEndpointAttributeResponseBodyData) SetCidrList(v []*GetEndpointAttributeResponseBodyDataCidrList) *GetEndpointAttributeResponseBodyData {
	s.CidrList = v
	return s
}

func (s *GetEndpointAttributeResponseBodyData) SetEndpointEnabled(v bool) *GetEndpointAttributeResponseBodyData {
	s.EndpointEnabled = &v
	return s
}

func (s *GetEndpointAttributeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetEndpointAttributeResponseBodyDataCidrList struct {
	// The ACL policy. Valid values:
	//
	// 	- **allow**: indicates that the current endpoint allows access from the corresponding CIDR block. (Only allow is supported.)
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// The CIDR block.
	//
	// example:
	//
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1701951224000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s GetEndpointAttributeResponseBodyDataCidrList) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointAttributeResponseBodyDataCidrList) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) GetAclStrategy() *string {
	return s.AclStrategy
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) GetCidr() *string {
	return s.Cidr
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) SetAclStrategy(v string) *GetEndpointAttributeResponseBodyDataCidrList {
	s.AclStrategy = &v
	return s
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) SetCidr(v string) *GetEndpointAttributeResponseBodyDataCidrList {
	s.Cidr = &v
	return s
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) SetCreateTime(v int64) *GetEndpointAttributeResponseBodyDataCidrList {
	s.CreateTime = &v
	return s
}

func (s *GetEndpointAttributeResponseBodyDataCidrList) Validate() error {
	return dara.Validate(s)
}
