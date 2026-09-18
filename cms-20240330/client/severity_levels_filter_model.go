// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeverityLevelsFilter interface {
	dara.Model
	String() string
	GoString() string
	SetContains(v []*string) *SeverityLevelsFilter
	GetContains() []*string
}

type SeverityLevelsFilter struct {
	// The hit condition: the set of severity levels covered by the rule contains at least one level in the array (OR semantics).
	//
	// example:
	//
	// ["CRITICAL","ERROR"]
	Contains []*string `json:"contains,omitempty" xml:"contains,omitempty" type:"Repeated"`
}

func (s SeverityLevelsFilter) String() string {
	return dara.Prettify(s)
}

func (s SeverityLevelsFilter) GoString() string {
	return s.String()
}

func (s *SeverityLevelsFilter) GetContains() []*string {
	return s.Contains
}

func (s *SeverityLevelsFilter) SetContains(v []*string) *SeverityLevelsFilter {
	s.Contains = v
	return s
}

func (s *SeverityLevelsFilter) Validate() error {
	return dara.Validate(s)
}
