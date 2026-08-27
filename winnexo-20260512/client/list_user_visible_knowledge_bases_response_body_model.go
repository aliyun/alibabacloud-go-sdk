// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserVisibleKnowledgeBasesResponseBody
	GetCode() *string
	SetItems(v []*ListUserVisibleKnowledgeBasesResponseBodyItems) *ListUserVisibleKnowledgeBasesResponseBody
	GetItems() []*ListUserVisibleKnowledgeBasesResponseBodyItems
	SetMessage(v string) *ListUserVisibleKnowledgeBasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserVisibleKnowledgeBasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUserVisibleKnowledgeBasesResponseBody
	GetTotalCount() *int64
}

type ListUserVisibleKnowledgeBasesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result list.
	//
	// Maximum size:
	//
	// 	50
	Items []*ListUserVisibleKnowledgeBasesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUserVisibleKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) GetItems() []*ListUserVisibleKnowledgeBasesResponseBodyItems {
	return s.Items
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) SetCode(v string) *ListUserVisibleKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) SetItems(v []*ListUserVisibleKnowledgeBasesResponseBodyItems) *ListUserVisibleKnowledgeBasesResponseBody {
	s.Items = v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) SetMessage(v string) *ListUserVisibleKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) SetRequestId(v string) *ListUserVisibleKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) SetTotalCount(v int64) *ListUserVisibleKnowledgeBasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserVisibleKnowledgeBasesResponseBodyItems struct {
	// The user ID of the creator.
	//
	// example:
	//
	// 1
	CreatorId *int64 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// John
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The description of the knowledge base.
	//
	// example:
	//
	// Product materials and usage instructions
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID of the enterprise knowledge base.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The directory type.
	//
	// example:
	//
	// normal
	DirectoryKind *string `json:"directoryKind,omitempty" xml:"directoryKind,omitempty"`
	// The directory type.
	//
	// example:
	//
	// TENANT
	DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The file name.
	//
	// example:
	//
	// Product Knowledge Base
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListUserVisibleKnowledgeBasesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBasesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetDirectoryKind() *string {
	return s.DirectoryKind
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetCreatorId(v int64) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.CreatorId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetCreatorName(v string) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetDescription(v string) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetDirectoryId(v string) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.DirectoryId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetDirectoryKind(v string) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.DirectoryKind = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetDirectoryType(v string) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.DirectoryType = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetGmtCreate(v int64) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetGmtModified(v int64) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) SetName(v string) *ListUserVisibleKnowledgeBasesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
