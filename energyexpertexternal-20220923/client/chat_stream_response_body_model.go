// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ChatItem) *ChatStreamResponseBody
	GetData() *ChatItem
	SetRequestId(v string) *ChatStreamResponseBody
	GetRequestId() *string
}

type ChatStreamResponseBody struct {
	// Q&A content.
	Data *ChatItem `json:"data,omitempty" xml:"data,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ChatStreamResponseBody) GetData() *ChatItem {
	return s.Data
}

func (s *ChatStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatStreamResponseBody) SetData(v *ChatItem) *ChatStreamResponseBody {
	s.Data = v
	return s
}

func (s *ChatStreamResponseBody) SetRequestId(v string) *ChatStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatStreamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
