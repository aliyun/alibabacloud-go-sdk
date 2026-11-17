// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFilterTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListWebFilterTemplatesRequest
	GetDescription() *string
	SetName(v string) *ListWebFilterTemplatesRequest
	GetName() *string
	SetPageNo(v int64) *ListWebFilterTemplatesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListWebFilterTemplatesRequest
	GetPageSize() *int64
	SetStatus(v string) *ListWebFilterTemplatesRequest
	GetStatus() *string
	SetWebFilterTemplateIds(v []*string) *ListWebFilterTemplatesRequest
	GetWebFilterTemplateIds() []*string
}

type ListWebFilterTemplatesRequest struct {
	// The description of the web filtering template. Fuzzy match is supported.
	//
	// example:
	//
	// URL access control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the web filtering template. Fuzzy match is supported.
	//
	// example:
	//
	// Web filtering template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the web filtering template. Exact match is supported. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of the web filtering templates for batch exact matching. You can specify up to 100 template IDs.
	WebFilterTemplateIds []*string `json:"WebFilterTemplateIds,omitempty" xml:"WebFilterTemplateIds,omitempty" type:"Repeated"`
}

func (s ListWebFilterTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListWebFilterTemplatesRequest) GetDescription() *string {
	return s.Description
}

func (s *ListWebFilterTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListWebFilterTemplatesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListWebFilterTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWebFilterTemplatesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWebFilterTemplatesRequest) GetWebFilterTemplateIds() []*string {
	return s.WebFilterTemplateIds
}

func (s *ListWebFilterTemplatesRequest) SetDescription(v string) *ListWebFilterTemplatesRequest {
	s.Description = &v
	return s
}

func (s *ListWebFilterTemplatesRequest) SetName(v string) *ListWebFilterTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListWebFilterTemplatesRequest) SetPageNo(v int64) *ListWebFilterTemplatesRequest {
	s.PageNo = &v
	return s
}

func (s *ListWebFilterTemplatesRequest) SetPageSize(v int64) *ListWebFilterTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWebFilterTemplatesRequest) SetStatus(v string) *ListWebFilterTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListWebFilterTemplatesRequest) SetWebFilterTemplateIds(v []*string) *ListWebFilterTemplatesRequest {
	s.WebFilterTemplateIds = v
	return s
}

func (s *ListWebFilterTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
