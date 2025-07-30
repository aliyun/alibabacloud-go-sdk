// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationResponseBody) *GetApplicationResponse
	GetBody() *GetApplicationResponseBody
}

type GetApplicationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationResponse) GetBody() *GetApplicationResponseBody {
	return s.Body
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetStatusCode(v int32) *GetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

func (s *GetApplicationResponse) Validate() error {
	return dara.Validate(s)
}
