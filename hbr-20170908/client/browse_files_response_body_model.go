// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowseFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BrowseFilesResponseBody
	GetCode() *string
	SetMaxResults(v int32) *BrowseFilesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *BrowseFilesResponseBody
	GetMessage() *string
	SetNextToken(v string) *BrowseFilesResponseBody
	GetNextToken() *string
	SetNodes(v *BrowseFilesResponseBodyNodes) *BrowseFilesResponseBody
	GetNodes() *BrowseFilesResponseBodyNodes
	SetPageNumber(v int32) *BrowseFilesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *BrowseFilesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *BrowseFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BrowseFilesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *BrowseFilesResponseBody
	GetTotalCount() *int32
}

type BrowseFilesResponseBody struct {
	// The HTTP status code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response message. The value `successful` is returned if the request was successful. An error message is returned if the request failed.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token for the next page of results. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// eyJ***Q==
	NextToken *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Nodes     *BrowseFilesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 843F7A45-8B***35-ECECBE5E5F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s BrowseFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BrowseFilesResponseBody) GoString() string {
	return s.String()
}

func (s *BrowseFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *BrowseFilesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *BrowseFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BrowseFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *BrowseFilesResponseBody) GetNodes() *BrowseFilesResponseBodyNodes {
	return s.Nodes
}

func (s *BrowseFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *BrowseFilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *BrowseFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BrowseFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BrowseFilesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *BrowseFilesResponseBody) SetCode(v string) *BrowseFilesResponseBody {
	s.Code = &v
	return s
}

func (s *BrowseFilesResponseBody) SetMaxResults(v int32) *BrowseFilesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *BrowseFilesResponseBody) SetMessage(v string) *BrowseFilesResponseBody {
	s.Message = &v
	return s
}

func (s *BrowseFilesResponseBody) SetNextToken(v string) *BrowseFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *BrowseFilesResponseBody) SetNodes(v *BrowseFilesResponseBodyNodes) *BrowseFilesResponseBody {
	s.Nodes = v
	return s
}

func (s *BrowseFilesResponseBody) SetPageNumber(v int32) *BrowseFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *BrowseFilesResponseBody) SetPageSize(v int32) *BrowseFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *BrowseFilesResponseBody) SetRequestId(v string) *BrowseFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BrowseFilesResponseBody) SetSuccess(v bool) *BrowseFilesResponseBody {
	s.Success = &v
	return s
}

func (s *BrowseFilesResponseBody) SetTotalCount(v int32) *BrowseFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *BrowseFilesResponseBody) Validate() error {
	if s.Nodes != nil {
		if err := s.Nodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BrowseFilesResponseBodyNodes struct {
	Node []*BrowseFilesResponseBodyNodesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s BrowseFilesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s BrowseFilesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *BrowseFilesResponseBodyNodes) GetNode() []*BrowseFilesResponseBodyNodesNode {
	return s.Node
}

func (s *BrowseFilesResponseBodyNodes) SetNode(v []*BrowseFilesResponseBodyNodesNode) *BrowseFilesResponseBodyNodes {
	s.Node = v
	return s
}

func (s *BrowseFilesResponseBodyNodes) Validate() error {
	if s.Node != nil {
		for _, item := range s.Node {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BrowseFilesResponseBodyNodesNode struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Subtree *string `json:"Subtree,omitempty" xml:"Subtree,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BrowseFilesResponseBodyNodesNode) String() string {
	return dara.Prettify(s)
}

func (s BrowseFilesResponseBodyNodesNode) GoString() string {
	return s.String()
}

func (s *BrowseFilesResponseBodyNodesNode) GetName() *string {
	return s.Name
}

func (s *BrowseFilesResponseBodyNodesNode) GetSize() *int64 {
	return s.Size
}

func (s *BrowseFilesResponseBodyNodesNode) GetSubtree() *string {
	return s.Subtree
}

func (s *BrowseFilesResponseBodyNodesNode) GetType() *string {
	return s.Type
}

func (s *BrowseFilesResponseBodyNodesNode) SetName(v string) *BrowseFilesResponseBodyNodesNode {
	s.Name = &v
	return s
}

func (s *BrowseFilesResponseBodyNodesNode) SetSize(v int64) *BrowseFilesResponseBodyNodesNode {
	s.Size = &v
	return s
}

func (s *BrowseFilesResponseBodyNodesNode) SetSubtree(v string) *BrowseFilesResponseBodyNodesNode {
	s.Subtree = &v
	return s
}

func (s *BrowseFilesResponseBodyNodesNode) SetType(v string) *BrowseFilesResponseBodyNodesNode {
	s.Type = &v
	return s
}

func (s *BrowseFilesResponseBodyNodesNode) Validate() error {
	return dara.Validate(s)
}
