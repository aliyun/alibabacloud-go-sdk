// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevokeDataServiceApiResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RevokeDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RevokeDataServiceApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeDataServiceApiResponseBody
	GetSuccess() *bool
}

type RevokeDataServiceApiResponseBody struct {
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

func (s RevokeDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeDataServiceApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RevokeDataServiceApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeDataServiceApiResponseBody) SetCode(v string) *RevokeDataServiceApiResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeDataServiceApiResponseBody) SetHttpStatusCode(v int32) *RevokeDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeDataServiceApiResponseBody) SetMessage(v string) *RevokeDataServiceApiResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeDataServiceApiResponseBody) SetRequestId(v string) *RevokeDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeDataServiceApiResponseBody) SetSuccess(v bool) *RevokeDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
