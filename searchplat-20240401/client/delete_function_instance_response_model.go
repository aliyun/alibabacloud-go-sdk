// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFunctionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFunctionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFunctionInstanceResponseBody) *DeleteFunctionInstanceResponse
	GetBody() *DeleteFunctionInstanceResponseBody
}

type DeleteFunctionInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFunctionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFunctionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFunctionInstanceResponse) GetBody() *DeleteFunctionInstanceResponseBody {
	return s.Body
}

func (s *DeleteFunctionInstanceResponse) SetHeaders(v map[string]*string) *DeleteFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionInstanceResponse) SetStatusCode(v int32) *DeleteFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionInstanceResponse) SetBody(v *DeleteFunctionInstanceResponseBody) *DeleteFunctionInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteFunctionInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
