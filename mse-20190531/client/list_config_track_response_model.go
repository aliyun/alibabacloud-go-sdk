// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigTrackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigTrackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigTrackResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigTrackResponseBody) *ListConfigTrackResponse
	GetBody() *ListConfigTrackResponseBody
}

type ListConfigTrackResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigTrackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigTrackResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTrackResponse) GoString() string {
	return s.String()
}

func (s *ListConfigTrackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigTrackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigTrackResponse) GetBody() *ListConfigTrackResponseBody {
	return s.Body
}

func (s *ListConfigTrackResponse) SetHeaders(v map[string]*string) *ListConfigTrackResponse {
	s.Headers = v
	return s
}

func (s *ListConfigTrackResponse) SetStatusCode(v int32) *ListConfigTrackResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigTrackResponse) SetBody(v *ListConfigTrackResponseBody) *ListConfigTrackResponse {
	s.Body = v
	return s
}

func (s *ListConfigTrackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
