// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountForAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountForAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountForAppResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountForAppResponseBody) *GetAccountForAppResponse
	GetBody() *GetAccountForAppResponseBody
}

type GetAccountForAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountForAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountForAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppResponse) GoString() string {
	return s.String()
}

func (s *GetAccountForAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountForAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountForAppResponse) GetBody() *GetAccountForAppResponseBody {
	return s.Body
}

func (s *GetAccountForAppResponse) SetHeaders(v map[string]*string) *GetAccountForAppResponse {
	s.Headers = v
	return s
}

func (s *GetAccountForAppResponse) SetStatusCode(v int32) *GetAccountForAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountForAppResponse) SetBody(v *GetAccountForAppResponseBody) *GetAccountForAppResponse {
	s.Body = v
	return s
}

func (s *GetAccountForAppResponse) Validate() error {
	return dara.Validate(s)
}
