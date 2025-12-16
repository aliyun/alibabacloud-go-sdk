// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionDefaultInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionDefaultInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionDefaultInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionDefaultInstanceResponseBody) *GetFunctionDefaultInstanceResponse
	GetBody() *GetFunctionDefaultInstanceResponseBody
}

type GetFunctionDefaultInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionDefaultInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionDefaultInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionDefaultInstanceResponse) GetBody() *GetFunctionDefaultInstanceResponseBody {
	return s.Body
}

func (s *GetFunctionDefaultInstanceResponse) SetHeaders(v map[string]*string) *GetFunctionDefaultInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionDefaultInstanceResponse) SetStatusCode(v int32) *GetFunctionDefaultInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponse) SetBody(v *GetFunctionDefaultInstanceResponseBody) *GetFunctionDefaultInstanceResponse {
	s.Body = v
	return s
}

func (s *GetFunctionDefaultInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
