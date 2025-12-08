// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeepfakeFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeepfakeFaceResponse
	GetStatusCode() *int32
	SetBody(v *DeepfakeFaceResponseBody) *DeepfakeFaceResponse
	GetBody() *DeepfakeFaceResponseBody
}

type DeepfakeFaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeepfakeFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeepfakeFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceResponse) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeepfakeFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeepfakeFaceResponse) GetBody() *DeepfakeFaceResponseBody {
	return s.Body
}

func (s *DeepfakeFaceResponse) SetHeaders(v map[string]*string) *DeepfakeFaceResponse {
	s.Headers = v
	return s
}

func (s *DeepfakeFaceResponse) SetStatusCode(v int32) *DeepfakeFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeepfakeFaceResponse) SetBody(v *DeepfakeFaceResponseBody) *DeepfakeFaceResponse {
	s.Body = v
	return s
}

func (s *DeepfakeFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
