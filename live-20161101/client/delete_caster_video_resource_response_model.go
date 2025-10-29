// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterVideoResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterVideoResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterVideoResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterVideoResourceResponseBody) *DeleteCasterVideoResourceResponse
	GetBody() *DeleteCasterVideoResourceResponseBody
}

type DeleteCasterVideoResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterVideoResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterVideoResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterVideoResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterVideoResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterVideoResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterVideoResourceResponse) GetBody() *DeleteCasterVideoResourceResponseBody {
	return s.Body
}

func (s *DeleteCasterVideoResourceResponse) SetHeaders(v map[string]*string) *DeleteCasterVideoResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterVideoResourceResponse) SetStatusCode(v int32) *DeleteCasterVideoResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterVideoResourceResponse) SetBody(v *DeleteCasterVideoResourceResponseBody) *DeleteCasterVideoResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterVideoResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
