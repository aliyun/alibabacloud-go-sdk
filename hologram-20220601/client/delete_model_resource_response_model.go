// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelResourceResponseBody) *DeleteModelResourceResponse
	GetBody() *DeleteModelResourceResponseBody
}

type DeleteModelResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelResourceResponse) GetBody() *DeleteModelResourceResponseBody {
	return s.Body
}

func (s *DeleteModelResourceResponse) SetHeaders(v map[string]*string) *DeleteModelResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResourceResponse) SetStatusCode(v int32) *DeleteModelResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelResourceResponse) SetBody(v *DeleteModelResourceResponseBody) *DeleteModelResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteModelResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
