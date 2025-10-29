// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackgroundPicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBackgroundPicResponseBody
	GetCode() *string
	SetData(v *CreateBackgroundPicResponseBodyData) *CreateBackgroundPicResponseBody
	GetData() *CreateBackgroundPicResponseBodyData
	SetHttpStatusCode(v int32) *CreateBackgroundPicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBackgroundPicResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBackgroundPicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBackgroundPicResponseBody
	GetSuccess() *bool
}

type CreateBackgroundPicResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateBackgroundPicResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateBackgroundPicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackgroundPicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackgroundPicResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBackgroundPicResponseBody) GetData() *CreateBackgroundPicResponseBodyData {
	return s.Data
}

func (s *CreateBackgroundPicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBackgroundPicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBackgroundPicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackgroundPicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBackgroundPicResponseBody) SetCode(v string) *CreateBackgroundPicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBackgroundPicResponseBody) SetData(v *CreateBackgroundPicResponseBodyData) *CreateBackgroundPicResponseBody {
	s.Data = v
	return s
}

func (s *CreateBackgroundPicResponseBody) SetHttpStatusCode(v int32) *CreateBackgroundPicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBackgroundPicResponseBody) SetMessage(v string) *CreateBackgroundPicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackgroundPicResponseBody) SetRequestId(v string) *CreateBackgroundPicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackgroundPicResponseBody) SetSuccess(v bool) *CreateBackgroundPicResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBackgroundPicResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBackgroundPicResponseBodyData struct {
	// example:
	//
	// M1lhKArheOyYdeYybDFqS1-Q
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateBackgroundPicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateBackgroundPicResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBackgroundPicResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateBackgroundPicResponseBodyData) SetId(v string) *CreateBackgroundPicResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateBackgroundPicResponseBodyData) Validate() error {
	return dara.Validate(s)
}
