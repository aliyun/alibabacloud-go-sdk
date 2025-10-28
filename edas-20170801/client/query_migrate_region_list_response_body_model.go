// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMigrateRegionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryMigrateRegionListResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryMigrateRegionListResponseBody
	GetMessage() *string
	SetRegionEntityList(v *QueryMigrateRegionListResponseBodyRegionEntityList) *QueryMigrateRegionListResponseBody
	GetRegionEntityList() *QueryMigrateRegionListResponseBodyRegionEntityList
	SetRequestId(v string) *QueryMigrateRegionListResponseBody
	GetRequestId() *string
}

type QueryMigrateRegionListResponseBody struct {
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
	// The namespaces.
	RegionEntityList *QueryMigrateRegionListResponseBodyRegionEntityList `json:"RegionEntityList,omitempty" xml:"RegionEntityList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMigrateRegionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryMigrateRegionListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMigrateRegionListResponseBody) GetRegionEntityList() *QueryMigrateRegionListResponseBodyRegionEntityList {
	return s.RegionEntityList
}

func (s *QueryMigrateRegionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMigrateRegionListResponseBody) SetCode(v int32) *QueryMigrateRegionListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMigrateRegionListResponseBody) SetMessage(v string) *QueryMigrateRegionListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMigrateRegionListResponseBody) SetRegionEntityList(v *QueryMigrateRegionListResponseBodyRegionEntityList) *QueryMigrateRegionListResponseBody {
	s.RegionEntityList = v
	return s
}

func (s *QueryMigrateRegionListResponseBody) SetRequestId(v string) *QueryMigrateRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMigrateRegionListResponseBody) Validate() error {
	if s.RegionEntityList != nil {
		if err := s.RegionEntityList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMigrateRegionListResponseBodyRegionEntityList struct {
	RegionEntity []*QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity `json:"RegionEntity,omitempty" xml:"RegionEntity,omitempty" type:"Repeated"`
}

func (s QueryMigrateRegionListResponseBodyRegionEntityList) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateRegionListResponseBodyRegionEntityList) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityList) GetRegionEntity() []*QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity {
	return s.RegionEntity
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityList) SetRegionEntity(v []*QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) *QueryMigrateRegionListResponseBodyRegionEntityList {
	s.RegionEntity = v
	return s
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityList) Validate() error {
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

type QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity struct {
	// The name of the namespace.
	//
	// example:
	//
	// Beta
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:beta
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) GetRegionName() *string {
	return s.RegionName
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) GetRegionNo() *string {
	return s.RegionNo
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) SetRegionName(v string) *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity {
	s.RegionName = &v
	return s
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) SetRegionNo(v string) *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity {
	s.RegionNo = &v
	return s
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) Validate() error {
	return dara.Validate(s)
}
