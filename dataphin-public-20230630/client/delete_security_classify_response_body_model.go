// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecurityClassifyResponseBody
	GetCode() *string
	SetData(v bool) *DeleteSecurityClassifyResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteSecurityClassifyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSecurityClassifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecurityClassifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSecurityClassifyResponseBody
	GetSuccess() *bool
}

type DeleteSecurityClassifyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s DeleteSecurityClassifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecurityClassifyResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteSecurityClassifyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSecurityClassifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecurityClassifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityClassifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecurityClassifyResponseBody) SetCode(v string) *DeleteSecurityClassifyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecurityClassifyResponseBody) SetData(v bool) *DeleteSecurityClassifyResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSecurityClassifyResponseBody) SetHttpStatusCode(v int32) *DeleteSecurityClassifyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSecurityClassifyResponseBody) SetMessage(v string) *DeleteSecurityClassifyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecurityClassifyResponseBody) SetRequestId(v string) *DeleteSecurityClassifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityClassifyResponseBody) SetSuccess(v bool) *DeleteSecurityClassifyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecurityClassifyResponseBody) Validate() error {
	return dara.Validate(s)
}
