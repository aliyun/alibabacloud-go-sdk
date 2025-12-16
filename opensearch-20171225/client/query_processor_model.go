// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProcessor interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *QueryProcessor
	GetActive() *bool
	SetCategory(v string) *QueryProcessor
	GetCategory() *string
	SetDomain(v string) *QueryProcessor
	GetDomain() *string
	SetIndexes(v []*string) *QueryProcessor
	GetIndexes() []*string
	SetName(v string) *QueryProcessor
	GetName() *string
	SetProcessors(v []map[string]interface{}) *QueryProcessor
	GetProcessors() []map[string]interface{}
}

type QueryProcessor struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Category   *string                  `json:"category,omitempty" xml:"category,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s QueryProcessor) String() string {
	return dara.Prettify(s)
}

func (s QueryProcessor) GoString() string {
	return s.String()
}

func (s *QueryProcessor) GetActive() *bool {
	return s.Active
}

func (s *QueryProcessor) GetCategory() *string {
	return s.Category
}

func (s *QueryProcessor) GetDomain() *string {
	return s.Domain
}

func (s *QueryProcessor) GetIndexes() []*string {
	return s.Indexes
}

func (s *QueryProcessor) GetName() *string {
	return s.Name
}

func (s *QueryProcessor) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *QueryProcessor) SetActive(v bool) *QueryProcessor {
	s.Active = &v
	return s
}

func (s *QueryProcessor) SetCategory(v string) *QueryProcessor {
	s.Category = &v
	return s
}

func (s *QueryProcessor) SetDomain(v string) *QueryProcessor {
	s.Domain = &v
	return s
}

func (s *QueryProcessor) SetIndexes(v []*string) *QueryProcessor {
	s.Indexes = v
	return s
}

func (s *QueryProcessor) SetName(v string) *QueryProcessor {
	s.Name = &v
	return s
}

func (s *QueryProcessor) SetProcessors(v []map[string]interface{}) *QueryProcessor {
	s.Processors = v
	return s
}

func (s *QueryProcessor) Validate() error {
	return dara.Validate(s)
}
