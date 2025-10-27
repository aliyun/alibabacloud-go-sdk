// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneypotProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneypotProbeResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneypotProbeResponseBody) *GetHoneypotProbeResponse
	GetBody() *GetHoneypotProbeResponseBody
}

type GetHoneypotProbeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneypotProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneypotProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeResponse) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneypotProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneypotProbeResponse) GetBody() *GetHoneypotProbeResponseBody {
	return s.Body
}

func (s *GetHoneypotProbeResponse) SetHeaders(v map[string]*string) *GetHoneypotProbeResponse {
	s.Headers = v
	return s
}

func (s *GetHoneypotProbeResponse) SetStatusCode(v int32) *GetHoneypotProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneypotProbeResponse) SetBody(v *GetHoneypotProbeResponseBody) *GetHoneypotProbeResponse {
	s.Body = v
	return s
}

func (s *GetHoneypotProbeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
