// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryYuqingMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryYuqingMessageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryYuqingMessageResponseBody
	GetTotalCount() *int64
	SetYuqingMessages(v []*YuqingMessage) *QueryYuqingMessageResponseBody
	GetYuqingMessages() []*YuqingMessage
}

type QueryYuqingMessageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// adacae47-6fc0-45c6-897c-26201aefbdfd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount     *int64           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	YuqingMessages []*YuqingMessage `json:"yuqingMessages,omitempty" xml:"yuqingMessages,omitempty" type:"Repeated"`
}

func (s QueryYuqingMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryYuqingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryYuqingMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryYuqingMessageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryYuqingMessageResponseBody) GetYuqingMessages() []*YuqingMessage {
	return s.YuqingMessages
}

func (s *QueryYuqingMessageResponseBody) SetRequestId(v string) *QueryYuqingMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryYuqingMessageResponseBody) SetTotalCount(v int64) *QueryYuqingMessageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryYuqingMessageResponseBody) SetYuqingMessages(v []*YuqingMessage) *QueryYuqingMessageResponseBody {
	s.YuqingMessages = v
	return s
}

func (s *QueryYuqingMessageResponseBody) Validate() error {
	if s.YuqingMessages != nil {
		for _, item := range s.YuqingMessages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
