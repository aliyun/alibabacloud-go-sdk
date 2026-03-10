// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetMemoryStoreResponseBody
	GetCreateTime() *string
	SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *GetMemoryStoreResponseBody
	GetCustomExtractionStrategies() []*CustomExtractionStrategy
	SetDescription(v string) *GetMemoryStoreResponseBody
	GetDescription() *string
	SetExtractionStrategies(v []*string) *GetMemoryStoreResponseBody
	GetExtractionStrategies() []*string
	SetMemoryStoreName(v string) *GetMemoryStoreResponseBody
	GetMemoryStoreName() *string
	SetRegionId(v string) *GetMemoryStoreResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetMemoryStoreResponseBody
	GetRequestId() *string
	SetShortTermTtl(v int32) *GetMemoryStoreResponseBody
	GetShortTermTtl() *int32
	SetUpdateTime(v string) *GetMemoryStoreResponseBody
	GetUpdateTime() *string
	SetWorkspace(v string) *GetMemoryStoreResponseBody
	GetWorkspace() *string
}

type GetMemoryStoreResponseBody struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1764556182850
	CreateTime                 *string                     `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CustomExtractionStrategies []*CustomExtractionStrategy `json:"customExtractionStrategies,omitempty" xml:"customExtractionStrategies,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description          *string   `json:"description,omitempty" xml:"description,omitempty"`
	ExtractionStrategies []*string `json:"extractionStrategies,omitempty" xml:"extractionStrategies,omitempty" type:"Repeated"`
	// example:
	//
	// test-memory-store
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	ShortTermTtl *int32 `json:"shortTermTtl,omitempty" xml:"shortTermTtl,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1764556182850
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// default-cms-xxxxxx-cn-beijing
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMemoryStoreResponseBody) GetCustomExtractionStrategies() []*CustomExtractionStrategy {
	return s.CustomExtractionStrategies
}

func (s *GetMemoryStoreResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMemoryStoreResponseBody) GetExtractionStrategies() []*string {
	return s.ExtractionStrategies
}

func (s *GetMemoryStoreResponseBody) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *GetMemoryStoreResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryStoreResponseBody) GetShortTermTtl() *int32 {
	return s.ShortTermTtl
}

func (s *GetMemoryStoreResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMemoryStoreResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetMemoryStoreResponseBody) SetCreateTime(v string) *GetMemoryStoreResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetCustomExtractionStrategies(v []*CustomExtractionStrategy) *GetMemoryStoreResponseBody {
	s.CustomExtractionStrategies = v
	return s
}

func (s *GetMemoryStoreResponseBody) SetDescription(v string) *GetMemoryStoreResponseBody {
	s.Description = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetExtractionStrategies(v []*string) *GetMemoryStoreResponseBody {
	s.ExtractionStrategies = v
	return s
}

func (s *GetMemoryStoreResponseBody) SetMemoryStoreName(v string) *GetMemoryStoreResponseBody {
	s.MemoryStoreName = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetRegionId(v string) *GetMemoryStoreResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetRequestId(v string) *GetMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetShortTermTtl(v int32) *GetMemoryStoreResponseBody {
	s.ShortTermTtl = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetUpdateTime(v string) *GetMemoryStoreResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetMemoryStoreResponseBody) SetWorkspace(v string) *GetMemoryStoreResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetMemoryStoreResponseBody) Validate() error {
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
