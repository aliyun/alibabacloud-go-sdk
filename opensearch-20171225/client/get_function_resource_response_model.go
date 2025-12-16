// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionResourceResponseBody) *GetFunctionResourceResponse
	GetBody() *GetFunctionResourceResponseBody
}

type GetFunctionResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionResourceResponse) GetBody() *GetFunctionResourceResponseBody {
	return s.Body
}

func (s *GetFunctionResourceResponse) SetHeaders(v map[string]*string) *GetFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResourceResponse) SetStatusCode(v int32) *GetFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResourceResponse) SetBody(v *GetFunctionResourceResponseBody) *GetFunctionResourceResponse {
	s.Body = v
	return s
}

func (s *GetFunctionResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
