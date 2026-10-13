// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncSourceTreeNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBody
	GetCode() *string
	SetData(v []*ListContextDatabaseSyncSourceTreeNodesResponseBodyData) *ListContextDatabaseSyncSourceTreeNodesResponseBody
	GetData() []*ListContextDatabaseSyncSourceTreeNodesResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncSourceTreeNodesResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncSourceTreeNodesResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncSourceTreeNodesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListContextDatabaseSyncSourceTreeNodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Example value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncSourceTreeNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourceTreeNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) GetData() []*ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) SetCode(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) SetData(v []*ListContextDatabaseSyncSourceTreeNodesResponseBodyData) *ListContextDatabaseSyncSourceTreeNodesResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncSourceTreeNodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) SetMessage(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) SetRequestId(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncSourceTreeNodesResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncSourceTreeNodesResponseBodyData struct {
	// The discovery time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	DiscoveredAt *string `json:"DiscoveredAt,omitempty" xml:"DiscoveredAt,omitempty"`
	// The file name extension.
	//
	// example:
	//
	// .docx
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Indicates whether the node has child nodes.
	//
	// example:
	//
	// true
	HasChildren *bool `json:"HasChildren,omitempty" xml:"HasChildren,omitempty"`
	// The type.
	//
	// example:
	//
	// folder
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The node identifier.
	//
	// example:
	//
	// doc:1234567890
	NodeKey *string `json:"NodeKey,omitempty" xml:"NodeKey,omitempty"`
	// The parent node identifier.
	//
	// example:
	//
	// folder:9876543210
	ParentKey *string `json:"ParentKey,omitempty" xml:"ParentKey,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	StaleAt *string `json:"StaleAt,omitempty" xml:"StaleAt,omitempty"`
	// The title.
	//
	// example:
	//
	// Project plan
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The reason why the operation is not supported.
	//
	// example:
	//
	// The current plugin does not support this operation
	UnsupportedReason *string `json:"UnsupportedReason,omitempty" xml:"UnsupportedReason,omitempty"`
}

func (s ListContextDatabaseSyncSourceTreeNodesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetDiscoveredAt() *string {
	return s.DiscoveredAt
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetKind() *string {
	return s.Kind
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetNodeKey() *string {
	return s.NodeKey
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetParentKey() *string {
	return s.ParentKey
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetStaleAt() *string {
	return s.StaleAt
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) GetUnsupportedReason() *string {
	return s.UnsupportedReason
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetDiscoveredAt(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.DiscoveredAt = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetExtension(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetHasChildren(v bool) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.HasChildren = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetKind(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.Kind = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetNodeKey(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.NodeKey = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetParentKey(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.ParentKey = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetStaleAt(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.StaleAt = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetTitle(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) SetUnsupportedReason(v string) *ListContextDatabaseSyncSourceTreeNodesResponseBodyData {
	s.UnsupportedReason = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
