// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteK8sApplicationResponseBody) *DeleteK8sApplicationResponse
	GetBody() *DeleteK8sApplicationResponseBody
}

type DeleteK8sApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteK8sApplicationResponse) GetBody() *DeleteK8sApplicationResponseBody {
	return s.Body
}

func (s *DeleteK8sApplicationResponse) SetHeaders(v map[string]*string) *DeleteK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteK8sApplicationResponse) SetStatusCode(v int32) *DeleteK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteK8sApplicationResponse) SetBody(v *DeleteK8sApplicationResponseBody) *DeleteK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *DeleteK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
