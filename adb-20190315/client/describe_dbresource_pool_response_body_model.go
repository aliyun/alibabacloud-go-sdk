// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourcePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBResourcePoolResponseBody
	GetDBClusterId() *string
	SetPoolsInfo(v []*DescribeDBResourcePoolResponseBodyPoolsInfo) *DescribeDBResourcePoolResponseBody
	GetPoolsInfo() []*DescribeDBResourcePoolResponseBodyPoolsInfo
	SetRequestId(v string) *DescribeDBResourcePoolResponseBody
	GetRequestId() *string
}

type DescribeDBResourcePoolResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Details of the resource group.
	PoolsInfo []*DescribeDBResourcePoolResponseBodyPoolsInfo `json:"PoolsInfo,omitempty" xml:"PoolsInfo,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBResourcePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBResourcePoolResponseBody) GetPoolsInfo() []*DescribeDBResourcePoolResponseBodyPoolsInfo {
	return s.PoolsInfo
}

func (s *DescribeDBResourcePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBResourcePoolResponseBody) SetDBClusterId(v string) *DescribeDBResourcePoolResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBody) SetPoolsInfo(v []*DescribeDBResourcePoolResponseBodyPoolsInfo) *DescribeDBResourcePoolResponseBody {
	s.PoolsInfo = v
	return s
}

func (s *DescribeDBResourcePoolResponseBody) SetRequestId(v string) *DescribeDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBody) Validate() error {
	if s.PoolsInfo != nil {
		for _, item := range s.PoolsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBResourcePoolResponseBodyPoolsInfo struct {
	// The time when the resource group was created.
	//
	// example:
	//
	// 2022-03-09 16:57:35.241
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of nodes.
	//
	// >  Each node consumes 16 cores and 64 GB memory.
	//
	// example:
	//
	// 2
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// USER_DEFAULT
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The database accounts that are associated with the resource group.
	//
	// example:
	//
	// testb,testc
	PoolUsers *string `json:"PoolUsers,omitempty" xml:"PoolUsers,omitempty"`
	// The mode in which SQL statements are executed.
	//
	// 	- **batch**
	//
	// 	- **interactive**
	//
	// >  For more information, see [Query execution modes](https://help.aliyun.com/document_detail/189502.html).
	//
	// example:
	//
	// default_type
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The time when the resource group was updated.
	//
	// example:
	//
	// 2022-03-09 16:57:35.241
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) GetPoolName() *string {
	return s.PoolName
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) GetPoolUsers() *string {
	return s.PoolUsers
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetCreateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetNodeNum(v int32) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.NodeNum = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolName(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolName = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolUsers(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolUsers = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetQueryType(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.QueryType = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetUpdateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) Validate() error {
	return dara.Validate(s)
}
