// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ScaleK8sApplicationResponseBody) *ScaleK8sApplicationResponse
	GetBody() *ScaleK8sApplicationResponseBody
}

type ScaleK8sApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *ScaleK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleK8sApplicationResponse) GetBody() *ScaleK8sApplicationResponseBody {
	return s.Body
}

func (s *ScaleK8sApplicationResponse) SetHeaders(v map[string]*string) *ScaleK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *ScaleK8sApplicationResponse) SetStatusCode(v int32) *ScaleK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleK8sApplicationResponse) SetBody(v *ScaleK8sApplicationResponseBody) *ScaleK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *ScaleK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
