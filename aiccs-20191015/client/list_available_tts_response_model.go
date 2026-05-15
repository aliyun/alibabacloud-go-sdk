// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableTtsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableTtsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableTtsResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableTtsResponseBody) *ListAvailableTtsResponse
	GetBody() *ListAvailableTtsResponseBody
}

type ListAvailableTtsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableTtsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableTtsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableTtsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableTtsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableTtsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableTtsResponse) GetBody() *ListAvailableTtsResponseBody {
	return s.Body
}

func (s *ListAvailableTtsResponse) SetHeaders(v map[string]*string) *ListAvailableTtsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableTtsResponse) SetStatusCode(v int32) *ListAvailableTtsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableTtsResponse) SetBody(v *ListAvailableTtsResponseBody) *ListAvailableTtsResponse {
	s.Body = v
	return s
}

func (s *ListAvailableTtsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
