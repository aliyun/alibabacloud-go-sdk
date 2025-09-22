// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgAndFactoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrgAndFactoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrgAndFactoryResponse
	GetStatusCode() *int32
	SetBody(v *GetOrgAndFactoryResponseBody) *GetOrgAndFactoryResponse
	GetBody() *GetOrgAndFactoryResponseBody
}

type GetOrgAndFactoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgAndFactoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgAndFactoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrgAndFactoryResponse) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrgAndFactoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrgAndFactoryResponse) GetBody() *GetOrgAndFactoryResponseBody {
	return s.Body
}

func (s *GetOrgAndFactoryResponse) SetHeaders(v map[string]*string) *GetOrgAndFactoryResponse {
	s.Headers = v
	return s
}

func (s *GetOrgAndFactoryResponse) SetStatusCode(v int32) *GetOrgAndFactoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgAndFactoryResponse) SetBody(v *GetOrgAndFactoryResponseBody) *GetOrgAndFactoryResponse {
	s.Body = v
	return s
}

func (s *GetOrgAndFactoryResponse) Validate() error {
	return dara.Validate(s)
}
