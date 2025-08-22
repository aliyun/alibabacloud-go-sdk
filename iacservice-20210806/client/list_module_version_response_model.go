// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *ListModuleVersionResponseBody) *ListModuleVersionResponse
	GetBody() *ListModuleVersionResponseBody
}

type ListModuleVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *ListModuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModuleVersionResponse) GetBody() *ListModuleVersionResponseBody {
	return s.Body
}

func (s *ListModuleVersionResponse) SetHeaders(v map[string]*string) *ListModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *ListModuleVersionResponse) SetStatusCode(v int32) *ListModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModuleVersionResponse) SetBody(v *ListModuleVersionResponseBody) *ListModuleVersionResponse {
	s.Body = v
	return s
}

func (s *ListModuleVersionResponse) Validate() error {
	return dara.Validate(s)
}
