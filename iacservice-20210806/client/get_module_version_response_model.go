// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetModuleVersionResponseBody) *GetModuleVersionResponse
	GetBody() *GetModuleVersionResponseBody
}

type GetModuleVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *GetModuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModuleVersionResponse) GetBody() *GetModuleVersionResponseBody {
	return s.Body
}

func (s *GetModuleVersionResponse) SetHeaders(v map[string]*string) *GetModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *GetModuleVersionResponse) SetStatusCode(v int32) *GetModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleVersionResponse) SetBody(v *GetModuleVersionResponseBody) *GetModuleVersionResponse {
	s.Body = v
	return s
}

func (s *GetModuleVersionResponse) Validate() error {
	return dara.Validate(s)
}
