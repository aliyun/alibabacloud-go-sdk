// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensNodeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *QuerySensNodeInfoRequest
	GetNodeId() *string
	SetPageNo(v int32) *QuerySensNodeInfoRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QuerySensNodeInfoRequest
	GetPageSize() *int32
	SetSensitiveName(v string) *QuerySensNodeInfoRequest
	GetSensitiveName() *string
	SetTemplateId(v string) *QuerySensNodeInfoRequest
	GetTemplateId() *string
	SetTenantId(v string) *QuerySensNodeInfoRequest
	GetTenantId() *string
	SetStatus(v int32) *QuerySensNodeInfoRequest
	GetStatus() *int32
}

type QuerySensNodeInfoRequest struct {
	// The ID of the data category. You can call the [QuerySensClassification](https://help.aliyun.com/document_detail/2746850.html) operation or log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Data Security Guard page to obtain the ID.
	//
	// example:
	//
	// 0ce67949-0810-400f-a24a-cc5ffafe1024
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 10 to 1000. The recommended number of entries per page ranges from 10 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the sensitive field. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Data Security Guard page to obtain the name.
	//
	// example:
	//
	// ID card
	SensitiveName *string `json:"SensitiveName,omitempty" xml:"SensitiveName,omitempty"`
	// The ID of the data category and data sensitivity level template. You can call the [QueryDefaultTemplate](https://help.aliyun.com/document_detail/2743948.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1970541-6cf5-4d23-b101-d5b66f6e1024
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The status of the sensitive field. Valid values:
	//
	// 	- 0: draft
	//
	// 	- 1: published
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QuerySensNodeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySensNodeInfoRequest) GoString() string {
	return s.String()
}

func (s *QuerySensNodeInfoRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *QuerySensNodeInfoRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QuerySensNodeInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySensNodeInfoRequest) GetSensitiveName() *string {
	return s.SensitiveName
}

func (s *QuerySensNodeInfoRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QuerySensNodeInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySensNodeInfoRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QuerySensNodeInfoRequest) SetNodeId(v string) *QuerySensNodeInfoRequest {
	s.NodeId = &v
	return s
}

func (s *QuerySensNodeInfoRequest) SetPageNo(v int32) *QuerySensNodeInfoRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySensNodeInfoRequest) SetPageSize(v int32) *QuerySensNodeInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySensNodeInfoRequest) SetSensitiveName(v string) *QuerySensNodeInfoRequest {
	s.SensitiveName = &v
	return s
}

func (s *QuerySensNodeInfoRequest) SetTemplateId(v string) *QuerySensNodeInfoRequest {
	s.TemplateId = &v
	return s
}

func (s *QuerySensNodeInfoRequest) SetTenantId(v string) *QuerySensNodeInfoRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySensNodeInfoRequest) SetStatus(v int32) *QuerySensNodeInfoRequest {
	s.Status = &v
	return s
}

func (s *QuerySensNodeInfoRequest) Validate() error {
	return dara.Validate(s)
}
