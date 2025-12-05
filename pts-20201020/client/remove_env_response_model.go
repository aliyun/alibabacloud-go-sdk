// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEnvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveEnvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveEnvResponse
	GetStatusCode() *int32
	SetBody(v *RemoveEnvResponseBody) *RemoveEnvResponse
	GetBody() *RemoveEnvResponseBody
}

type RemoveEnvResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveEnvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveEnvResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveEnvResponse) GoString() string {
	return s.String()
}

func (s *RemoveEnvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveEnvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveEnvResponse) GetBody() *RemoveEnvResponseBody {
	return s.Body
}

func (s *RemoveEnvResponse) SetHeaders(v map[string]*string) *RemoveEnvResponse {
	s.Headers = v
	return s
}

func (s *RemoveEnvResponse) SetStatusCode(v int32) *RemoveEnvResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEnvResponse) SetBody(v *RemoveEnvResponseBody) *RemoveEnvResponse {
	s.Body = v
	return s
}

func (s *RemoveEnvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
