// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceIpWhitelistResponseBody
	GetCode() *string
	SetData(v *GetInstanceIpWhitelistResponseBodyData) *GetInstanceIpWhitelistResponseBody
	GetData() *GetInstanceIpWhitelistResponseBodyData
	SetDynamicCode(v string) *GetInstanceIpWhitelistResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetInstanceIpWhitelistResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceIpWhitelistResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceIpWhitelistResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceIpWhitelistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceIpWhitelistResponseBody
	GetSuccess() *bool
}

type GetInstanceIpWhitelistResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetInstanceIpWhitelistResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// xxx
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B962390-D84B-5D44-8C11-79DF40299D41
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInstanceIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhitelistResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceIpWhitelistResponseBody) GetData() *GetInstanceIpWhitelistResponseBodyData {
	return s.Data
}

func (s *GetInstanceIpWhitelistResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetInstanceIpWhitelistResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetInstanceIpWhitelistResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceIpWhitelistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceIpWhitelistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceIpWhitelistResponseBody) SetCode(v string) *GetInstanceIpWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) SetData(v *GetInstanceIpWhitelistResponseBodyData) *GetInstanceIpWhitelistResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) SetDynamicCode(v string) *GetInstanceIpWhitelistResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) SetDynamicMessage(v string) *GetInstanceIpWhitelistResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) SetHttpStatusCode(v int32) *GetInstanceIpWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) SetMessage(v string) *GetInstanceIpWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) SetRequestId(v string) *GetInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) SetSuccess(v bool) *GetInstanceIpWhitelistResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceIpWhitelistResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The IP address whitelists.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetInstanceIpWhitelistResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhitelistResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhitelistResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceIpWhitelistResponseBodyData) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *GetInstanceIpWhitelistResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceIpWhitelistResponseBodyData) SetInstanceId(v string) *GetInstanceIpWhitelistResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBodyData) SetIpWhitelists(v []*string) *GetInstanceIpWhitelistResponseBodyData {
	s.IpWhitelists = v
	return s
}

func (s *GetInstanceIpWhitelistResponseBodyData) SetRegionId(v string) *GetInstanceIpWhitelistResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceIpWhitelistResponseBodyData) Validate() error {
	return dara.Validate(s)
}
