// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQueryProcessorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQueryProcessorResponseBody
	GetRequestId() *string
	SetResult(v *ModifyQueryProcessorResponseBodyResult) *ModifyQueryProcessorResponseBody
	GetResult() *ModifyQueryProcessorResponseBodyResult
}

type ModifyQueryProcessorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	//
	// example:
	//
	// {}
	Result *ModifyQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyQueryProcessorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQueryProcessorResponseBody) GetResult() *ModifyQueryProcessorResponseBodyResult {
	return s.Result
}

func (s *ModifyQueryProcessorResponseBody) SetRequestId(v string) *ModifyQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQueryProcessorResponseBody) SetResult(v *ModifyQueryProcessorResponseBodyResult) *ModifyQueryProcessorResponseBody {
	s.Result = v
	return s
}

func (s *ModifyQueryProcessorResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyQueryProcessorResponseBodyResult struct {
	// Indicates whether the query analysis rule is a default rule.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the rule was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry to which the query analysis rule is applied. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule is applied.
	//
	// example:
	//
	// ["default"]
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// synonym
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The analysis rule.
	//
	// example:
	//
	// []
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the rule was updated.
	//
	// example:
	//
	// 1
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyQueryProcessorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ModifyQueryProcessorResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ModifyQueryProcessorResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ModifyQueryProcessorResponseBodyResult) GetIndexes() []*string {
	return s.Indexes
}

func (s *ModifyQueryProcessorResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ModifyQueryProcessorResponseBodyResult) GetProcessors() []map[string]interface{} {
	return s.Processors
}

func (s *ModifyQueryProcessorResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ModifyQueryProcessorResponseBodyResult) SetActive(v bool) *ModifyQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetCreated(v int32) *ModifyQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetDomain(v string) *ModifyQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetIndexes(v []*string) *ModifyQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetName(v string) *ModifyQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *ModifyQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetUpdated(v int32) *ModifyQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
