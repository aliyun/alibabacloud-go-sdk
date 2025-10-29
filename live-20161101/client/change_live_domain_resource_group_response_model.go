// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeLiveDomainResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeLiveDomainResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeLiveDomainResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeLiveDomainResourceGroupResponseBody) *ChangeLiveDomainResourceGroupResponse
	GetBody() *ChangeLiveDomainResourceGroupResponseBody
}

type ChangeLiveDomainResourceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeLiveDomainResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeLiveDomainResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeLiveDomainResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeLiveDomainResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeLiveDomainResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeLiveDomainResourceGroupResponse) GetBody() *ChangeLiveDomainResourceGroupResponseBody {
	return s.Body
}

func (s *ChangeLiveDomainResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeLiveDomainResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeLiveDomainResourceGroupResponse) SetStatusCode(v int32) *ChangeLiveDomainResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeLiveDomainResourceGroupResponse) SetBody(v *ChangeLiveDomainResourceGroupResponseBody) *ChangeLiveDomainResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeLiveDomainResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
