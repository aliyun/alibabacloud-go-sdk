// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStandardStatisticsResponseBody
	GetCode() *string
	SetData(v *GetStandardStatisticsResponseBodyData) *GetStandardStatisticsResponseBody
	GetData() *GetStandardStatisticsResponseBodyData
	SetHttpStatusCode(v int32) *GetStandardStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStandardStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStandardStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStandardStatisticsResponseBody
	GetSuccess() *bool
}

type GetStandardStatisticsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetStandardStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStandardStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStandardStatisticsResponseBody) GetData() *GetStandardStatisticsResponseBodyData {
	return s.Data
}

func (s *GetStandardStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStandardStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStandardStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandardStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandardStatisticsResponseBody) SetCode(v string) *GetStandardStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetStandardStatisticsResponseBody) SetData(v *GetStandardStatisticsResponseBodyData) *GetStandardStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetStandardStatisticsResponseBody) SetHttpStatusCode(v int32) *GetStandardStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStandardStatisticsResponseBody) SetMessage(v string) *GetStandardStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetStandardStatisticsResponseBody) SetRequestId(v string) *GetStandardStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardStatisticsResponseBody) SetSuccess(v bool) *GetStandardStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandardStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardStatisticsResponseBodyData struct {
	StandardTypeCountList []*GetStandardStatisticsResponseBodyDataStandardTypeCountList `json:"StandardTypeCountList,omitempty" xml:"StandardTypeCountList,omitempty" type:"Repeated"`
	// example:
	//
	// 101
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetStandardStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsResponseBodyData) GetStandardTypeCountList() []*GetStandardStatisticsResponseBodyDataStandardTypeCountList {
	return s.StandardTypeCountList
}

func (s *GetStandardStatisticsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetStandardStatisticsResponseBodyData) SetStandardTypeCountList(v []*GetStandardStatisticsResponseBodyDataStandardTypeCountList) *GetStandardStatisticsResponseBodyData {
	s.StandardTypeCountList = v
	return s
}

func (s *GetStandardStatisticsResponseBodyData) SetTotalCount(v int32) *GetStandardStatisticsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetStandardStatisticsResponseBodyData) Validate() error {
	if s.StandardTypeCountList != nil {
		for _, item := range s.StandardTypeCountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStandardStatisticsResponseBodyDataStandardTypeCountList struct {
	// example:
	//
	// 11
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// test
	StandardType *string `json:"StandardType,omitempty" xml:"StandardType,omitempty"`
}

func (s GetStandardStatisticsResponseBodyDataStandardTypeCountList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsResponseBodyDataStandardTypeCountList) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsResponseBodyDataStandardTypeCountList) GetCount() *int32 {
	return s.Count
}

func (s *GetStandardStatisticsResponseBodyDataStandardTypeCountList) GetStandardType() *string {
	return s.StandardType
}

func (s *GetStandardStatisticsResponseBodyDataStandardTypeCountList) SetCount(v int32) *GetStandardStatisticsResponseBodyDataStandardTypeCountList {
	s.Count = &v
	return s
}

func (s *GetStandardStatisticsResponseBodyDataStandardTypeCountList) SetStandardType(v string) *GetStandardStatisticsResponseBodyDataStandardTypeCountList {
	s.StandardType = &v
	return s
}

func (s *GetStandardStatisticsResponseBodyDataStandardTypeCountList) Validate() error {
	return dara.Validate(s)
}
