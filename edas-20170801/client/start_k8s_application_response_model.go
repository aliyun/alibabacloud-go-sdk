// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *StartK8sApplicationResponseBody) *StartK8sApplicationResponse
	GetBody() *StartK8sApplicationResponseBody
}

type StartK8sApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *StartK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartK8sApplicationResponse) GetBody() *StartK8sApplicationResponseBody {
	return s.Body
}

func (s *StartK8sApplicationResponse) SetHeaders(v map[string]*string) *StartK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *StartK8sApplicationResponse) SetStatusCode(v int32) *StartK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartK8sApplicationResponse) SetBody(v *StartK8sApplicationResponseBody) *StartK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *StartK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
