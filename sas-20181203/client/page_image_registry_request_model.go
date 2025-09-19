// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageImageRegistryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *PageImageRegistryRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *PageImageRegistryRequest
	GetPageSize() *int32
	SetRegistryNameLike(v string) *PageImageRegistryRequest
	GetRegistryNameLike() *string
	SetRegistryTypeInList(v []*string) *PageImageRegistryRequest
	GetRegistryTypeInList() []*string
	SetRegistryTypeNotInList(v []*string) *PageImageRegistryRequest
	GetRegistryTypeNotInList() []*string
	SetSourceIp(v string) *PageImageRegistryRequest
	GetSourceIp() *string
}

type PageImageRegistryRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the image repository. Fuzzy match is supported.
	//
	// example:
	//
	// asa
	RegistryNameLike *string `json:"RegistryNameLike,omitempty" xml:"RegistryNameLike,omitempty"`
	// The types of image repositories.
	RegistryTypeInList []*string `json:"RegistryTypeInList,omitempty" xml:"RegistryTypeInList,omitempty" type:"Repeated"`
	// The types of excluded image repositories.
	RegistryTypeNotInList []*string `json:"RegistryTypeNotInList,omitempty" xml:"RegistryTypeNotInList,omitempty" type:"Repeated"`
	// The source IP address of the request.
	//
	// example:
	//
	// 140.207.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PageImageRegistryRequest) String() string {
	return dara.Prettify(s)
}

func (s PageImageRegistryRequest) GoString() string {
	return s.String()
}

func (s *PageImageRegistryRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *PageImageRegistryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageImageRegistryRequest) GetRegistryNameLike() *string {
	return s.RegistryNameLike
}

func (s *PageImageRegistryRequest) GetRegistryTypeInList() []*string {
	return s.RegistryTypeInList
}

func (s *PageImageRegistryRequest) GetRegistryTypeNotInList() []*string {
	return s.RegistryTypeNotInList
}

func (s *PageImageRegistryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PageImageRegistryRequest) SetCurrentPage(v int32) *PageImageRegistryRequest {
	s.CurrentPage = &v
	return s
}

func (s *PageImageRegistryRequest) SetPageSize(v int32) *PageImageRegistryRequest {
	s.PageSize = &v
	return s
}

func (s *PageImageRegistryRequest) SetRegistryNameLike(v string) *PageImageRegistryRequest {
	s.RegistryNameLike = &v
	return s
}

func (s *PageImageRegistryRequest) SetRegistryTypeInList(v []*string) *PageImageRegistryRequest {
	s.RegistryTypeInList = v
	return s
}

func (s *PageImageRegistryRequest) SetRegistryTypeNotInList(v []*string) *PageImageRegistryRequest {
	s.RegistryTypeNotInList = v
	return s
}

func (s *PageImageRegistryRequest) SetSourceIp(v string) *PageImageRegistryRequest {
	s.SourceIp = &v
	return s
}

func (s *PageImageRegistryRequest) Validate() error {
	return dara.Validate(s)
}
