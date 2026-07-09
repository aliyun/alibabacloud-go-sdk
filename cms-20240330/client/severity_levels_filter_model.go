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
	// Matches a log entry if its severity level appears in this array of strings.
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
