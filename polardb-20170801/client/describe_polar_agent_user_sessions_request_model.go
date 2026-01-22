// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentUserSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *DescribePolarAgentUserSessionsRequest
	GetSource() *string
}

type DescribePolarAgentUserSessionsRequest struct {
	// example:
	//
	// polardb-console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePolarAgentUserSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentUserSessionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentUserSessionsRequest) GetSource() *string {
	return s.Source
}

func (s *DescribePolarAgentUserSessionsRequest) SetSource(v string) *DescribePolarAgentUserSessionsRequest {
	s.Source = &v
	return s
}

func (s *DescribePolarAgentUserSessionsRequest) Validate() error {
	return dara.Validate(s)
}
