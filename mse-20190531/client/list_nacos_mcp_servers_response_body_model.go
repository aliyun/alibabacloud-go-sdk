// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosMcpServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListNacosMcpServersResponseBodyData) *ListNacosMcpServersResponseBody
	GetData() *ListNacosMcpServersResponseBodyData
	SetRequestId(v string) *ListNacosMcpServersResponseBody
	GetRequestId() *string
}

type ListNacosMcpServersResponseBody struct {
	Data *ListNacosMcpServersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ADDD8AB7-8D1C-4697-A83E-410D2607****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNacosMcpServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNacosMcpServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListNacosMcpServersResponseBody) GetData() *ListNacosMcpServersResponseBodyData {
	return s.Data
}

func (s *ListNacosMcpServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNacosMcpServersResponseBody) SetData(v *ListNacosMcpServersResponseBodyData) *ListNacosMcpServersResponseBody {
	s.Data = v
	return s
}

func (s *ListNacosMcpServersResponseBody) SetRequestId(v string) *ListNacosMcpServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNacosMcpServersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNacosMcpServersResponseBodyData struct {
	PageItems []*ListNacosMcpServersResponseBodyDataPageItems `json:"PageItems,omitempty" xml:"PageItems,omitempty" type:"Repeated"`
	// pageNumber.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pagesAvailable.
	//
	// example:
	//
	// 1
	PagesAvailable *int32 `json:"PagesAvailable,omitempty" xml:"PagesAvailable,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNacosMcpServersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNacosMcpServersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNacosMcpServersResponseBodyData) GetPageItems() []*ListNacosMcpServersResponseBodyDataPageItems {
	return s.PageItems
}

func (s *ListNacosMcpServersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNacosMcpServersResponseBodyData) GetPagesAvailable() *int32 {
	return s.PagesAvailable
}

func (s *ListNacosMcpServersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNacosMcpServersResponseBodyData) SetPageItems(v []*ListNacosMcpServersResponseBodyDataPageItems) *ListNacosMcpServersResponseBodyData {
	s.PageItems = v
	return s
}

func (s *ListNacosMcpServersResponseBodyData) SetPageNumber(v int32) *ListNacosMcpServersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyData) SetPagesAvailable(v int32) *ListNacosMcpServersResponseBodyData {
	s.PagesAvailable = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyData) SetTotalCount(v int32) *ListNacosMcpServersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyData) Validate() error {
	if s.PageItems != nil {
		for _, item := range s.PageItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNacosMcpServersResponseBodyDataPageItems struct {
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mcp-sse
	FrontProtocol *string `json:"FrontProtocol,omitempty" xml:"FrontProtocol,omitempty"`
	// ID。
	//
	// example:
	//
	// 2385420b-6176-4a37-aba4-d3d382e4bb84
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// stdio
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 1.1.0
	Version       *string                                                    `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionDetail *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail `json:"VersionDetail,omitempty" xml:"VersionDetail,omitempty" type:"Struct"`
}

func (s ListNacosMcpServersResponseBodyDataPageItems) String() string {
	return dara.Prettify(s)
}

func (s ListNacosMcpServersResponseBodyDataPageItems) GoString() string {
	return s.String()
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetDescription() *string {
	return s.Description
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetFrontProtocol() *string {
	return s.FrontProtocol
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetId() *string {
	return s.Id
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetName() *string {
	return s.Name
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetVersion() *string {
	return s.Version
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) GetVersionDetail() *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail {
	return s.VersionDetail
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetCapabilities(v []*string) *ListNacosMcpServersResponseBodyDataPageItems {
	s.Capabilities = v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetDescription(v string) *ListNacosMcpServersResponseBodyDataPageItems {
	s.Description = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetFrontProtocol(v string) *ListNacosMcpServersResponseBodyDataPageItems {
	s.FrontProtocol = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetId(v string) *ListNacosMcpServersResponseBodyDataPageItems {
	s.Id = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetName(v string) *ListNacosMcpServersResponseBodyDataPageItems {
	s.Name = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetProtocol(v string) *ListNacosMcpServersResponseBodyDataPageItems {
	s.Protocol = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetVersion(v string) *ListNacosMcpServersResponseBodyDataPageItems {
	s.Version = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) SetVersionDetail(v *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) *ListNacosMcpServersResponseBodyDataPageItems {
	s.VersionDetail = v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItems) Validate() error {
	if s.VersionDetail != nil {
		if err := s.VersionDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNacosMcpServersResponseBodyDataPageItemsVersionDetail struct {
	// example:
	//
	// true
	IsLatest *bool `json:"Is_latest,omitempty" xml:"Is_latest,omitempty"`
	// example:
	//
	// 2025-07-16
	ReleaseDate *string `json:"Release_date,omitempty" xml:"Release_date,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) GoString() string {
	return s.String()
}

func (s *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) GetIsLatest() *bool {
	return s.IsLatest
}

func (s *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) GetVersion() *string {
	return s.Version
}

func (s *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) SetIsLatest(v bool) *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail {
	s.IsLatest = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) SetReleaseDate(v string) *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail {
	s.ReleaseDate = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) SetVersion(v string) *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail {
	s.Version = &v
	return s
}

func (s *ListNacosMcpServersResponseBodyDataPageItemsVersionDetail) Validate() error {
	return dara.Validate(s)
}
