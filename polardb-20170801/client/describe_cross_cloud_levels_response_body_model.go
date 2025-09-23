// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudLevelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrossCloudLevelList(v []*DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) *DescribeCrossCloudLevelsResponseBody
	GetCrossCloudLevelList() []*DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList
	SetRequestId(v string) *DescribeCrossCloudLevelsResponseBody
	GetRequestId() *string
}

type DescribeCrossCloudLevelsResponseBody struct {
	CrossCloudLevelList []*DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList `json:"CrossCloudLevelList,omitempty" xml:"CrossCloudLevelList,omitempty" type:"Repeated"`
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCrossCloudLevelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudLevelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudLevelsResponseBody) GetCrossCloudLevelList() []*DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList {
	return s.CrossCloudLevelList
}

func (s *DescribeCrossCloudLevelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossCloudLevelsResponseBody) SetCrossCloudLevelList(v []*DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) *DescribeCrossCloudLevelsResponseBody {
	s.CrossCloudLevelList = v
	return s
}

func (s *DescribeCrossCloudLevelsResponseBody) SetRequestId(v string) *DescribeCrossCloudLevelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossCloudLevelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList struct {
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// polar.mysql.g4.medium.c
	LevelCode *string `json:"LevelCode,omitempty" xml:"LevelCode,omitempty"`
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
}

func (s DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) GetDBType() *string {
	return s.DBType
}

func (s *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) GetLevelCode() *string {
	return s.LevelCode
}

func (s *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) GetLevelName() *string {
	return s.LevelName
}

func (s *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) SetDBType(v string) *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList {
	s.DBType = &v
	return s
}

func (s *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) SetLevelCode(v string) *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList {
	s.LevelCode = &v
	return s
}

func (s *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) SetLevelName(v string) *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList {
	s.LevelName = &v
	return s
}

func (s *DescribeCrossCloudLevelsResponseBodyCrossCloudLevelList) Validate() error {
	return dara.Validate(s)
}
