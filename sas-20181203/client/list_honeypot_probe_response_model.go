// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotProbeResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotProbeResponseBody) *ListHoneypotProbeResponse
	GetBody() *ListHoneypotProbeResponseBody
}

type ListHoneypotProbeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotProbeResponse) GetBody() *ListHoneypotProbeResponseBody {
	return s.Body
}

func (s *ListHoneypotProbeResponse) SetHeaders(v map[string]*string) *ListHoneypotProbeResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotProbeResponse) SetStatusCode(v int32) *ListHoneypotProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotProbeResponse) SetBody(v *ListHoneypotProbeResponseBody) *ListHoneypotProbeResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotProbeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
