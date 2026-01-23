// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishStandardResponse
	GetStatusCode() *int32
	SetBody(v *PublishStandardResponseBody) *PublishStandardResponse
	GetBody() *PublishStandardResponseBody
}

type PublishStandardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishStandardResponse) GoString() string {
	return s.String()
}

func (s *PublishStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishStandardResponse) GetBody() *PublishStandardResponseBody {
	return s.Body
}

func (s *PublishStandardResponse) SetHeaders(v map[string]*string) *PublishStandardResponse {
	s.Headers = v
	return s
}

func (s *PublishStandardResponse) SetStatusCode(v int32) *PublishStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishStandardResponse) SetBody(v *PublishStandardResponseBody) *PublishStandardResponse {
	s.Body = v
	return s
}

func (s *PublishStandardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
