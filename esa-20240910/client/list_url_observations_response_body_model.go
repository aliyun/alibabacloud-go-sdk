// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUrlObservationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListUrlObservationsResponseBodyConfigs) *ListUrlObservationsResponseBody
	GetConfigs() []*ListUrlObservationsResponseBodyConfigs
	SetPageNumber(v int32) *ListUrlObservationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUrlObservationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUrlObservationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUrlObservationsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListUrlObservationsResponseBody
	GetTotalPage() *int32
}

type ListUrlObservationsResponseBody struct {
	Configs    []*ListUrlObservationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListUrlObservationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUrlObservationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUrlObservationsResponseBody) GetConfigs() []*ListUrlObservationsResponseBodyConfigs {
	return s.Configs
}

func (s *ListUrlObservationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUrlObservationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUrlObservationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUrlObservationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUrlObservationsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListUrlObservationsResponseBody) SetConfigs(v []*ListUrlObservationsResponseBodyConfigs) *ListUrlObservationsResponseBody {
	s.Configs = v
	return s
}

func (s *ListUrlObservationsResponseBody) SetPageNumber(v int32) *ListUrlObservationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUrlObservationsResponseBody) SetPageSize(v int32) *ListUrlObservationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUrlObservationsResponseBody) SetRequestId(v string) *ListUrlObservationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUrlObservationsResponseBody) SetTotalCount(v int32) *ListUrlObservationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUrlObservationsResponseBody) SetTotalPage(v int32) *ListUrlObservationsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListUrlObservationsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUrlObservationsResponseBodyConfigs struct {
	ConfigId *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	SdkType  *string `json:"SdkType,omitempty" xml:"SdkType,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListUrlObservationsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListUrlObservationsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListUrlObservationsResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListUrlObservationsResponseBodyConfigs) GetSdkType() *string {
	return s.SdkType
}

func (s *ListUrlObservationsResponseBodyConfigs) GetUrl() *string {
	return s.Url
}

func (s *ListUrlObservationsResponseBodyConfigs) SetConfigId(v int64) *ListUrlObservationsResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListUrlObservationsResponseBodyConfigs) SetSdkType(v string) *ListUrlObservationsResponseBodyConfigs {
	s.SdkType = &v
	return s
}

func (s *ListUrlObservationsResponseBodyConfigs) SetUrl(v string) *ListUrlObservationsResponseBodyConfigs {
	s.Url = &v
	return s
}

func (s *ListUrlObservationsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
