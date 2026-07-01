// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDatasetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDatasetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDatasetResponseBody
	GetSuccess() *bool
}

type UpdateDatasetResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDatasetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDatasetResponseBody) SetCode(v string) *UpdateDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetHttpStatusCode(v int32) *UpdateDatasetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetMessage(v string) *UpdateDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetSuccess(v bool) *UpdateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
