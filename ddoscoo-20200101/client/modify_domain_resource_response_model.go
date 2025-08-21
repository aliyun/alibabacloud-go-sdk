// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDomainResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDomainResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDomainResourceResponseBody) *ModifyDomainResourceResponse
	GetBody() *ModifyDomainResourceResponseBody
}

type ModifyDomainResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDomainResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDomainResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDomainResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDomainResourceResponse) GetBody() *ModifyDomainResourceResponseBody {
	return s.Body
}

func (s *ModifyDomainResourceResponse) SetHeaders(v map[string]*string) *ModifyDomainResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainResourceResponse) SetStatusCode(v int32) *ModifyDomainResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainResourceResponse) SetBody(v *ModifyDomainResourceResponseBody) *ModifyDomainResourceResponse {
	s.Body = v
	return s
}

func (s *ModifyDomainResourceResponse) Validate() error {
	return dara.Validate(s)
}
