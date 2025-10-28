// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *RestartK8sApplicationResponseBody) *RestartK8sApplicationResponse
	GetBody() *RestartK8sApplicationResponseBody
}

type RestartK8sApplicationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *RestartK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartK8sApplicationResponse) GetBody() *RestartK8sApplicationResponseBody {
	return s.Body
}

func (s *RestartK8sApplicationResponse) SetHeaders(v map[string]*string) *RestartK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *RestartK8sApplicationResponse) SetStatusCode(v int32) *RestartK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartK8sApplicationResponse) SetBody(v *RestartK8sApplicationResponseBody) *RestartK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *RestartK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
