// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustVariableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustVariableResponseBody) *DeleteCustVariableResponse
	GetBody() *DeleteCustVariableResponseBody
}

type DeleteCustVariableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustVariableResponse) GetBody() *DeleteCustVariableResponseBody {
	return s.Body
}

func (s *DeleteCustVariableResponse) SetHeaders(v map[string]*string) *DeleteCustVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustVariableResponse) SetStatusCode(v int32) *DeleteCustVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustVariableResponse) SetBody(v *DeleteCustVariableResponseBody) *DeleteCustVariableResponse {
	s.Body = v
	return s
}

func (s *DeleteCustVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
