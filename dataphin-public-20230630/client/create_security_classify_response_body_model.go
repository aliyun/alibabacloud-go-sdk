// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSecurityClassifyResponseBody
	GetCode() *string
	SetData(v int64) *CreateSecurityClassifyResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateSecurityClassifyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSecurityClassifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecurityClassifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSecurityClassifyResponseBody
	GetSuccess() *bool
}

type CreateSecurityClassifyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateSecurityClassifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecurityClassifyResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateSecurityClassifyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSecurityClassifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecurityClassifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityClassifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSecurityClassifyResponseBody) SetCode(v string) *CreateSecurityClassifyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecurityClassifyResponseBody) SetData(v int64) *CreateSecurityClassifyResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSecurityClassifyResponseBody) SetHttpStatusCode(v int32) *CreateSecurityClassifyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSecurityClassifyResponseBody) SetMessage(v string) *CreateSecurityClassifyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecurityClassifyResponseBody) SetRequestId(v string) *CreateSecurityClassifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityClassifyResponseBody) SetSuccess(v bool) *CreateSecurityClassifyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecurityClassifyResponseBody) Validate() error {
	return dara.Validate(s)
}
