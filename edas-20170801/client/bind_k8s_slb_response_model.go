// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindK8sSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindK8sSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindK8sSlbResponse
	GetStatusCode() *int32
	SetBody(v *BindK8sSlbResponseBody) *BindK8sSlbResponse
	GetBody() *BindK8sSlbResponseBody
}

type BindK8sSlbResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindK8sSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindK8sSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s BindK8sSlbResponse) GoString() string {
	return s.String()
}

func (s *BindK8sSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindK8sSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindK8sSlbResponse) GetBody() *BindK8sSlbResponseBody {
	return s.Body
}

func (s *BindK8sSlbResponse) SetHeaders(v map[string]*string) *BindK8sSlbResponse {
	s.Headers = v
	return s
}

func (s *BindK8sSlbResponse) SetStatusCode(v int32) *BindK8sSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *BindK8sSlbResponse) SetBody(v *BindK8sSlbResponseBody) *BindK8sSlbResponse {
	s.Body = v
	return s
}

func (s *BindK8sSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
