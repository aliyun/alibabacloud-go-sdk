// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCeaseFunctionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CeaseFunctionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CeaseFunctionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CeaseFunctionInstanceResponseBody) *CeaseFunctionInstanceResponse
	GetBody() *CeaseFunctionInstanceResponseBody
}

type CeaseFunctionInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CeaseFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CeaseFunctionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CeaseFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *CeaseFunctionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CeaseFunctionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CeaseFunctionInstanceResponse) GetBody() *CeaseFunctionInstanceResponseBody {
	return s.Body
}

func (s *CeaseFunctionInstanceResponse) SetHeaders(v map[string]*string) *CeaseFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *CeaseFunctionInstanceResponse) SetStatusCode(v int32) *CeaseFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CeaseFunctionInstanceResponse) SetBody(v *CeaseFunctionInstanceResponseBody) *CeaseFunctionInstanceResponse {
	s.Body = v
	return s
}

func (s *CeaseFunctionInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
