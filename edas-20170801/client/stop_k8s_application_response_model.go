// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *StopK8sApplicationResponseBody) *StopK8sApplicationResponse
	GetBody() *StopK8sApplicationResponseBody
}

type StopK8sApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s StopK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *StopK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopK8sApplicationResponse) GetBody() *StopK8sApplicationResponseBody {
	return s.Body
}

func (s *StopK8sApplicationResponse) SetHeaders(v map[string]*string) *StopK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *StopK8sApplicationResponse) SetStatusCode(v int32) *StopK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopK8sApplicationResponse) SetBody(v *StopK8sApplicationResponseBody) *StopK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *StopK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
