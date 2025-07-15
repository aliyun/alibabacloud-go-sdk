// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneratedContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGeneratedContentResponseBody
	GetCode() *string
	SetData(v bool) *UpdateGeneratedContentResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateGeneratedContentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGeneratedContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGeneratedContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGeneratedContentResponseBody
	GetSuccess() *bool
}

type UpdateGeneratedContentResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGeneratedContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGeneratedContentResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGeneratedContentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGeneratedContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGeneratedContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGeneratedContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGeneratedContentResponseBody) SetCode(v string) *UpdateGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetData(v bool) *UpdateGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetHttpStatusCode(v int32) *UpdateGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetMessage(v string) *UpdateGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetRequestId(v string) *UpdateGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetSuccess(v bool) *UpdateGeneratedContentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) Validate() error {
	return dara.Validate(s)
}
