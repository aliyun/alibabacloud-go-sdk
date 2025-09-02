// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeTypeListInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeTypeInfoList(v *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) *GetNodeTypeListInfoResponseBody
	GetNodeTypeInfoList() *GetNodeTypeListInfoResponseBodyNodeTypeInfoList
	SetRequestId(v string) *GetNodeTypeListInfoResponseBody
	GetRequestId() *string
}

type GetNodeTypeListInfoResponseBody struct {
	// The list of node types.
	NodeTypeInfoList *GetNodeTypeListInfoResponseBodyNodeTypeInfoList `json:"NodeTypeInfoList,omitempty" xml:"NodeTypeInfoList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeTypeListInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeTypeListInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeTypeListInfoResponseBody) GetNodeTypeInfoList() *GetNodeTypeListInfoResponseBodyNodeTypeInfoList {
	return s.NodeTypeInfoList
}

func (s *GetNodeTypeListInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeTypeListInfoResponseBody) SetNodeTypeInfoList(v *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) *GetNodeTypeListInfoResponseBody {
	s.NodeTypeInfoList = v
	return s
}

func (s *GetNodeTypeListInfoResponseBody) SetRequestId(v string) *GetNodeTypeListInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeTypeListInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNodeTypeListInfoResponseBodyNodeTypeInfoList struct {
	// The information about a node type.
	NodeTypeInfo []*GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo `json:"NodeTypeInfo,omitempty" xml:"NodeTypeInfo,omitempty" type:"Repeated"`
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

func (s GetNodeTypeListInfoResponseBodyNodeTypeInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetNodeTypeListInfoResponseBodyNodeTypeInfoList) GoString() string {
	return s.String()
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) GetNodeTypeInfo() []*GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo {
	return s.NodeTypeInfo
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) SetNodeTypeInfo(v []*GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) *GetNodeTypeListInfoResponseBodyNodeTypeInfoList {
	s.NodeTypeInfo = v
	return s
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) SetPageNumber(v int32) *GetNodeTypeListInfoResponseBodyNodeTypeInfoList {
	s.PageNumber = &v
	return s
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) SetPageSize(v int32) *GetNodeTypeListInfoResponseBodyNodeTypeInfoList {
	s.PageSize = &v
	return s
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) SetTotalCount(v int32) *GetNodeTypeListInfoResponseBodyNodeTypeInfoList {
	s.TotalCount = &v
	return s
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoList) Validate() error {
	return dara.Validate(s)
}

type GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo struct {
	// The type of the node, which is specified by a number.
	//
	// example:
	//
	// 10
	NodeType *int32 `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The name of the node type.
	//
	// example:
	//
	// ODPS SQL
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
}

func (s GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) GoString() string {
	return s.String()
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) GetNodeType() *int32 {
	return s.NodeType
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) GetNodeTypeName() *string {
	return s.NodeTypeName
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) SetNodeType(v int32) *GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo {
	s.NodeType = &v
	return s
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) SetNodeTypeName(v string) *GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo {
	s.NodeTypeName = &v
	return s
}

func (s *GetNodeTypeListInfoResponseBodyNodeTypeInfoListNodeTypeInfo) Validate() error {
	return dara.Validate(s)
}
