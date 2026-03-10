// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *CreateMemoryStoreRequest
	GetCustomExtractionStrategies() []*CustomExtractionStrategy
	SetDescription(v string) *CreateMemoryStoreRequest
	GetDescription() *string
	SetExtractionStrategies(v []*string) *CreateMemoryStoreRequest
	GetExtractionStrategies() []*string
	SetMemoryStoreName(v string) *CreateMemoryStoreRequest
	GetMemoryStoreName() *string
	SetShortTermTtl(v int32) *CreateMemoryStoreRequest
	GetShortTermTtl() *int32
}

type CreateMemoryStoreRequest struct {
	CustomExtractionStrategies []*CustomExtractionStrategy `json:"customExtractionStrategies,omitempty" xml:"customExtractionStrategies,omitempty" type:"Repeated"`
	// example:
	//
	// Test memory store for demonstration.
	Description          *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExtractionStrategies []*string `json:"extractionStrategies,omitempty" xml:"extractionStrategies,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-memory-store
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	ShortTermTtl *int32 `json:"shortTermTtl,omitempty" xml:"shortTermTtl,omitempty"`
}

func (s CreateMemoryStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreRequest) GetCustomExtractionStrategies() []*CustomExtractionStrategy {
	return s.CustomExtractionStrategies
}

func (s *CreateMemoryStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryStoreRequest) GetExtractionStrategies() []*string {
	return s.ExtractionStrategies
}

func (s *CreateMemoryStoreRequest) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *CreateMemoryStoreRequest) GetShortTermTtl() *int32 {
	return s.ShortTermTtl
}

func (s *CreateMemoryStoreRequest) SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *CreateMemoryStoreRequest {
	s.CustomExtractionStrategies = v
	return s
}

func (s *CreateMemoryStoreRequest) SetDescription(v string) *CreateMemoryStoreRequest {
	s.Description = &v
	return s
}

func (s *CreateMemoryStoreRequest) SetExtractionStrategies(v []*string) *CreateMemoryStoreRequest {
	s.ExtractionStrategies = v
	return s
}

func (s *CreateMemoryStoreRequest) SetMemoryStoreName(v string) *CreateMemoryStoreRequest {
	s.MemoryStoreName = &v
	return s
}

func (s *CreateMemoryStoreRequest) SetShortTermTtl(v int32) *CreateMemoryStoreRequest {
	s.ShortTermTtl = &v
	return s
}

func (s *CreateMemoryStoreRequest) Validate() error {
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
