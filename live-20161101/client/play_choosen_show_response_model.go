// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayChoosenShowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PlayChoosenShowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PlayChoosenShowResponse
	GetStatusCode() *int32
	SetBody(v *PlayChoosenShowResponseBody) *PlayChoosenShowResponse
	GetBody() *PlayChoosenShowResponseBody
}

type PlayChoosenShowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PlayChoosenShowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PlayChoosenShowResponse) String() string {
	return dara.Prettify(s)
}

func (s PlayChoosenShowResponse) GoString() string {
	return s.String()
}

func (s *PlayChoosenShowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PlayChoosenShowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PlayChoosenShowResponse) GetBody() *PlayChoosenShowResponseBody {
	return s.Body
}

func (s *PlayChoosenShowResponse) SetHeaders(v map[string]*string) *PlayChoosenShowResponse {
	s.Headers = v
	return s
}

func (s *PlayChoosenShowResponse) SetStatusCode(v int32) *PlayChoosenShowResponse {
	s.StatusCode = &v
	return s
}

func (s *PlayChoosenShowResponse) SetBody(v *PlayChoosenShowResponseBody) *PlayChoosenShowResponse {
	s.Body = v
	return s
}

func (s *PlayChoosenShowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
