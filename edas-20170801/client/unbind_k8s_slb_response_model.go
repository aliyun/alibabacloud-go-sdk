// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindK8sSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindK8sSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindK8sSlbResponse
	GetStatusCode() *int32
	SetBody(v *UnbindK8sSlbResponseBody) *UnbindK8sSlbResponse
	GetBody() *UnbindK8sSlbResponseBody
}

type UnbindK8sSlbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindK8sSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindK8sSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindK8sSlbResponse) GoString() string {
	return s.String()
}

func (s *UnbindK8sSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindK8sSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindK8sSlbResponse) GetBody() *UnbindK8sSlbResponseBody {
	return s.Body
}

func (s *UnbindK8sSlbResponse) SetHeaders(v map[string]*string) *UnbindK8sSlbResponse {
	s.Headers = v
	return s
}

func (s *UnbindK8sSlbResponse) SetStatusCode(v int32) *UnbindK8sSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindK8sSlbResponse) SetBody(v *UnbindK8sSlbResponseBody) *UnbindK8sSlbResponse {
	s.Body = v
	return s
}

func (s *UnbindK8sSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
