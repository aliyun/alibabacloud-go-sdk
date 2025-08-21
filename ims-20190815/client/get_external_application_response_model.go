// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExternalApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExternalApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetExternalApplicationResponseBody) *GetExternalApplicationResponse
	GetBody() *GetExternalApplicationResponseBody
}

type GetExternalApplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExternalApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExternalApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExternalApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExternalApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExternalApplicationResponse) GetBody() *GetExternalApplicationResponseBody {
	return s.Body
}

func (s *GetExternalApplicationResponse) SetHeaders(v map[string]*string) *GetExternalApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetExternalApplicationResponse) SetStatusCode(v int32) *GetExternalApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExternalApplicationResponse) SetBody(v *GetExternalApplicationResponseBody) *GetExternalApplicationResponse {
	s.Body = v
	return s
}

func (s *GetExternalApplicationResponse) Validate() error {
	return dara.Validate(s)
}
