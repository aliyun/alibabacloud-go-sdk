// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImages2LibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddImages2LibResponseBody
	GetCode() *int32
	SetData(v *AddImages2LibResponseBodyData) *AddImages2LibResponseBody
	GetData() *AddImages2LibResponseBodyData
	SetHttpStatusCode(v int32) *AddImages2LibResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *AddImages2LibResponseBody
	GetMsg() *string
	SetRequestId(v string) *AddImages2LibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddImages2LibResponseBody
	GetSuccess() *bool
}

type AddImages2LibResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *AddImages2LibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImages2LibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImages2LibResponseBody) GoString() string {
	return s.String()
}

func (s *AddImages2LibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddImages2LibResponseBody) GetData() *AddImages2LibResponseBodyData {
	return s.Data
}

func (s *AddImages2LibResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddImages2LibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *AddImages2LibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImages2LibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddImages2LibResponseBody) SetCode(v int32) *AddImages2LibResponseBody {
	s.Code = &v
	return s
}

func (s *AddImages2LibResponseBody) SetData(v *AddImages2LibResponseBodyData) *AddImages2LibResponseBody {
	s.Data = v
	return s
}

func (s *AddImages2LibResponseBody) SetHttpStatusCode(v int32) *AddImages2LibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddImages2LibResponseBody) SetMsg(v string) *AddImages2LibResponseBody {
	s.Msg = &v
	return s
}

func (s *AddImages2LibResponseBody) SetRequestId(v string) *AddImages2LibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImages2LibResponseBody) SetSuccess(v bool) *AddImages2LibResponseBody {
	s.Success = &v
	return s
}

func (s *AddImages2LibResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddImages2LibResponseBodyData struct {
	// The id of the uploaded image.
	//
	// example:
	//
	// 100001
	ImgId *string `json:"ImgId,omitempty" xml:"ImgId,omitempty"`
}

func (s AddImages2LibResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddImages2LibResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddImages2LibResponseBodyData) GetImgId() *string {
	return s.ImgId
}

func (s *AddImages2LibResponseBodyData) SetImgId(v string) *AddImages2LibResponseBodyData {
	s.ImgId = &v
	return s
}

func (s *AddImages2LibResponseBodyData) Validate() error {
	return dara.Validate(s)
}
