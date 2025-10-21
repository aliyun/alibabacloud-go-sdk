// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicResponseBody) *GetTopicResponse
	GetBody() *GetTopicResponseBody
}

type GetTopicResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponse) GoString() string {
	return s.String()
}

func (s *GetTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicResponse) GetBody() *GetTopicResponseBody {
	return s.Body
}

func (s *GetTopicResponse) SetHeaders(v map[string]*string) *GetTopicResponse {
	s.Headers = v
	return s
}

func (s *GetTopicResponse) SetStatusCode(v int32) *GetTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicResponse) SetBody(v *GetTopicResponseBody) *GetTopicResponse {
	s.Body = v
	return s
}

func (s *GetTopicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
