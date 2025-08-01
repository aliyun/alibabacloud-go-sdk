// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServiceSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddServiceSourceResponseBody
	GetCode() *int32
	SetData(v int64) *AddServiceSourceResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddServiceSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddServiceSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddServiceSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddServiceSourceResponseBody
	GetSuccess() *bool
}

type AddServiceSourceResponseBody struct {
	// duplicatedClusterAliasName
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// mse-100-007
	//
	// example:
	//
	// 63
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// mse-200-105
	//
	// example:
	//
	// You are not authorized to perform this operation. Action: mse:AddServiceSource, Resource: acs:mse:cn-hangzhou:1105471854403716:instance/gw-082c943a8c304e48a37a7a29a5ddeda7
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The response data.
	//
	// example:
	//
	// D0DB055C-51F2-5BB2-82A6-CD8A3C6EE6BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// duplicated cluster alias name
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddServiceSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddServiceSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddServiceSourceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddServiceSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddServiceSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddServiceSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddServiceSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddServiceSourceResponseBody) SetCode(v int32) *AddServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetData(v int64) *AddServiceSourceResponseBody {
	s.Data = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetHttpStatusCode(v int32) *AddServiceSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetMessage(v string) *AddServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetRequestId(v string) *AddServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetSuccess(v bool) *AddServiceSourceResponseBody {
	s.Success = &v
	return s
}

func (s *AddServiceSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
