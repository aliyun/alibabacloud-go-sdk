// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionCurrentVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionCurrentVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionCurrentVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionCurrentVersionResponseBody) *GetFunctionCurrentVersionResponse
	GetBody() *GetFunctionCurrentVersionResponseBody
}

type GetFunctionCurrentVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionCurrentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionCurrentVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionCurrentVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionCurrentVersionResponse) GetBody() *GetFunctionCurrentVersionResponseBody {
	return s.Body
}

func (s *GetFunctionCurrentVersionResponse) SetHeaders(v map[string]*string) *GetFunctionCurrentVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionCurrentVersionResponse) SetStatusCode(v int32) *GetFunctionCurrentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionCurrentVersionResponse) SetBody(v *GetFunctionCurrentVersionResponseBody) *GetFunctionCurrentVersionResponse {
	s.Body = v
	return s
}

func (s *GetFunctionCurrentVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
