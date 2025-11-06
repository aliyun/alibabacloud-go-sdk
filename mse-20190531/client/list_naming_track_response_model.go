// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamingTrackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNamingTrackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNamingTrackResponse
	GetStatusCode() *int32
	SetBody(v *ListNamingTrackResponseBody) *ListNamingTrackResponse
	GetBody() *ListNamingTrackResponseBody
}

type ListNamingTrackResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNamingTrackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNamingTrackResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNamingTrackResponse) GoString() string {
	return s.String()
}

func (s *ListNamingTrackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNamingTrackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNamingTrackResponse) GetBody() *ListNamingTrackResponseBody {
	return s.Body
}

func (s *ListNamingTrackResponse) SetHeaders(v map[string]*string) *ListNamingTrackResponse {
	s.Headers = v
	return s
}

func (s *ListNamingTrackResponse) SetStatusCode(v int32) *ListNamingTrackResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamingTrackResponse) SetBody(v *ListNamingTrackResponseBody) *ListNamingTrackResponse {
	s.Body = v
	return s
}

func (s *ListNamingTrackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
