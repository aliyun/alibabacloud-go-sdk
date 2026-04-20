// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChatMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChatMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChatMessageResponse
	GetStatusCode() *int32
	SetId(v string) *DescribeChatMessageResponse
	GetId() *string
	SetEvent(v string) *DescribeChatMessageResponse
	GetEvent() *string
	SetBody(v *DescribeChatMessageResponseBody) *DescribeChatMessageResponse
	GetBody() *DescribeChatMessageResponseBody
}

type DescribeChatMessageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                          `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                          `json:"event,omitempty" xml:"event,omitempty"`
	Body       *DescribeChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChatMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChatMessageResponse) GoString() string {
	return s.String()
}

func (s *DescribeChatMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChatMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChatMessageResponse) GetId() *string {
	return s.Id
}

func (s *DescribeChatMessageResponse) GetEvent() *string {
	return s.Event
}

func (s *DescribeChatMessageResponse) GetBody() *DescribeChatMessageResponseBody {
	return s.Body
}

func (s *DescribeChatMessageResponse) SetHeaders(v map[string]*string) *DescribeChatMessageResponse {
	s.Headers = v
	return s
}

func (s *DescribeChatMessageResponse) SetStatusCode(v int32) *DescribeChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChatMessageResponse) SetId(v string) *DescribeChatMessageResponse {
	s.Id = &v
	return s
}

func (s *DescribeChatMessageResponse) SetEvent(v string) *DescribeChatMessageResponse {
	s.Event = &v
	return s
}

func (s *DescribeChatMessageResponse) SetBody(v *DescribeChatMessageResponseBody) *DescribeChatMessageResponse {
	s.Body = v
	return s
}

func (s *DescribeChatMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
