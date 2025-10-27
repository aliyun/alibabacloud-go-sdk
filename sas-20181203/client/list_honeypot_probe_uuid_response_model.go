// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotProbeUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotProbeUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotProbeUuidResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotProbeUuidResponseBody) *ListHoneypotProbeUuidResponse
	GetBody() *ListHoneypotProbeUuidResponseBody
}

type ListHoneypotProbeUuidResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotProbeUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotProbeUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeUuidResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotProbeUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotProbeUuidResponse) GetBody() *ListHoneypotProbeUuidResponseBody {
	return s.Body
}

func (s *ListHoneypotProbeUuidResponse) SetHeaders(v map[string]*string) *ListHoneypotProbeUuidResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotProbeUuidResponse) SetStatusCode(v int32) *ListHoneypotProbeUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotProbeUuidResponse) SetBody(v *ListHoneypotProbeUuidResponseBody) *ListHoneypotProbeUuidResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotProbeUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
