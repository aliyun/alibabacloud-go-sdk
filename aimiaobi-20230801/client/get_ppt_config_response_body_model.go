// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPptConfigResponseBody
	GetCode() *string
	SetData(v *GetPptConfigResponseBodyData) *GetPptConfigResponseBody
	GetData() *GetPptConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetPptConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPptConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPptConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPptConfigResponseBody
	GetSuccess() *bool
}

type GetPptConfigResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPptConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPptConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPptConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPptConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPptConfigResponseBody) GetData() *GetPptConfigResponseBodyData {
	return s.Data
}

func (s *GetPptConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPptConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPptConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPptConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPptConfigResponseBody) SetCode(v string) *GetPptConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetPptConfigResponseBody) SetData(v *GetPptConfigResponseBodyData) *GetPptConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetPptConfigResponseBody) SetHttpStatusCode(v int32) *GetPptConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPptConfigResponseBody) SetMessage(v string) *GetPptConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetPptConfigResponseBody) SetRequestId(v string) *GetPptConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPptConfigResponseBody) SetSuccess(v bool) *GetPptConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetPptConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPptConfigResponseBodyData struct {
	// AppKey
	//
	// example:
	//
	// 333444766
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Code
	//
	// example:
	//
	// r4wr5y6g
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPptConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPptConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPptConfigResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *GetPptConfigResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetPptConfigResponseBodyData) SetAppKey(v string) *GetPptConfigResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *GetPptConfigResponseBodyData) SetCode(v string) *GetPptConfigResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetPptConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
