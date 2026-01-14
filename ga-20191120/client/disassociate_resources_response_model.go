// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateResourcesResponseBody) *DisassociateResourcesResponse
	GetBody() *DisassociateResourcesResponseBody
}

type DisassociateResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourcesResponse) GoString() string {
	return s.String()
}

func (s *DisassociateResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateResourcesResponse) GetBody() *DisassociateResourcesResponseBody {
	return s.Body
}

func (s *DisassociateResourcesResponse) SetHeaders(v map[string]*string) *DisassociateResourcesResponse {
	s.Headers = v
	return s
}

func (s *DisassociateResourcesResponse) SetStatusCode(v int32) *DisassociateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateResourcesResponse) SetBody(v *DisassociateResourcesResponseBody) *DisassociateResourcesResponse {
	s.Body = v
	return s
}

func (s *DisassociateResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
