// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShowStorageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeShowStorageInfoResponseBody
	GetCode() *string
	SetData(v *DescribeShowStorageInfoResponseBodyData) *DescribeShowStorageInfoResponseBody
	GetData() *DescribeShowStorageInfoResponseBodyData
	SetRequestId(v string) *DescribeShowStorageInfoResponseBody
	GetRequestId() *string
}

type DescribeShowStorageInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeShowStorageInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeShowStorageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowStorageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShowStorageInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeShowStorageInfoResponseBody) GetData() *DescribeShowStorageInfoResponseBodyData {
	return s.Data
}

func (s *DescribeShowStorageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeShowStorageInfoResponseBody) SetCode(v string) *DescribeShowStorageInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBody) SetData(v *DescribeShowStorageInfoResponseBodyData) *DescribeShowStorageInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeShowStorageInfoResponseBody) SetRequestId(v string) *DescribeShowStorageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeShowStorageInfoResponseBodyData struct {
	StorageInfos []*DescribeShowStorageInfoResponseBodyDataStorageInfos `json:"StorageInfos,omitempty" xml:"StorageInfos,omitempty" type:"Repeated"`
}

func (s DescribeShowStorageInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowStorageInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeShowStorageInfoResponseBodyData) GetStorageInfos() []*DescribeShowStorageInfoResponseBodyDataStorageInfos {
	return s.StorageInfos
}

func (s *DescribeShowStorageInfoResponseBodyData) SetStorageInfos(v []*DescribeShowStorageInfoResponseBodyDataStorageInfos) *DescribeShowStorageInfoResponseBodyData {
	s.StorageInfos = v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyData) Validate() error {
	if s.StorageInfos != nil {
		for _, item := range s.StorageInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeShowStorageInfoResponseBodyDataStorageInfos struct {
	// example:
	//
	// polar.mysql.x4.large
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// example:
	//
	// 3
	DbCount *int64 `json:"DbCount,omitempty" xml:"DbCount,omitempty"`
	// example:
	//
	// True
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// example:
	//
	// 12
	GroupCount *int64 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// example:
	//
	// MASTER
	InstKind *string `json:"InstKind,omitempty" xml:"InstKind,omitempty"`
	// example:
	//
	// true
	IsHealthy *bool `json:"IsHealthy,omitempty" xml:"IsHealthy,omitempty"`
	// example:
	//
	// 11.117.237.205:3029
	LeaderNode *string `json:"LeaderNode,omitempty" xml:"LeaderNode,omitempty"`
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// pxc-xdb-m-pxcbjrylg49skcxb17394
	StorageInstName *string `json:"StorageInstName,omitempty" xml:"StorageInstName,omitempty"`
}

func (s DescribeShowStorageInfoResponseBodyDataStorageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowStorageInfoResponseBodyDataStorageInfos) GoString() string {
	return s.String()
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetClass() *string {
	return s.Class
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetDbCount() *int64 {
	return s.DbCount
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetDeletable() *bool {
	return s.Deletable
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetGroupCount() *int64 {
	return s.GroupCount
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetInstKind() *string {
	return s.InstKind
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetIsHealthy() *bool {
	return s.IsHealthy
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetLeaderNode() *string {
	return s.LeaderNode
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) GetStorageInstName() *string {
	return s.StorageInstName
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetClass(v string) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.Class = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetDbCount(v int64) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.DbCount = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetDeletable(v bool) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.Deletable = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetGroupCount(v int64) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.GroupCount = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetInstKind(v string) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.InstKind = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetIsHealthy(v bool) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.IsHealthy = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetLeaderNode(v string) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.LeaderNode = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetStatus(v int64) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.Status = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) SetStorageInstName(v string) *DescribeShowStorageInfoResponseBodyDataStorageInfos {
	s.StorageInstName = &v
	return s
}

func (s *DescribeShowStorageInfoResponseBodyDataStorageInfos) Validate() error {
	return dara.Validate(s)
}
