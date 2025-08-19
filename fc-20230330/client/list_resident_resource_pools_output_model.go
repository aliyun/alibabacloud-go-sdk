// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResidentResourcePoolsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListResidentResourcePoolsOutput
	GetNextToken() *string
	SetResidentResourcePools(v []*ResidentResourcePool) *ListResidentResourcePoolsOutput
	GetResidentResourcePools() []*ResidentResourcePool
}

type ListResidentResourcePoolsOutput struct {
	NextToken             *string                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResidentResourcePools []*ResidentResourcePool `json:"residentResourcePools" xml:"residentResourcePools" type:"Repeated"`
}

func (s ListResidentResourcePoolsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListResidentResourcePoolsOutput) GoString() string {
	return s.String()
}

func (s *ListResidentResourcePoolsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResidentResourcePoolsOutput) GetResidentResourcePools() []*ResidentResourcePool {
	return s.ResidentResourcePools
}

func (s *ListResidentResourcePoolsOutput) SetNextToken(v string) *ListResidentResourcePoolsOutput {
	s.NextToken = &v
	return s
}

func (s *ListResidentResourcePoolsOutput) SetResidentResourcePools(v []*ResidentResourcePool) *ListResidentResourcePoolsOutput {
	s.ResidentResourcePools = v
	return s
}

func (s *ListResidentResourcePoolsOutput) Validate() error {
	return dara.Validate(s)
}
