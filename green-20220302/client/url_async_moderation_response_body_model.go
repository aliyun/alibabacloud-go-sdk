// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUrlAsyncModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UrlAsyncModerationResponseBody
	GetCode() *int32
	SetData(v *UrlAsyncModerationResponseBodyData) *UrlAsyncModerationResponseBody
	GetData() *UrlAsyncModerationResponseBodyData
	SetMsg(v string) *UrlAsyncModerationResponseBody
	GetMsg() *string
	SetRequestId(v string) *UrlAsyncModerationResponseBody
	GetRequestId() *string
}

type UrlAsyncModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *UrlAsyncModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6CF2815C-****-****-B52E-FF6E2****492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UrlAsyncModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UrlAsyncModerationResponseBody) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UrlAsyncModerationResponseBody) GetData() *UrlAsyncModerationResponseBodyData {
	return s.Data
}

func (s *UrlAsyncModerationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *UrlAsyncModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UrlAsyncModerationResponseBody) SetCode(v int32) *UrlAsyncModerationResponseBody {
	s.Code = &v
	return s
}

func (s *UrlAsyncModerationResponseBody) SetData(v *UrlAsyncModerationResponseBodyData) *UrlAsyncModerationResponseBody {
	s.Data = v
	return s
}

func (s *UrlAsyncModerationResponseBody) SetMsg(v string) *UrlAsyncModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *UrlAsyncModerationResponseBody) SetRequestId(v string) *UrlAsyncModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UrlAsyncModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UrlAsyncModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The reqId field returned by the Url Async Moderation API.
	//
	// example:
	//
	// A07B3DB9-D762-5C56-95B1-8EC55CF176D2
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s UrlAsyncModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UrlAsyncModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *UrlAsyncModerationResponseBodyData) GetReqId() *string {
	return s.ReqId
}

func (s *UrlAsyncModerationResponseBodyData) SetDataId(v string) *UrlAsyncModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *UrlAsyncModerationResponseBodyData) SetReqId(v string) *UrlAsyncModerationResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *UrlAsyncModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
