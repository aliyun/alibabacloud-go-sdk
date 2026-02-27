// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataServiceAppGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDataServiceAppGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDataServiceAppGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataServiceAppGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataServiceAppGroupResponseBody
	GetSuccess() *bool
}

type UpdateDataServiceAppGroupResponseBody struct {
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
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataServiceAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataServiceAppGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataServiceAppGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataServiceAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataServiceAppGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataServiceAppGroupResponseBody) SetCode(v string) *UpdateDataServiceAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataServiceAppGroupResponseBody) SetHttpStatusCode(v int32) *UpdateDataServiceAppGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataServiceAppGroupResponseBody) SetMessage(v string) *UpdateDataServiceAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataServiceAppGroupResponseBody) SetRequestId(v string) *UpdateDataServiceAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataServiceAppGroupResponseBody) SetSuccess(v bool) *UpdateDataServiceAppGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataServiceAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
