// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sConfigMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateK8sConfigMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateK8sConfigMapResponse
	GetStatusCode() *int32
	SetBody(v *CreateK8sConfigMapResponseBody) *CreateK8sConfigMapResponse
	GetBody() *CreateK8sConfigMapResponseBody
}

type CreateK8sConfigMapResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateK8sConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateK8sConfigMapResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sConfigMapResponse) GoString() string {
	return s.String()
}

func (s *CreateK8sConfigMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateK8sConfigMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateK8sConfigMapResponse) GetBody() *CreateK8sConfigMapResponseBody {
	return s.Body
}

func (s *CreateK8sConfigMapResponse) SetHeaders(v map[string]*string) *CreateK8sConfigMapResponse {
	s.Headers = v
	return s
}

func (s *CreateK8sConfigMapResponse) SetStatusCode(v int32) *CreateK8sConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateK8sConfigMapResponse) SetBody(v *CreateK8sConfigMapResponseBody) *CreateK8sConfigMapResponse {
	s.Body = v
	return s
}

func (s *CreateK8sConfigMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
