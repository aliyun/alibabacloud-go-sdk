// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastStickerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBroadcastStickerResponseBody
	GetCode() *string
	SetData(v *CreateBroadcastStickerResponseBodyData) *CreateBroadcastStickerResponseBody
	GetData() *CreateBroadcastStickerResponseBodyData
	SetHttpStatusCode(v int32) *CreateBroadcastStickerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBroadcastStickerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBroadcastStickerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBroadcastStickerResponseBody
	GetSuccess() *bool
}

type CreateBroadcastStickerResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateBroadcastStickerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
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

func (s CreateBroadcastStickerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastStickerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBroadcastStickerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBroadcastStickerResponseBody) GetData() *CreateBroadcastStickerResponseBodyData {
	return s.Data
}

func (s *CreateBroadcastStickerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBroadcastStickerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBroadcastStickerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBroadcastStickerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBroadcastStickerResponseBody) SetCode(v string) *CreateBroadcastStickerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBroadcastStickerResponseBody) SetData(v *CreateBroadcastStickerResponseBodyData) *CreateBroadcastStickerResponseBody {
	s.Data = v
	return s
}

func (s *CreateBroadcastStickerResponseBody) SetHttpStatusCode(v int32) *CreateBroadcastStickerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBroadcastStickerResponseBody) SetMessage(v string) *CreateBroadcastStickerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBroadcastStickerResponseBody) SetRequestId(v string) *CreateBroadcastStickerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBroadcastStickerResponseBody) SetSuccess(v bool) *CreateBroadcastStickerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBroadcastStickerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBroadcastStickerResponseBodyData struct {
	// example:
	//
	// M1lhKArheOyYdeYybDFqS1-Q
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateBroadcastStickerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastStickerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBroadcastStickerResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateBroadcastStickerResponseBodyData) SetId(v string) *CreateBroadcastStickerResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateBroadcastStickerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
