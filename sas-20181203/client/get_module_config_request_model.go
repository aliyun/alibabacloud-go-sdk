// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *GetModuleConfigRequest
	GetCurrentPage() *string
	SetPageSize(v string) *GetModuleConfigRequest
	GetPageSize() *string
}

type GetModuleConfigRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetModuleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetModuleConfigRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *GetModuleConfigRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetModuleConfigRequest) SetCurrentPage(v string) *GetModuleConfigRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetModuleConfigRequest) SetPageSize(v string) *GetModuleConfigRequest {
	s.PageSize = &v
	return s
}

func (s *GetModuleConfigRequest) Validate() error {
	return dara.Validate(s)
}
