// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeTypeInfoList(v *ListFileTypeResponseBodyNodeTypeInfoList) *ListFileTypeResponseBody
	GetNodeTypeInfoList() *ListFileTypeResponseBodyNodeTypeInfoList
	SetRequestId(v string) *ListFileTypeResponseBody
	GetRequestId() *string
}

type ListFileTypeResponseBody struct {
	// The information about node types.
	NodeTypeInfoList *ListFileTypeResponseBodyNodeTypeInfoList `json:"NodeTypeInfoList,omitempty" xml:"NodeTypeInfoList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileTypeResponseBody) GetNodeTypeInfoList() *ListFileTypeResponseBodyNodeTypeInfoList {
	return s.NodeTypeInfoList
}

func (s *ListFileTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileTypeResponseBody) SetNodeTypeInfoList(v *ListFileTypeResponseBodyNodeTypeInfoList) *ListFileTypeResponseBody {
	s.NodeTypeInfoList = v
	return s
}

func (s *ListFileTypeResponseBody) SetRequestId(v string) *ListFileTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFileTypeResponseBodyNodeTypeInfoList struct {
	// The information about the node type.
	NodeTypeInfo []*ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo `json:"NodeTypeInfo,omitempty" xml:"NodeTypeInfo,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 127
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileTypeResponseBodyNodeTypeInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListFileTypeResponseBodyNodeTypeInfoList) GoString() string {
	return s.String()
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) GetNodeTypeInfo() []*ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo {
	return s.NodeTypeInfo
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) SetNodeTypeInfo(v []*ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) *ListFileTypeResponseBodyNodeTypeInfoList {
	s.NodeTypeInfo = v
	return s
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) SetPageNumber(v int32) *ListFileTypeResponseBodyNodeTypeInfoList {
	s.PageNumber = &v
	return s
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) SetPageSize(v int32) *ListFileTypeResponseBodyNodeTypeInfoList {
	s.PageSize = &v
	return s
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) SetTotalCount(v int32) *ListFileTypeResponseBodyNodeTypeInfoList {
	s.TotalCount = &v
	return s
}

func (s *ListFileTypeResponseBodyNodeTypeInfoList) Validate() error {
	return dara.Validate(s)
}

type ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo struct {
	// The code of the node type. The codes and names of node types have the following mappings: 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 23 (Data Integration), 24 (ODPS Script), 99 (zero load), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (real-time synchronization), 1089 (cross-tenant collaboration), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (assignment), and 1221 (PyODPS 3)
	//
	// example:
	//
	// 10
	NodeType *int32 `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The name of the node type. The codes and names of node types have the following mappings: 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 23 (Data Integration), 24 (ODPS Script), 99 (zero load), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (real-time synchronization), 1089 (cross-tenant collaboration), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (assignment), and 1221 (PyODPS 3)
	//
	// example:
	//
	// ODPS SQL
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
}

func (s ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) GoString() string {
	return s.String()
}

func (s *ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) GetNodeType() *int32 {
	return s.NodeType
}

func (s *ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) GetNodeTypeName() *string {
	return s.NodeTypeName
}

func (s *ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) SetNodeType(v int32) *ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo {
	s.NodeType = &v
	return s
}

func (s *ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) SetNodeTypeName(v string) *ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo {
	s.NodeTypeName = &v
	return s
}

func (s *ListFileTypeResponseBodyNodeTypeInfoListNodeTypeInfo) Validate() error {
	return dara.Validate(s)
}
