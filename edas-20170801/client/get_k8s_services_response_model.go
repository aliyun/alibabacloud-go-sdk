// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetK8sServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetK8sServicesResponse
	GetStatusCode() *int32
	SetBody(v *GetK8sServicesResponseBody) *GetK8sServicesResponse
	GetBody() *GetK8sServicesResponseBody
}

type GetK8sServicesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetK8sServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetK8sServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetK8sServicesResponse) GoString() string {
	return s.String()
}

func (s *GetK8sServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetK8sServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetK8sServicesResponse) GetBody() *GetK8sServicesResponseBody {
	return s.Body
}

func (s *GetK8sServicesResponse) SetHeaders(v map[string]*string) *GetK8sServicesResponse {
	s.Headers = v
	return s
}

func (s *GetK8sServicesResponse) SetStatusCode(v int32) *GetK8sServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetK8sServicesResponse) SetBody(v *GetK8sServicesResponseBody) *GetK8sServicesResponse {
	s.Body = v
	return s
}

func (s *GetK8sServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
