// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExternalAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeExternalAgentResponseBody
	GetConfig() *string
}

type DescribeExternalAgentResponseBody struct {
	// The agent configurations in the YAML format.
	//
	// example:
	//
	// apiVersion: v1****
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s DescribeExternalAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExternalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeExternalAgentResponseBody) SetConfig(v string) *DescribeExternalAgentResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeExternalAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
