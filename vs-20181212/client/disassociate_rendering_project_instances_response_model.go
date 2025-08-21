// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateRenderingProjectInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateRenderingProjectInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateRenderingProjectInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateRenderingProjectInstancesResponseBody) *DisassociateRenderingProjectInstancesResponse
	GetBody() *DisassociateRenderingProjectInstancesResponseBody
}

type DisassociateRenderingProjectInstancesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateRenderingProjectInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateRenderingProjectInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateRenderingProjectInstancesResponse) GoString() string {
	return s.String()
}

func (s *DisassociateRenderingProjectInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateRenderingProjectInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateRenderingProjectInstancesResponse) GetBody() *DisassociateRenderingProjectInstancesResponseBody {
	return s.Body
}

func (s *DisassociateRenderingProjectInstancesResponse) SetHeaders(v map[string]*string) *DisassociateRenderingProjectInstancesResponse {
	s.Headers = v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponse) SetStatusCode(v int32) *DisassociateRenderingProjectInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponse) SetBody(v *DisassociateRenderingProjectInstancesResponseBody) *DisassociateRenderingProjectInstancesResponse {
	s.Body = v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponse) Validate() error {
	return dara.Validate(s)
}
