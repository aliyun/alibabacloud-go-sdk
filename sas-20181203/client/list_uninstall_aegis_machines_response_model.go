// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUninstallAegisMachinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUninstallAegisMachinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUninstallAegisMachinesResponse
	GetStatusCode() *int32
	SetBody(v *ListUninstallAegisMachinesResponseBody) *ListUninstallAegisMachinesResponse
	GetBody() *ListUninstallAegisMachinesResponseBody
}

type ListUninstallAegisMachinesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUninstallAegisMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUninstallAegisMachinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallAegisMachinesResponse) GoString() string {
	return s.String()
}

func (s *ListUninstallAegisMachinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUninstallAegisMachinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUninstallAegisMachinesResponse) GetBody() *ListUninstallAegisMachinesResponseBody {
	return s.Body
}

func (s *ListUninstallAegisMachinesResponse) SetHeaders(v map[string]*string) *ListUninstallAegisMachinesResponse {
	s.Headers = v
	return s
}

func (s *ListUninstallAegisMachinesResponse) SetStatusCode(v int32) *ListUninstallAegisMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUninstallAegisMachinesResponse) SetBody(v *ListUninstallAegisMachinesResponseBody) *ListUninstallAegisMachinesResponse {
	s.Body = v
	return s
}

func (s *ListUninstallAegisMachinesResponse) Validate() error {
	return dara.Validate(s)
}
