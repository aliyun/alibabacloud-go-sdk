// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListVersionsOutput
	GetDirection() *string
	SetNextToken(v string) *ListVersionsOutput
	GetNextToken() *string
	SetVersions(v []*Version) *ListVersionsOutput
	GetVersions() []*Version
}

type ListVersionsOutput struct {
	// example:
	//
	// FORWARD
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// example:
	//
	// 3
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Versions  []*Version `json:"versions" xml:"versions" type:"Repeated"`
}

func (s ListVersionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsOutput) GoString() string {
	return s.String()
}

func (s *ListVersionsOutput) GetDirection() *string {
	return s.Direction
}

func (s *ListVersionsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVersionsOutput) GetVersions() []*Version {
	return s.Versions
}

func (s *ListVersionsOutput) SetDirection(v string) *ListVersionsOutput {
	s.Direction = &v
	return s
}

func (s *ListVersionsOutput) SetNextToken(v string) *ListVersionsOutput {
	s.NextToken = &v
	return s
}

func (s *ListVersionsOutput) SetVersions(v []*Version) *ListVersionsOutput {
	s.Versions = v
	return s
}

func (s *ListVersionsOutput) Validate() error {
	return dara.Validate(s)
}
