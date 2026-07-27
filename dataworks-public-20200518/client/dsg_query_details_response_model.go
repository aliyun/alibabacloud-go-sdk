// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgQueryDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgQueryDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DsgQueryDetailsResponseBody) *DsgQueryDetailsResponse
	GetBody() *DsgQueryDetailsResponseBody
}

type DsgQueryDetailsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgQueryDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgQueryDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDetailsResponse) GoString() string {
	return s.String()
}

func (s *DsgQueryDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgQueryDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgQueryDetailsResponse) GetBody() *DsgQueryDetailsResponseBody {
	return s.Body
}

func (s *DsgQueryDetailsResponse) SetHeaders(v map[string]*string) *DsgQueryDetailsResponse {
	s.Headers = v
	return s
}

func (s *DsgQueryDetailsResponse) SetStatusCode(v int32) *DsgQueryDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgQueryDetailsResponse) SetBody(v *DsgQueryDetailsResponseBody) *DsgQueryDetailsResponse {
	s.Body = v
	return s
}

func (s *DsgQueryDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
