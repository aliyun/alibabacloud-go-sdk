// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetsGovernObjectStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataAssetsGovernObjectStatusResponseBody
	GetCode() *string
	SetData(v int32) *UpdateDataAssetsGovernObjectStatusResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *UpdateDataAssetsGovernObjectStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDataAssetsGovernObjectStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataAssetsGovernObjectStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataAssetsGovernObjectStatusResponseBody
	GetSuccess() *bool
}

type UpdateDataAssetsGovernObjectStatusResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataAssetsGovernObjectStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetsGovernObjectStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) GetData() *int32 {
	return s.Data
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) SetCode(v string) *UpdateDataAssetsGovernObjectStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) SetData(v int32) *UpdateDataAssetsGovernObjectStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) SetHttpStatusCode(v int32) *UpdateDataAssetsGovernObjectStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) SetMessage(v string) *UpdateDataAssetsGovernObjectStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) SetRequestId(v string) *UpdateDataAssetsGovernObjectStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) SetSuccess(v bool) *UpdateDataAssetsGovernObjectStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
