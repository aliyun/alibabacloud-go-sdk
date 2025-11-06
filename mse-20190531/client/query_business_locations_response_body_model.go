// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBusinessLocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryBusinessLocationsResponseBodyData) *QueryBusinessLocationsResponseBody
	GetData() []*QueryBusinessLocationsResponseBodyData
	SetErrorCode(v string) *QueryBusinessLocationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryBusinessLocationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBusinessLocationsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *QueryBusinessLocationsResponseBody
	GetSuccess() *string
}

type QueryBusinessLocationsResponseBody struct {
	// The details of the data.
	Data []*QueryBusinessLocationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The parameter is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3369AD10-F1A6-4E6F-B99E-20F51826****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBusinessLocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBusinessLocationsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBusinessLocationsResponseBody) GetData() []*QueryBusinessLocationsResponseBodyData {
	return s.Data
}

func (s *QueryBusinessLocationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryBusinessLocationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBusinessLocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBusinessLocationsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QueryBusinessLocationsResponseBody) SetData(v []*QueryBusinessLocationsResponseBodyData) *QueryBusinessLocationsResponseBody {
	s.Data = v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetErrorCode(v string) *QueryBusinessLocationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetMessage(v string) *QueryBusinessLocationsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetRequestId(v string) *QueryBusinessLocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetSuccess(v string) *QueryBusinessLocationsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBusinessLocationsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryBusinessLocationsResponseBodyData struct {
	// The Chinese name of the region.
	//
	// example:
	//
	// China (Shanghai)
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// The description.
	//
	// example:
	//
	// China (Shanghai)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Chinese name of the district.
	//
	// example:
	//
	// Asia Pacific
	DistrictCnName *string `json:"DistrictCnName,omitempty" xml:"DistrictCnName,omitempty"`
	// The English name of the district.
	//
	// example:
	//
	// Asia Pacific
	DistrictEnName *string `json:"DistrictEnName,omitempty" xml:"DistrictEnName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// asia-pacific
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// The ordering information of the district.
	//
	// example:
	//
	// 101
	DistrictOrdering *int32 `json:"DistrictOrdering,omitempty" xml:"DistrictOrdering,omitempty"`
	// The display name of the district.
	//
	// example:
	//
	// Asia Pacific
	DistrictShowName *string `json:"DistrictShowName,omitempty" xml:"DistrictShowName,omitempty"`
	// The complete description of the region.
	//
	// example:
	//
	// China East 2 (Shanghai)
	EnDescription *string `json:"EnDescription,omitempty" xml:"EnDescription,omitempty"`
	// The English name of the region.
	//
	// example:
	//
	// China (Shanghai)
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// The name.
	//
	// example:
	//
	// cn-shanghai
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ordering information.
	//
	// example:
	//
	// 6
	Ordering *int32 `json:"Ordering,omitempty" xml:"Ordering,omitempty"`
	// The display name.
	//
	// example:
	//
	// China (Shanghai)
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The type.
	//
	// example:
	//
	// region
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryBusinessLocationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBusinessLocationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBusinessLocationsResponseBodyData) GetCnName() *string {
	return s.CnName
}

func (s *QueryBusinessLocationsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QueryBusinessLocationsResponseBodyData) GetDistrictCnName() *string {
	return s.DistrictCnName
}

func (s *QueryBusinessLocationsResponseBodyData) GetDistrictEnName() *string {
	return s.DistrictEnName
}

func (s *QueryBusinessLocationsResponseBodyData) GetDistrictId() *string {
	return s.DistrictId
}

func (s *QueryBusinessLocationsResponseBodyData) GetDistrictOrdering() *int32 {
	return s.DistrictOrdering
}

func (s *QueryBusinessLocationsResponseBodyData) GetDistrictShowName() *string {
	return s.DistrictShowName
}

func (s *QueryBusinessLocationsResponseBodyData) GetEnDescription() *string {
	return s.EnDescription
}

func (s *QueryBusinessLocationsResponseBodyData) GetEnName() *string {
	return s.EnName
}

func (s *QueryBusinessLocationsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryBusinessLocationsResponseBodyData) GetOrdering() *int32 {
	return s.Ordering
}

func (s *QueryBusinessLocationsResponseBodyData) GetShowName() *string {
	return s.ShowName
}

func (s *QueryBusinessLocationsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *QueryBusinessLocationsResponseBodyData) SetCnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.CnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDescription(v string) *QueryBusinessLocationsResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictCnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictCnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictEnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictEnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictId(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictId = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictOrdering(v int32) *QueryBusinessLocationsResponseBodyData {
	s.DistrictOrdering = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictShowName(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictShowName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetEnDescription(v string) *QueryBusinessLocationsResponseBodyData {
	s.EnDescription = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetEnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.EnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetName(v string) *QueryBusinessLocationsResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetOrdering(v int32) *QueryBusinessLocationsResponseBodyData {
	s.Ordering = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetShowName(v string) *QueryBusinessLocationsResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetType(v string) *QueryBusinessLocationsResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
