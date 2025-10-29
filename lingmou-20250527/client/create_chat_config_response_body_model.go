// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateChatConfigResponseBody
	GetCode() *string
	SetData(v *CreateChatConfigResponseBodyData) *CreateChatConfigResponseBody
	GetData() *CreateChatConfigResponseBodyData
	SetHttpStatusCode(v int32) *CreateChatConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateChatConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateChatConfigResponseBody
	GetSuccess() *bool
}

type CreateChatConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateChatConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0EC3BA89-13F5-5766-A0BA-85096092A032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateChatConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatConfigResponseBody) GetData() *CreateChatConfigResponseBodyData {
	return s.Data
}

func (s *CreateChatConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateChatConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateChatConfigResponseBody) SetCode(v string) *CreateChatConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatConfigResponseBody) SetData(v *CreateChatConfigResponseBodyData) *CreateChatConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatConfigResponseBody) SetHttpStatusCode(v int32) *CreateChatConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateChatConfigResponseBody) SetMessage(v string) *CreateChatConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatConfigResponseBody) SetRequestId(v string) *CreateChatConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatConfigResponseBody) SetSuccess(v bool) *CreateChatConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateChatConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateChatConfigResponseBodyData struct {
	// example:
	//
	// C1RznvtlM-JO6HuPHqNC-Xxg
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateChatConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatConfigResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateChatConfigResponseBodyData) SetId(v string) *CreateChatConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateChatConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
