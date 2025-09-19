// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageRegistryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyImageRegistryResponseBody
	GetCode() *string
	SetData(v interface{}) *ModifyImageRegistryResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *ModifyImageRegistryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyImageRegistryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyImageRegistryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyImageRegistryResponseBody
	GetSuccess() *bool
	SetTimeCost(v int64) *ModifyImageRegistryResponseBody
	GetTimeCost() *int64
}

type ModifyImageRegistryResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// N/A
	//
	// example:
	//
	// N/A
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 379a9b8f-107b-4630-9e95-2299a1ea****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s ModifyImageRegistryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageRegistryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyImageRegistryResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ModifyImageRegistryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyImageRegistryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyImageRegistryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImageRegistryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyImageRegistryResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *ModifyImageRegistryResponseBody) SetCode(v string) *ModifyImageRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyImageRegistryResponseBody) SetData(v interface{}) *ModifyImageRegistryResponseBody {
	s.Data = v
	return s
}

func (s *ModifyImageRegistryResponseBody) SetHttpStatusCode(v int32) *ModifyImageRegistryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyImageRegistryResponseBody) SetMessage(v string) *ModifyImageRegistryResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyImageRegistryResponseBody) SetRequestId(v string) *ModifyImageRegistryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImageRegistryResponseBody) SetSuccess(v bool) *ModifyImageRegistryResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyImageRegistryResponseBody) SetTimeCost(v int64) *ModifyImageRegistryResponseBody {
	s.TimeCost = &v
	return s
}

func (s *ModifyImageRegistryResponseBody) Validate() error {
	return dara.Validate(s)
}
