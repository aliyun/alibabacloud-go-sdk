// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCredentialResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCredentialResponseBody
	GetSuccess() *bool
}

type DeleteCredentialResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCredentialResponseBody) SetCode(v string) *DeleteCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetHttpStatusCode(v int32) *DeleteCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetMessage(v string) *DeleteCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetRequestId(v string) *DeleteCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetSuccess(v bool) *DeleteCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
