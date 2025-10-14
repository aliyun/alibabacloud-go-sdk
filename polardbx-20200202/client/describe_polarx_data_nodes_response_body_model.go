// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarxDataNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDataNodes(v []*DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) *DescribePolarxDataNodesResponseBody
	GetDBInstanceDataNodes() []*DescribePolarxDataNodesResponseBodyDBInstanceDataNodes
	SetPageNumber(v int32) *DescribePolarxDataNodesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarxDataNodesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePolarxDataNodesResponseBody
	GetRequestId() *string
	SetTotalNumber(v int32) *DescribePolarxDataNodesResponseBody
	GetTotalNumber() *int32
}

type DescribePolarxDataNodesResponseBody struct {
	DBInstanceDataNodes []*DescribePolarxDataNodesResponseBodyDBInstanceDataNodes `json:"DBInstanceDataNodes,omitempty" xml:"DBInstanceDataNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 0
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BD99340C-4A45-548B-****-27584B0BCFFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribePolarxDataNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarxDataNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesResponseBody) GetDBInstanceDataNodes() []*DescribePolarxDataNodesResponseBodyDBInstanceDataNodes {
	return s.DBInstanceDataNodes
}

func (s *DescribePolarxDataNodesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarxDataNodesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarxDataNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarxDataNodesResponseBody) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *DescribePolarxDataNodesResponseBody) SetDBInstanceDataNodes(v []*DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) *DescribePolarxDataNodesResponseBody {
	s.DBInstanceDataNodes = v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetPageNumber(v int32) *DescribePolarxDataNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetPageSize(v int32) *DescribePolarxDataNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetRequestId(v string) *DescribePolarxDataNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) SetTotalNumber(v int32) *DescribePolarxDataNodesResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBody) Validate() error {
	if s.DBInstanceDataNodes != nil {
		for _, item := range s.DBInstanceDataNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarxDataNodesResponseBodyDBInstanceDataNodes struct {
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// example:
	//
	// pxc-hzrlcjc****sz9
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// pxc-hzrp5m****04w4
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) GoString() string {
	return s.String()
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) SetDBInstanceDescription(v string) *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) SetDBInstanceId(v string) *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) SetDBInstanceName(v string) *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePolarxDataNodesResponseBodyDBInstanceDataNodes) Validate() error {
	return dara.Validate(s)
}
