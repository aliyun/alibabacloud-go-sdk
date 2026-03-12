// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrayPublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrayPublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrayPublishResponse
	GetStatusCode() *int32
	SetBody(v *GrayPublishResponseBody) *GrayPublishResponse
	GetBody() *GrayPublishResponseBody
}

type GrayPublishResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrayPublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrayPublishResponse) String() string {
	return dara.Prettify(s)
}

func (s GrayPublishResponse) GoString() string {
	return s.String()
}

func (s *GrayPublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrayPublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrayPublishResponse) GetBody() *GrayPublishResponseBody {
	return s.Body
}

func (s *GrayPublishResponse) SetHeaders(v map[string]*string) *GrayPublishResponse {
	s.Headers = v
	return s
}

func (s *GrayPublishResponse) SetStatusCode(v int32) *GrayPublishResponse {
	s.StatusCode = &v
	return s
}

func (s *GrayPublishResponse) SetBody(v *GrayPublishResponseBody) *GrayPublishResponse {
	s.Body = v
	return s
}

func (s *GrayPublishResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
