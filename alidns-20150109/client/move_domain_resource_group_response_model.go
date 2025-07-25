// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveDomainResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveDomainResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveDomainResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *MoveDomainResourceGroupResponseBody) *MoveDomainResourceGroupResponse
	GetBody() *MoveDomainResourceGroupResponseBody
}

type MoveDomainResourceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveDomainResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveDomainResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveDomainResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveDomainResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveDomainResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveDomainResourceGroupResponse) GetBody() *MoveDomainResourceGroupResponseBody {
	return s.Body
}

func (s *MoveDomainResourceGroupResponse) SetHeaders(v map[string]*string) *MoveDomainResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveDomainResourceGroupResponse) SetStatusCode(v int32) *MoveDomainResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveDomainResourceGroupResponse) SetBody(v *MoveDomainResourceGroupResponseBody) *MoveDomainResourceGroupResponse {
	s.Body = v
	return s
}

func (s *MoveDomainResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
