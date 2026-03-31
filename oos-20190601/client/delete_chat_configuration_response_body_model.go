// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DeleteChatConfigurationResponseBody
	GetContent() *string
	SetRequestId(v string) *DeleteChatConfigurationResponseBody
	GetRequestId() *string
}

type DeleteChatConfigurationResponseBody struct {
	// example:
	//
	// Chat configuration deleted successfully.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 4DB0****1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChatConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatConfigurationResponseBody) GetContent() *string {
	return s.Content
}

func (s *DeleteChatConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatConfigurationResponseBody) SetContent(v string) *DeleteChatConfigurationResponseBody {
	s.Content = &v
	return s
}

func (s *DeleteChatConfigurationResponseBody) SetRequestId(v string) *DeleteChatConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
