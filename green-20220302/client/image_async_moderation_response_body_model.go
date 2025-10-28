// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageAsyncModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImageAsyncModerationResponseBody
	GetCode() *int32
	SetData(v *ImageAsyncModerationResponseBodyData) *ImageAsyncModerationResponseBody
	GetData() *ImageAsyncModerationResponseBodyData
	SetMsg(v string) *ImageAsyncModerationResponseBody
	GetMsg() *string
	SetRequestId(v string) *ImageAsyncModerationResponseBody
	GetRequestId() *string
}

type ImageAsyncModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ImageAsyncModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A926AE2-4C96-573F-824F-0532960799F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageAsyncModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageAsyncModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImageAsyncModerationResponseBody) GetData() *ImageAsyncModerationResponseBodyData {
	return s.Data
}

func (s *ImageAsyncModerationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ImageAsyncModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageAsyncModerationResponseBody) SetCode(v int32) *ImageAsyncModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageAsyncModerationResponseBody) SetData(v *ImageAsyncModerationResponseBodyData) *ImageAsyncModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageAsyncModerationResponseBody) SetMsg(v string) *ImageAsyncModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageAsyncModerationResponseBody) SetRequestId(v string) *ImageAsyncModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageAsyncModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageAsyncModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// fb5ffab1-993b-449f-b8d6-b97d5e3331f2
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The reqId field returned by the Image Async Moderation API. You can use this field to query the detection results.
	//
	// example:
	//
	// A07B3DB9-D762-5C56-95B1-8EC55CF176D2
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s ImageAsyncModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageAsyncModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ImageAsyncModerationResponseBodyData) GetReqId() *string {
	return s.ReqId
}

func (s *ImageAsyncModerationResponseBodyData) SetDataId(v string) *ImageAsyncModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageAsyncModerationResponseBodyData) SetReqId(v string) *ImageAsyncModerationResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *ImageAsyncModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
