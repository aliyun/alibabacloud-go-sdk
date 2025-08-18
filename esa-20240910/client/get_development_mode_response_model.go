// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDevelopmentModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDevelopmentModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDevelopmentModeResponse
	GetStatusCode() *int32
	SetBody(v *GetDevelopmentModeResponseBody) *GetDevelopmentModeResponse
	GetBody() *GetDevelopmentModeResponseBody
}

type GetDevelopmentModeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDevelopmentModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDevelopmentModeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDevelopmentModeResponse) GoString() string {
	return s.String()
}

func (s *GetDevelopmentModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDevelopmentModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDevelopmentModeResponse) GetBody() *GetDevelopmentModeResponseBody {
	return s.Body
}

func (s *GetDevelopmentModeResponse) SetHeaders(v map[string]*string) *GetDevelopmentModeResponse {
	s.Headers = v
	return s
}

func (s *GetDevelopmentModeResponse) SetStatusCode(v int32) *GetDevelopmentModeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevelopmentModeResponse) SetBody(v *GetDevelopmentModeResponseBody) *GetDevelopmentModeResponse {
	s.Body = v
	return s
}

func (s *GetDevelopmentModeResponse) Validate() error {
	return dara.Validate(s)
}
