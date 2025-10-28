// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sSlbResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sSlbResponseBody) *UpdateK8sSlbResponse
	GetBody() *UpdateK8sSlbResponseBody
}

type UpdateK8sSlbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sSlbResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sSlbResponse) GetBody() *UpdateK8sSlbResponseBody {
	return s.Body
}

func (s *UpdateK8sSlbResponse) SetHeaders(v map[string]*string) *UpdateK8sSlbResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sSlbResponse) SetStatusCode(v int32) *UpdateK8sSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sSlbResponse) SetBody(v *UpdateK8sSlbResponseBody) *UpdateK8sSlbResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
