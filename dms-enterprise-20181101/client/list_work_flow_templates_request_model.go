// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchName(v string) *ListWorkFlowTemplatesRequest
	GetSearchName() *string
	SetTid(v int64) *ListWorkFlowTemplatesRequest
	GetTid() *int64
}

type ListWorkFlowTemplatesRequest struct {
	// The name that is used to query approval templates.
	//
	// example:
	//
	// Admin
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListWorkFlowTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesRequest) GetSearchName() *string {
	return s.SearchName
}

func (s *ListWorkFlowTemplatesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListWorkFlowTemplatesRequest) SetSearchName(v string) *ListWorkFlowTemplatesRequest {
	s.SearchName = &v
	return s
}

func (s *ListWorkFlowTemplatesRequest) SetTid(v int64) *ListWorkFlowTemplatesRequest {
	s.Tid = &v
	return s
}

func (s *ListWorkFlowTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
