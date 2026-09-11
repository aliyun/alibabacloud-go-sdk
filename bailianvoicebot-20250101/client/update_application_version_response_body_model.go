// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApplicationVersionResponseBody
	GetCode() *string
	SetData(v string) *UpdateApplicationVersionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateApplicationVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateApplicationVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationVersionResponseBody
	GetRequestId() *string
}

type UpdateApplicationVersionResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApplicationVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateApplicationVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateApplicationVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationVersionResponseBody) SetCode(v string) *UpdateApplicationVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationVersionResponseBody) SetData(v string) *UpdateApplicationVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateApplicationVersionResponseBody) SetHttpStatusCode(v int32) *UpdateApplicationVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateApplicationVersionResponseBody) SetMessage(v string) *UpdateApplicationVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationVersionResponseBody) SetRequestId(v string) *UpdateApplicationVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
