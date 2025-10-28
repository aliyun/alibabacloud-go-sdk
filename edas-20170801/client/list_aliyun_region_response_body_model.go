// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliyunRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAliyunRegionResponseBody
	GetCode() *int32
	SetMessage(v string) *ListAliyunRegionResponseBody
	GetMessage() *string
	SetRegionEntityList(v *ListAliyunRegionResponseBodyRegionEntityList) *ListAliyunRegionResponseBody
	GetRegionEntityList() *ListAliyunRegionResponseBodyRegionEntityList
	SetRequestId(v string) *ListAliyunRegionResponseBody
	GetRequestId() *string
}

type ListAliyunRegionResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The details of the regions.
	RegionEntityList *ListAliyunRegionResponseBodyRegionEntityList `json:"RegionEntityList,omitempty" xml:"RegionEntityList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAliyunRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAliyunRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAliyunRegionResponseBody) GetRegionEntityList() *ListAliyunRegionResponseBodyRegionEntityList {
	return s.RegionEntityList
}

func (s *ListAliyunRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAliyunRegionResponseBody) SetCode(v int32) *ListAliyunRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListAliyunRegionResponseBody) SetMessage(v string) *ListAliyunRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListAliyunRegionResponseBody) SetRegionEntityList(v *ListAliyunRegionResponseBodyRegionEntityList) *ListAliyunRegionResponseBody {
	s.RegionEntityList = v
	return s
}

func (s *ListAliyunRegionResponseBody) SetRequestId(v string) *ListAliyunRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliyunRegionResponseBody) Validate() error {
	if s.RegionEntityList != nil {
		if err := s.RegionEntityList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAliyunRegionResponseBodyRegionEntityList struct {
	RegionEntity []*ListAliyunRegionResponseBodyRegionEntityListRegionEntity `json:"RegionEntity,omitempty" xml:"RegionEntity,omitempty" type:"Repeated"`
}

func (s ListAliyunRegionResponseBodyRegionEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunRegionResponseBodyRegionEntityList) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponseBodyRegionEntityList) GetRegionEntity() []*ListAliyunRegionResponseBodyRegionEntityListRegionEntity {
	return s.RegionEntity
}

func (s *ListAliyunRegionResponseBodyRegionEntityList) SetRegionEntity(v []*ListAliyunRegionResponseBodyRegionEntityListRegionEntity) *ListAliyunRegionResponseBodyRegionEntityList {
	s.RegionEntity = v
	return s
}

func (s *ListAliyunRegionResponseBodyRegionEntityList) Validate() error {
	if s.RegionEntity != nil {
		for _, item := range s.RegionEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAliyunRegionResponseBodyRegionEntityListRegionEntity struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China East 1 (Hangzhou)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAliyunRegionResponseBodyRegionEntityListRegionEntity) String() string {
	return dara.Prettify(s)
}

func (s ListAliyunRegionResponseBodyRegionEntityListRegionEntity) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponseBodyRegionEntityListRegionEntity) GetId() *string {
	return s.Id
}

func (s *ListAliyunRegionResponseBodyRegionEntityListRegionEntity) GetName() *string {
	return s.Name
}

func (s *ListAliyunRegionResponseBodyRegionEntityListRegionEntity) SetId(v string) *ListAliyunRegionResponseBodyRegionEntityListRegionEntity {
	s.Id = &v
	return s
}

func (s *ListAliyunRegionResponseBodyRegionEntityListRegionEntity) SetName(v string) *ListAliyunRegionResponseBodyRegionEntityListRegionEntity {
	s.Name = &v
	return s
}

func (s *ListAliyunRegionResponseBodyRegionEntityListRegionEntity) Validate() error {
	return dara.Validate(s)
}
