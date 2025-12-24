// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *PublishPlaybookResponseBody) *PublishPlaybookResponse
	GetBody() *PublishPlaybookResponseBody
}

type PublishPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishPlaybookResponse) GoString() string {
	return s.String()
}

func (s *PublishPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishPlaybookResponse) GetBody() *PublishPlaybookResponseBody {
	return s.Body
}

func (s *PublishPlaybookResponse) SetHeaders(v map[string]*string) *PublishPlaybookResponse {
	s.Headers = v
	return s
}

func (s *PublishPlaybookResponse) SetStatusCode(v int32) *PublishPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishPlaybookResponse) SetBody(v *PublishPlaybookResponseBody) *PublishPlaybookResponse {
	s.Body = v
	return s
}

func (s *PublishPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
