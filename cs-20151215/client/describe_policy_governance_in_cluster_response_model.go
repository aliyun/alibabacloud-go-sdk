// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyGovernanceInClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyGovernanceInClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyGovernanceInClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyGovernanceInClusterResponseBody) *DescribePolicyGovernanceInClusterResponse
	GetBody() *DescribePolicyGovernanceInClusterResponseBody
}

type DescribePolicyGovernanceInClusterResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyGovernanceInClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyGovernanceInClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyGovernanceInClusterResponse) GetBody() *DescribePolicyGovernanceInClusterResponseBody {
	return s.Body
}

func (s *DescribePolicyGovernanceInClusterResponse) SetHeaders(v map[string]*string) *DescribePolicyGovernanceInClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponse) SetStatusCode(v int32) *DescribePolicyGovernanceInClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponse) SetBody(v *DescribePolicyGovernanceInClusterResponseBody) *DescribePolicyGovernanceInClusterResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
