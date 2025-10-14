// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafUsageOfRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchConfigUsage(v int64) *ListWafUsageOfRulesResponseBody
	GetBatchConfigUsage() *int64
	SetInstanceUsage(v int64) *ListWafUsageOfRulesResponseBody
	GetInstanceUsage() *int64
	SetRequestId(v string) *ListWafUsageOfRulesResponseBody
	GetRequestId() *string
	SetSites(v []*ListWafUsageOfRulesResponseBodySites) *ListWafUsageOfRulesResponseBody
	GetSites() []*ListWafUsageOfRulesResponseBodySites
}

type ListWafUsageOfRulesResponseBody struct {
	BatchConfigUsage *int64 `json:"BatchConfigUsage,omitempty" xml:"BatchConfigUsage,omitempty"`
	InstanceUsage    *int64 `json:"InstanceUsage,omitempty" xml:"InstanceUsage,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of site usage.
	Sites []*ListWafUsageOfRulesResponseBodySites `json:"Sites,omitempty" xml:"Sites,omitempty" type:"Repeated"`
}

func (s ListWafUsageOfRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWafUsageOfRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesResponseBody) GetBatchConfigUsage() *int64 {
	return s.BatchConfigUsage
}

func (s *ListWafUsageOfRulesResponseBody) GetInstanceUsage() *int64 {
	return s.InstanceUsage
}

func (s *ListWafUsageOfRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWafUsageOfRulesResponseBody) GetSites() []*ListWafUsageOfRulesResponseBodySites {
	return s.Sites
}

func (s *ListWafUsageOfRulesResponseBody) SetBatchConfigUsage(v int64) *ListWafUsageOfRulesResponseBody {
	s.BatchConfigUsage = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBody) SetInstanceUsage(v int64) *ListWafUsageOfRulesResponseBody {
	s.InstanceUsage = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBody) SetRequestId(v string) *ListWafUsageOfRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBody) SetSites(v []*ListWafUsageOfRulesResponseBodySites) *ListWafUsageOfRulesResponseBody {
	s.Sites = v
	return s
}

func (s *ListWafUsageOfRulesResponseBody) Validate() error {
	if s.Sites != nil {
		for _, item := range s.Sites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWafUsageOfRulesResponseBodySites struct {
	// Site ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Site name.
	//
	// example:
	//
	// example.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Usage of WAF rules/WAF rule sets.
	//
	// example:
	//
	// 1
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListWafUsageOfRulesResponseBodySites) String() string {
	return dara.Prettify(s)
}

func (s ListWafUsageOfRulesResponseBodySites) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesResponseBodySites) GetId() *int64 {
	return s.Id
}

func (s *ListWafUsageOfRulesResponseBodySites) GetName() *string {
	return s.Name
}

func (s *ListWafUsageOfRulesResponseBodySites) GetUsage() *int64 {
	return s.Usage
}

func (s *ListWafUsageOfRulesResponseBodySites) SetId(v int64) *ListWafUsageOfRulesResponseBodySites {
	s.Id = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBodySites) SetName(v string) *ListWafUsageOfRulesResponseBodySites {
	s.Name = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBodySites) SetUsage(v int64) *ListWafUsageOfRulesResponseBodySites {
	s.Usage = &v
	return s
}

func (s *ListWafUsageOfRulesResponseBodySites) Validate() error {
	return dara.Validate(s)
}
