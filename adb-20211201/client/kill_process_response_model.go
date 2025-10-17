// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillProcessResponse
	GetStatusCode() *int32
	SetBody(v *KillProcessResponseBody) *KillProcessResponse
	GetBody() *KillProcessResponseBody
}

type KillProcessResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s KillProcessResponse) GoString() string {
	return s.String()
}

func (s *KillProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillProcessResponse) GetBody() *KillProcessResponseBody {
	return s.Body
}

func (s *KillProcessResponse) SetHeaders(v map[string]*string) *KillProcessResponse {
	s.Headers = v
	return s
}

func (s *KillProcessResponse) SetStatusCode(v int32) *KillProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *KillProcessResponse) SetBody(v *KillProcessResponseBody) *KillProcessResponse {
	s.Body = v
	return s
}

func (s *KillProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
