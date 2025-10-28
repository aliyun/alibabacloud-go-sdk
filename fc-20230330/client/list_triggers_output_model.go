// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTriggersOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTriggersOutput
	GetNextToken() *string
	SetTriggers(v []*Trigger) *ListTriggersOutput
	GetTriggers() []*Trigger
}

type ListTriggersOutput struct {
	// example:
	//
	// next_token
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Triggers  []*Trigger `json:"triggers" xml:"triggers" type:"Repeated"`
}

func (s ListTriggersOutput) String() string {
	return dara.Prettify(s)
}

func (s ListTriggersOutput) GoString() string {
	return s.String()
}

func (s *ListTriggersOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTriggersOutput) GetTriggers() []*Trigger {
	return s.Triggers
}

func (s *ListTriggersOutput) SetNextToken(v string) *ListTriggersOutput {
	s.NextToken = &v
	return s
}

func (s *ListTriggersOutput) SetTriggers(v []*Trigger) *ListTriggersOutput {
	s.Triggers = v
	return s
}

func (s *ListTriggersOutput) Validate() error {
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
