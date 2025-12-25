// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishStatusResponse
	GetStatusCode() *int32
	SetBody(v *PublishStatusResponseBody) *PublishStatusResponse
	GetBody() *PublishStatusResponseBody
}

type PublishStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishStatusResponse) GoString() string {
	return s.String()
}

func (s *PublishStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishStatusResponse) GetBody() *PublishStatusResponseBody {
	return s.Body
}

func (s *PublishStatusResponse) SetHeaders(v map[string]*string) *PublishStatusResponse {
	s.Headers = v
	return s
}

func (s *PublishStatusResponse) SetStatusCode(v int32) *PublishStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishStatusResponse) SetBody(v *PublishStatusResponseBody) *PublishStatusResponse {
	s.Body = v
	return s
}

func (s *PublishStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
