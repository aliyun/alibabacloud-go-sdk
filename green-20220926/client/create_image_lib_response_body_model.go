// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateImageLibResponseBody
	GetCode() *int32
	SetData(v *CreateImageLibResponseBodyData) *CreateImageLibResponseBody
	GetData() *CreateImageLibResponseBodyData
	SetHttpStatusCode(v int32) *CreateImageLibResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *CreateImageLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateImageLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImageLibResponseBody
	GetSuccess() *bool
}

type CreateImageLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateImageLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateImageLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateImageLibResponseBody) GetData() *CreateImageLibResponseBodyData {
	return s.Data
}

func (s *CreateImageLibResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateImageLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateImageLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageLibResponseBody) SetCode(v int32) *CreateImageLibResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageLibResponseBody) SetData(v *CreateImageLibResponseBodyData) *CreateImageLibResponseBody {
	s.Data = v
	return s
}

func (s *CreateImageLibResponseBody) SetHttpStatusCode(v int32) *CreateImageLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateImageLibResponseBody) SetMsg(v string) *CreateImageLibResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateImageLibResponseBody) SetRequestId(v string) *CreateImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageLibResponseBody) SetSuccess(v bool) *CreateImageLibResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageLibResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageLibResponseBodyData struct {
	// example:
	//
	// 138xxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
}

func (s CreateImageLibResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateImageLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateImageLibResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *CreateImageLibResponseBodyData) SetLibId(v string) *CreateImageLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *CreateImageLibResponseBodyData) Validate() error {
	return dara.Validate(s)
}
