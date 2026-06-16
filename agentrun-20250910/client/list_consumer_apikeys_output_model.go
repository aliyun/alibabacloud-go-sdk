// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerAPIKeysOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ConsumerAPIKey) *ListConsumerAPIKeysOutput
	GetItems() []*ConsumerAPIKey
	SetPageNumber(v int) *ListConsumerAPIKeysOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListConsumerAPIKeysOutput
	GetPageSize() *int
	SetTotal(v int) *ListConsumerAPIKeysOutput
	GetTotal() *int
}

type ListConsumerAPIKeysOutput struct {
	// A list of consumer API keys.
	Items []*ConsumerAPIKey `json:"items" xml:"items" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of matching consumer API keys.
	//
	// example:
	//
	// 10
	Total *int `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListConsumerAPIKeysOutput) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerAPIKeysOutput) GoString() string {
	return s.String()
}

func (s *ListConsumerAPIKeysOutput) GetItems() []*ConsumerAPIKey {
	return s.Items
}

func (s *ListConsumerAPIKeysOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListConsumerAPIKeysOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListConsumerAPIKeysOutput) GetTotal() *int {
	return s.Total
}

func (s *ListConsumerAPIKeysOutput) SetItems(v []*ConsumerAPIKey) *ListConsumerAPIKeysOutput {
	s.Items = v
	return s
}

func (s *ListConsumerAPIKeysOutput) SetPageNumber(v int) *ListConsumerAPIKeysOutput {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerAPIKeysOutput) SetPageSize(v int) *ListConsumerAPIKeysOutput {
	s.PageSize = &v
	return s
}

func (s *ListConsumerAPIKeysOutput) SetTotal(v int) *ListConsumerAPIKeysOutput {
	s.Total = &v
	return s
}

func (s *ListConsumerAPIKeysOutput) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
