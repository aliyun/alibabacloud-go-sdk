// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sConfigMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteK8sConfigMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteK8sConfigMapResponse
	GetStatusCode() *int32
	SetBody(v *DeleteK8sConfigMapResponseBody) *DeleteK8sConfigMapResponse
	GetBody() *DeleteK8sConfigMapResponseBody
}

type DeleteK8sConfigMapResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteK8sConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteK8sConfigMapResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sConfigMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteK8sConfigMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteK8sConfigMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteK8sConfigMapResponse) GetBody() *DeleteK8sConfigMapResponseBody {
	return s.Body
}

func (s *DeleteK8sConfigMapResponse) SetHeaders(v map[string]*string) *DeleteK8sConfigMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteK8sConfigMapResponse) SetStatusCode(v int32) *DeleteK8sConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteK8sConfigMapResponse) SetBody(v *DeleteK8sConfigMapResponseBody) *DeleteK8sConfigMapResponse {
	s.Body = v
	return s
}

func (s *DeleteK8sConfigMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
