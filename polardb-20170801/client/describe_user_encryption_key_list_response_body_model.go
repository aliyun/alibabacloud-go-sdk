// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeUserEncryptionKeyListResponseBody
	GetDBClusterId() *string
	SetKeyList(v []*string) *DescribeUserEncryptionKeyListResponseBody
	GetKeyList() []*string
	SetPageNumber(v int32) *DescribeUserEncryptionKeyListResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeUserEncryptionKeyListResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeUserEncryptionKeyListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeUserEncryptionKeyListResponseBody
	GetTotalRecordCount() *int32
}

type DescribeUserEncryptionKeyListResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Cluster key list.
	KeyList []*string `json:"KeyList,omitempty" xml:"KeyList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetKeyList() []*string {
	return s.KeyList
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetDBClusterId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetKeyList(v []*string) *DescribeUserEncryptionKeyListResponseBody {
	s.KeyList = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetPageNumber(v int32) *DescribeUserEncryptionKeyListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetPageRecordCount(v int32) *DescribeUserEncryptionKeyListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetPageSize(v int32) *DescribeUserEncryptionKeyListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetTotalRecordCount(v int32) *DescribeUserEncryptionKeyListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) Validate() error {
	return dara.Validate(s)
}
