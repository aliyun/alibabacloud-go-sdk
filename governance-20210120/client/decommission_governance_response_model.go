// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecommissionGovernanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DecommissionGovernanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DecommissionGovernanceResponse
	GetStatusCode() *int32
	SetBody(v *DecommissionGovernanceResponseBody) *DecommissionGovernanceResponse
	GetBody() *DecommissionGovernanceResponseBody
}

type DecommissionGovernanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecommissionGovernanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecommissionGovernanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DecommissionGovernanceResponse) GoString() string {
	return s.String()
}

func (s *DecommissionGovernanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DecommissionGovernanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DecommissionGovernanceResponse) GetBody() *DecommissionGovernanceResponseBody {
	return s.Body
}

func (s *DecommissionGovernanceResponse) SetHeaders(v map[string]*string) *DecommissionGovernanceResponse {
	s.Headers = v
	return s
}

func (s *DecommissionGovernanceResponse) SetStatusCode(v int32) *DecommissionGovernanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DecommissionGovernanceResponse) SetBody(v *DecommissionGovernanceResponseBody) *DecommissionGovernanceResponse {
	s.Body = v
	return s
}

func (s *DecommissionGovernanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
