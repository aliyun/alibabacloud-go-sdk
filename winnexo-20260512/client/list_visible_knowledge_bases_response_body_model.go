// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVisibleKnowledgeBasesResponseBody
	GetCode() *string
	SetItems(v []*ListVisibleKnowledgeBasesResponseBodyItems) *ListVisibleKnowledgeBasesResponseBody
	GetItems() []*ListVisibleKnowledgeBasesResponseBodyItems
	SetMessage(v string) *ListVisibleKnowledgeBasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVisibleKnowledgeBasesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListVisibleKnowledgeBasesResponseBody
	GetTotal() *int64
}

type ListVisibleKnowledgeBasesResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The file information.
	Items []*ListVisibleKnowledgeBasesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 190F5425-A145-5BBA-980F-082ADB0CA6AF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 3
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListVisibleKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetItems() []*ListVisibleKnowledgeBasesResponseBodyItems {
	return s.Items
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVisibleKnowledgeBasesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetCode(v string) *ListVisibleKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetItems(v []*ListVisibleKnowledgeBasesResponseBodyItems) *ListVisibleKnowledgeBasesResponseBody {
	s.Items = v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetMessage(v string) *ListVisibleKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetRequestId(v string) *ListVisibleKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) SetTotal(v int64) *ListVisibleKnowledgeBasesResponseBody {
	s.Total = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBody) Validate() error {
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

type ListVisibleKnowledgeBasesResponseBodyItems struct {
	// The creator.
	//
	// example:
	//
	// admin
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The description.
	//
	// example:
	//
	// {{7*7}}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID. You can obtain this value by calling the API operation for retrieving the knowledge base directory.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time. The value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1763086707000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1774533462
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The skill name.
	//
	// example:
	//
	// cs-default-umodel-1785637207863_k8s.metric.k8s_csi_node_pv_node_cn-heyuan-acdr-1/c80cf3a4f9d6c496781591bd17d006c6f
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The file directory information.
	//
	// example:
	//
	// /test-folder-path1_1773194924773
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The number of resources in the FAILED state. This parameter is returned only when the top-level knowledge base directory list is queried.
	//
	// example:
	//
	// 1
	SourceFailedCount *int64 `json:"sourceFailedCount,omitempty" xml:"sourceFailedCount,omitempty"`
	// The number of resources in the READY state. This parameter is returned only when the top-level knowledge base directory list is queried.
	//
	// example:
	//
	// 1
	SourceReadyCount *int64 `json:"sourceReadyCount,omitempty" xml:"sourceReadyCount,omitempty"`
	// The total number of resources in the directory and its subdirectories. This parameter is returned only when the top-level knowledge base directory list is queried.
	//
	// example:
	//
	// 1
	SourceTotalCount *int64 `json:"sourceTotalCount,omitempty" xml:"sourceTotalCount,omitempty"`
}

func (s ListVisibleKnowledgeBasesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetPath() *string {
	return s.Path
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetSourceFailedCount() *int64 {
	return s.SourceFailedCount
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetSourceReadyCount() *int64 {
	return s.SourceReadyCount
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) GetSourceTotalCount() *int64 {
	return s.SourceTotalCount
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetCreatorName(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetDescription(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetDirectoryId(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.DirectoryId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetGmtCreate(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetGmtModified(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetName(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetPath(v string) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.Path = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetSourceFailedCount(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.SourceFailedCount = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetSourceReadyCount(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.SourceReadyCount = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) SetSourceTotalCount(v int64) *ListVisibleKnowledgeBasesResponseBodyItems {
	s.SourceTotalCount = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
