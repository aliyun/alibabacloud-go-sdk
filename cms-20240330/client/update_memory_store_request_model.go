// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *UpdateMemoryStoreRequest
	GetCustomExtractionStrategies() []*CustomExtractionStrategy
	SetDescription(v string) *UpdateMemoryStoreRequest
	GetDescription() *string
	SetExtractionStrategies(v []*string) *UpdateMemoryStoreRequest
	GetExtractionStrategies() []*string
	SetShortTermTtl(v int32) *UpdateMemoryStoreRequest
	GetShortTermTtl() *int32
}

type UpdateMemoryStoreRequest struct {
	CustomExtractionStrategies []*CustomExtractionStrategy `json:"customExtractionStrategies,omitempty" xml:"customExtractionStrategies,omitempty" type:"Repeated"`
	// example:
	//
	// memoryStore test
	Description          *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExtractionStrategies []*string `json:"extractionStrategies,omitempty" xml:"extractionStrategies,omitempty" type:"Repeated"`
	// example:
	//
	// 7
	ShortTermTtl *int32 `json:"shortTermTtl,omitempty" xml:"shortTermTtl,omitempty"`
}

func (s UpdateMemoryStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreRequest) GetCustomExtractionStrategies() []*CustomExtractionStrategy {
	return s.CustomExtractionStrategies
}

func (s *UpdateMemoryStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryStoreRequest) GetExtractionStrategies() []*string {
	return s.ExtractionStrategies
}

func (s *UpdateMemoryStoreRequest) GetShortTermTtl() *int32 {
	return s.ShortTermTtl
}

func (s *UpdateMemoryStoreRequest) SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *UpdateMemoryStoreRequest {
	s.CustomExtractionStrategies = v
	return s
}

func (s *UpdateMemoryStoreRequest) SetDescription(v string) *UpdateMemoryStoreRequest {
	s.Description = &v
	return s
}

func (s *UpdateMemoryStoreRequest) SetExtractionStrategies(v []*string) *UpdateMemoryStoreRequest {
	s.ExtractionStrategies = v
	return s
}

func (s *UpdateMemoryStoreRequest) SetShortTermTtl(v int32) *UpdateMemoryStoreRequest {
	s.ShortTermTtl = &v
	return s
}

func (s *UpdateMemoryStoreRequest) Validate() error {
	if s.CustomExtractionStrategies != nil {
		for _, item := range s.CustomExtractionStrategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
