// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionVersionResponseBody) *GetFunctionVersionResponse
	GetBody() *GetFunctionVersionResponseBody
}

type GetFunctionVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionVersionResponse) GetBody() *GetFunctionVersionResponseBody {
	return s.Body
}

func (s *GetFunctionVersionResponse) SetHeaders(v map[string]*string) *GetFunctionVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionVersionResponse) SetStatusCode(v int32) *GetFunctionVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionVersionResponse) SetBody(v *GetFunctionVersionResponseBody) *GetFunctionVersionResponse {
	s.Body = v
	return s
}

func (s *GetFunctionVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
