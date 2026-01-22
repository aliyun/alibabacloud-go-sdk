// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentChatRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarAgentChatRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarAgentChatRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarAgentChatRecordsResponseBody) *DescribePolarAgentChatRecordsResponse
	GetBody() *DescribePolarAgentChatRecordsResponseBody
}

type DescribePolarAgentChatRecordsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarAgentChatRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarAgentChatRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentChatRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentChatRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarAgentChatRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarAgentChatRecordsResponse) GetBody() *DescribePolarAgentChatRecordsResponseBody {
	return s.Body
}

func (s *DescribePolarAgentChatRecordsResponse) SetHeaders(v map[string]*string) *DescribePolarAgentChatRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarAgentChatRecordsResponse) SetStatusCode(v int32) *DescribePolarAgentChatRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarAgentChatRecordsResponse) SetBody(v *DescribePolarAgentChatRecordsResponseBody) *DescribePolarAgentChatRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarAgentChatRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
