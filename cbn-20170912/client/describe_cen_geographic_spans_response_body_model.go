// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenGeographicSpansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGeographicSpanModels(v *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) *DescribeCenGeographicSpansResponseBody
	GetGeographicSpanModels() *DescribeCenGeographicSpansResponseBodyGeographicSpanModels
	SetPageNumber(v int32) *DescribeCenGeographicSpansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenGeographicSpansResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenGeographicSpansResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenGeographicSpansResponseBody
	GetTotalCount() *int32
}

type DescribeCenGeographicSpansResponseBody struct {
	// A list of areas.
	GeographicSpanModels *DescribeCenGeographicSpansResponseBodyGeographicSpanModels `json:"GeographicSpanModels,omitempty" xml:"GeographicSpanModels,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 09780287-BC24-4164-8334-773432E32696
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenGeographicSpansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBody) GetGeographicSpanModels() *DescribeCenGeographicSpansResponseBodyGeographicSpanModels {
	return s.GeographicSpanModels
}

func (s *DescribeCenGeographicSpansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenGeographicSpansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenGeographicSpansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenGeographicSpansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenGeographicSpansResponseBody) SetGeographicSpanModels(v *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) *DescribeCenGeographicSpansResponseBody {
	s.GeographicSpanModels = v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetPageNumber(v int32) *DescribeCenGeographicSpansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetPageSize(v int32) *DescribeCenGeographicSpansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetRequestId(v string) *DescribeCenGeographicSpansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) SetTotalCount(v int32) *DescribeCenGeographicSpansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBody) Validate() error {
	if s.GeographicSpanModels != nil {
		if err := s.GeographicSpanModels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenGeographicSpansResponseBodyGeographicSpanModels struct {
	GeographicSpanModel []*DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel `json:"GeographicSpanModel,omitempty" xml:"GeographicSpanModel,omitempty" type:"Repeated"`
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModels) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) GetGeographicSpanModel() []*DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	return s.GeographicSpanModel
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) SetGeographicSpanModel(v []*DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) *DescribeCenGeographicSpansResponseBodyGeographicSpanModels {
	s.GeographicSpanModel = v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModels) Validate() error {
	if s.GeographicSpanModel != nil {
		for _, item := range s.GeographicSpanModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel struct {
	// The ID of the pair of connected areas.
	//
	// example:
	//
	// china_asia-pacific
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	// The ID of the local area.
	//
	// example:
	//
	// asia-pacific
	LocalGeoRegionId *string `json:"LocalGeoRegionId,omitempty" xml:"LocalGeoRegionId,omitempty"`
	// The ID of the peer area.
	//
	// example:
	//
	// china
	OppositeGeoRegionId *string `json:"OppositeGeoRegionId,omitempty" xml:"OppositeGeoRegionId,omitempty"`
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) GetGeographicSpanId() *string {
	return s.GeographicSpanId
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) GetLocalGeoRegionId() *string {
	return s.LocalGeoRegionId
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) GetOppositeGeoRegionId() *string {
	return s.OppositeGeoRegionId
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetGeographicSpanId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetLocalGeoRegionId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.LocalGeoRegionId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) SetOppositeGeoRegionId(v string) *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel {
	s.OppositeGeoRegionId = &v
	return s
}

func (s *DescribeCenGeographicSpansResponseBodyGeographicSpanModelsGeographicSpanModel) Validate() error {
	return dara.Validate(s)
}
