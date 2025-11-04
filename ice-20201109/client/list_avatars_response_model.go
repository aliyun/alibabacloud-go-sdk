// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvatarsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvatarsResponse
	GetStatusCode() *int32
	SetBody(v *ListAvatarsResponseBody) *ListAvatarsResponse
	GetBody() *ListAvatarsResponseBody
}

type ListAvatarsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvatarsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvatarsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarsResponse) GoString() string {
	return s.String()
}

func (s *ListAvatarsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvatarsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvatarsResponse) GetBody() *ListAvatarsResponseBody {
	return s.Body
}

func (s *ListAvatarsResponse) SetHeaders(v map[string]*string) *ListAvatarsResponse {
	s.Headers = v
	return s
}

func (s *ListAvatarsResponse) SetStatusCode(v int32) *ListAvatarsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvatarsResponse) SetBody(v *ListAvatarsResponseBody) *ListAvatarsResponse {
	s.Body = v
	return s
}

func (s *ListAvatarsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
