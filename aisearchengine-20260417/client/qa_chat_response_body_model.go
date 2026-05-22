// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQaChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QaChatResponseBodyData) *QaChatResponseBody
	GetData() *QaChatResponseBodyData
	SetEvent(v string) *QaChatResponseBody
	GetEvent() *string
	SetId(v string) *QaChatResponseBody
	GetId() *string
}

type QaChatResponseBody struct {
	Data *QaChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// start
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 2a127bc9-9474-405d-916d-8bc4475fa459
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s QaChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QaChatResponseBody) GoString() string {
	return s.String()
}

func (s *QaChatResponseBody) GetData() *QaChatResponseBodyData {
	return s.Data
}

func (s *QaChatResponseBody) GetEvent() *string {
	return s.Event
}

func (s *QaChatResponseBody) GetId() *string {
	return s.Id
}

func (s *QaChatResponseBody) SetData(v *QaChatResponseBodyData) *QaChatResponseBody {
	s.Data = v
	return s
}

func (s *QaChatResponseBody) SetEvent(v string) *QaChatResponseBody {
	s.Event = &v
	return s
}

func (s *QaChatResponseBody) SetId(v string) *QaChatResponseBody {
	s.Id = &v
	return s
}

func (s *QaChatResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QaChatResponseBodyData struct {
	// example:
	//
	// {
	//
	//   "items": [
	//
	//     {
	//
	//       "title": "MacBook Neo 深度测评：全面超越，最轻盈的生产力",
	//
	//       "videoUrl": "https://video.example.com/macbook-neo-review.mp4",
	//
	//       "coverUrl": "https://images.unsplash.com/photo-1529139574466-a303027c1d8b?auto=format&fit=crop&w=800&q=80",
	//
	//       "tags": [
	//
	//         "数码产品",
	//
	//         "评测",
	//
	//         "MacBook",
	//
	//         "生产力"
	//
	//       ]
	//
	//     }
	//
	//   ]
	//
	// }
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 如果你想更快看到上身效果，这 3 条短视频会更直观。\\n\\n**第一条：完整通勤 look**\\n\\n3 套通勤到周末无缝切换的浅色运动鞋穿搭。
	Delta *string `json:"delta,omitempty" xml:"delta,omitempty"`
	// example:
	//
	// 400
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 参数错误
	ErrorText *string `json:"errorText,omitempty" xml:"errorText,omitempty"`
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 2a127bc9-9474-405d-916d-8bc4475fa459
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2a127bc9-9474-405d-916d-8bc4475fa459
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
	// example:
	//
	// v
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QaChatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QaChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *QaChatResponseBodyData) GetData() *string {
	return s.Data
}

func (s *QaChatResponseBodyData) GetDelta() *string {
	return s.Delta
}

func (s *QaChatResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QaChatResponseBodyData) GetErrorText() *string {
	return s.ErrorText
}

func (s *QaChatResponseBodyData) GetFinishReason() *string {
	return s.FinishReason
}

func (s *QaChatResponseBodyData) GetId() *string {
	return s.Id
}

func (s *QaChatResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *QaChatResponseBodyData) GetRetryable() *bool {
	return s.Retryable
}

func (s *QaChatResponseBodyData) GetType() *string {
	return s.Type
}

func (s *QaChatResponseBodyData) SetData(v string) *QaChatResponseBodyData {
	s.Data = &v
	return s
}

func (s *QaChatResponseBodyData) SetDelta(v string) *QaChatResponseBodyData {
	s.Delta = &v
	return s
}

func (s *QaChatResponseBodyData) SetErrorCode(v string) *QaChatResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *QaChatResponseBodyData) SetErrorText(v string) *QaChatResponseBodyData {
	s.ErrorText = &v
	return s
}

func (s *QaChatResponseBodyData) SetFinishReason(v string) *QaChatResponseBodyData {
	s.FinishReason = &v
	return s
}

func (s *QaChatResponseBodyData) SetId(v string) *QaChatResponseBodyData {
	s.Id = &v
	return s
}

func (s *QaChatResponseBodyData) SetRequestId(v string) *QaChatResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *QaChatResponseBodyData) SetRetryable(v bool) *QaChatResponseBodyData {
	s.Retryable = &v
	return s
}

func (s *QaChatResponseBodyData) SetType(v string) *QaChatResponseBodyData {
	s.Type = &v
	return s
}

func (s *QaChatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
