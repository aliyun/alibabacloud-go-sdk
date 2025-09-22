// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ChatItem) *ChatResponseBody
	GetData() *ChatItem
	SetRequestId(v string) *ChatResponseBody
	GetRequestId() *string
}

type ChatResponseBody struct {
	// Details of the Q&A.
	//
	// example:
	//
	// true
	Data *ChatItem `json:"data,omitempty" xml:"data,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) GetData() *ChatItem {
	return s.Data
}

func (s *ChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBody) SetData(v *ChatItem) *ChatResponseBody {
	s.Data = v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) Validate() error {
	return dara.Validate(s)
}
