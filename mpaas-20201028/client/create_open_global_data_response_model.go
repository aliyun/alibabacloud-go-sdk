// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenGlobalDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpenGlobalDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpenGlobalDataResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpenGlobalDataResponseBody) *CreateOpenGlobalDataResponse
	GetBody() *CreateOpenGlobalDataResponseBody
}

type CreateOpenGlobalDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpenGlobalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpenGlobalDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenGlobalDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOpenGlobalDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpenGlobalDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpenGlobalDataResponse) GetBody() *CreateOpenGlobalDataResponseBody {
	return s.Body
}

func (s *CreateOpenGlobalDataResponse) SetHeaders(v map[string]*string) *CreateOpenGlobalDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOpenGlobalDataResponse) SetStatusCode(v int32) *CreateOpenGlobalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpenGlobalDataResponse) SetBody(v *CreateOpenGlobalDataResponseBody) *CreateOpenGlobalDataResponse {
	s.Body = v
	return s
}

func (s *CreateOpenGlobalDataResponse) Validate() error {
	return dara.Validate(s)
}
