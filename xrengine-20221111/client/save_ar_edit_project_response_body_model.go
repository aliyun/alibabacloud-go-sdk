// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveArEditProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveArEditProjectResponseBody
	GetCode() *string
	SetData(v bool) *SaveArEditProjectResponseBody
	GetData() *bool
	SetErrorName(v string) *SaveArEditProjectResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *SaveArEditProjectResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *SaveArEditProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveArEditProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveArEditProjectResponseBody
	GetSuccess() *bool
}

type SaveArEditProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveArEditProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveArEditProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SaveArEditProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveArEditProjectResponseBody) GetData() *bool {
	return s.Data
}

func (s *SaveArEditProjectResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *SaveArEditProjectResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SaveArEditProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveArEditProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveArEditProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveArEditProjectResponseBody) SetCode(v string) *SaveArEditProjectResponseBody {
	s.Code = &v
	return s
}

func (s *SaveArEditProjectResponseBody) SetData(v bool) *SaveArEditProjectResponseBody {
	s.Data = &v
	return s
}

func (s *SaveArEditProjectResponseBody) SetErrorName(v string) *SaveArEditProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *SaveArEditProjectResponseBody) SetHttpCode(v int32) *SaveArEditProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SaveArEditProjectResponseBody) SetMessage(v string) *SaveArEditProjectResponseBody {
	s.Message = &v
	return s
}

func (s *SaveArEditProjectResponseBody) SetRequestId(v string) *SaveArEditProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveArEditProjectResponseBody) SetSuccess(v bool) *SaveArEditProjectResponseBody {
	s.Success = &v
	return s
}

func (s *SaveArEditProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
