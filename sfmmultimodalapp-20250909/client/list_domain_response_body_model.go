// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainInfoList(v []*ListDomainResponseBodyDomainInfoList) *ListDomainResponseBody
	GetDomainInfoList() []*ListDomainResponseBodyDomainInfoList
	SetPageNumber(v int32) *ListDomainResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDomainResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDomainResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDomainResponseBody
	GetTotalCount() *int32
}

type ListDomainResponseBody struct {
	DomainInfoList []*ListDomainResponseBodyDomainInfoList `json:"DomainInfoList,omitempty" xml:"DomainInfoList,omitempty" type:"Repeated"`
	PageNumber     *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainResponseBody) GetDomainInfoList() []*ListDomainResponseBodyDomainInfoList {
	return s.DomainInfoList
}

func (s *ListDomainResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDomainResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDomainResponseBody) SetDomainInfoList(v []*ListDomainResponseBodyDomainInfoList) *ListDomainResponseBody {
	s.DomainInfoList = v
	return s
}

func (s *ListDomainResponseBody) SetPageNumber(v int32) *ListDomainResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDomainResponseBody) SetPageSize(v int32) *ListDomainResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDomainResponseBody) SetRequestId(v string) *ListDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainResponseBody) SetTotalCount(v int32) *ListDomainResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDomainResponseBody) Validate() error {
	if s.DomainInfoList != nil {
		for _, item := range s.DomainInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainResponseBodyDomainInfoList struct {
	DomainCode *string                                         `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	DomainName *string                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ToolCount  *int32                                          `json:"ToolCount,omitempty" xml:"ToolCount,omitempty"`
	ToolList   []*ListDomainResponseBodyDomainInfoListToolList `json:"ToolList,omitempty" xml:"ToolList,omitempty" type:"Repeated"`
}

func (s ListDomainResponseBodyDomainInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListDomainResponseBodyDomainInfoList) GoString() string {
	return s.String()
}

func (s *ListDomainResponseBodyDomainInfoList) GetDomainCode() *string {
	return s.DomainCode
}

func (s *ListDomainResponseBodyDomainInfoList) GetDomainName() *string {
	return s.DomainName
}

func (s *ListDomainResponseBodyDomainInfoList) GetToolCount() *int32 {
	return s.ToolCount
}

func (s *ListDomainResponseBodyDomainInfoList) GetToolList() []*ListDomainResponseBodyDomainInfoListToolList {
	return s.ToolList
}

func (s *ListDomainResponseBodyDomainInfoList) SetDomainCode(v string) *ListDomainResponseBodyDomainInfoList {
	s.DomainCode = &v
	return s
}

func (s *ListDomainResponseBodyDomainInfoList) SetDomainName(v string) *ListDomainResponseBodyDomainInfoList {
	s.DomainName = &v
	return s
}

func (s *ListDomainResponseBodyDomainInfoList) SetToolCount(v int32) *ListDomainResponseBodyDomainInfoList {
	s.ToolCount = &v
	return s
}

func (s *ListDomainResponseBodyDomainInfoList) SetToolList(v []*ListDomainResponseBodyDomainInfoListToolList) *ListDomainResponseBodyDomainInfoList {
	s.ToolList = v
	return s
}

func (s *ListDomainResponseBodyDomainInfoList) Validate() error {
	if s.ToolList != nil {
		for _, item := range s.ToolList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainResponseBodyDomainInfoListToolList struct {
	ToolCode *string `json:"ToolCode,omitempty" xml:"ToolCode,omitempty"`
	ToolName *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
}

func (s ListDomainResponseBodyDomainInfoListToolList) String() string {
	return dara.Prettify(s)
}

func (s ListDomainResponseBodyDomainInfoListToolList) GoString() string {
	return s.String()
}

func (s *ListDomainResponseBodyDomainInfoListToolList) GetToolCode() *string {
	return s.ToolCode
}

func (s *ListDomainResponseBodyDomainInfoListToolList) GetToolName() *string {
	return s.ToolName
}

func (s *ListDomainResponseBodyDomainInfoListToolList) SetToolCode(v string) *ListDomainResponseBodyDomainInfoListToolList {
	s.ToolCode = &v
	return s
}

func (s *ListDomainResponseBodyDomainInfoListToolList) SetToolName(v string) *ListDomainResponseBodyDomainInfoListToolList {
	s.ToolName = &v
	return s
}

func (s *ListDomainResponseBodyDomainInfoListToolList) Validate() error {
	return dara.Validate(s)
}
