// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFunctionResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFunctionResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFunctionResourceResponseBody) *DeleteFunctionResourceResponse
	GetBody() *DeleteFunctionResourceResponseBody
}

type DeleteFunctionResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFunctionResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFunctionResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFunctionResourceResponse) GetBody() *DeleteFunctionResourceResponseBody {
	return s.Body
}

func (s *DeleteFunctionResourceResponse) SetHeaders(v map[string]*string) *DeleteFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResourceResponse) SetStatusCode(v int32) *DeleteFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionResourceResponse) SetBody(v *DeleteFunctionResourceResponseBody) *DeleteFunctionResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteFunctionResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
