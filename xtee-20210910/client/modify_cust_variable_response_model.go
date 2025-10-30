// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustVariableResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustVariableResponseBody) *ModifyCustVariableResponse
	GetBody() *ModifyCustVariableResponseBody
}

type ModifyCustVariableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustVariableResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustVariableResponse) GetBody() *ModifyCustVariableResponseBody {
	return s.Body
}

func (s *ModifyCustVariableResponse) SetHeaders(v map[string]*string) *ModifyCustVariableResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustVariableResponse) SetStatusCode(v int32) *ModifyCustVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustVariableResponse) SetBody(v *ModifyCustVariableResponseBody) *ModifyCustVariableResponse {
	s.Body = v
	return s
}

func (s *ModifyCustVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
