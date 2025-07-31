// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GrantDataServiceApiResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GrantDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GrantDataServiceApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GrantDataServiceApiResponseBody
	GetSuccess() *bool
}

type GrantDataServiceApiResponseBody struct {
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

func (s GrantDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *GrantDataServiceApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *GrantDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GrantDataServiceApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantDataServiceApiResponseBody) SetCode(v string) *GrantDataServiceApiResponseBody {
	s.Code = &v
	return s
}

func (s *GrantDataServiceApiResponseBody) SetHttpStatusCode(v int32) *GrantDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantDataServiceApiResponseBody) SetMessage(v string) *GrantDataServiceApiResponseBody {
	s.Message = &v
	return s
}

func (s *GrantDataServiceApiResponseBody) SetRequestId(v string) *GrantDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantDataServiceApiResponseBody) SetSuccess(v bool) *GrantDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *GrantDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
