// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAncillarySuggestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AncillarySuggestResponseBody
	GetRequestId() *string
	SetData(v *AncillarySuggestResponseBodyData) *AncillarySuggestResponseBody
	GetData() *AncillarySuggestResponseBodyData
	SetErrorCode(v string) *AncillarySuggestResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *AncillarySuggestResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *AncillarySuggestResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *AncillarySuggestResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *AncillarySuggestResponseBody
	GetSuccess() *bool
}

type AncillarySuggestResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Properly processed return data
	Data *AncillarySuggestResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// Data carried in error handling
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// Error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http request successful, status value is always 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AncillarySuggestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AncillarySuggestResponseBody) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AncillarySuggestResponseBody) GetData() *AncillarySuggestResponseBodyData {
	return s.Data
}

func (s *AncillarySuggestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AncillarySuggestResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *AncillarySuggestResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *AncillarySuggestResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *AncillarySuggestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AncillarySuggestResponseBody) SetRequestId(v string) *AncillarySuggestResponseBody {
	s.RequestId = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetData(v *AncillarySuggestResponseBodyData) *AncillarySuggestResponseBody {
	s.Data = v
	return s
}

func (s *AncillarySuggestResponseBody) SetErrorCode(v string) *AncillarySuggestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetErrorData(v interface{}) *AncillarySuggestResponseBody {
	s.ErrorData = v
	return s
}

func (s *AncillarySuggestResponseBody) SetErrorMsg(v string) *AncillarySuggestResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetStatus(v int32) *AncillarySuggestResponseBody {
	s.Status = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetSuccess(v bool) *AncillarySuggestResponseBody {
	s.Success = &v
	return s
}

func (s *AncillarySuggestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AncillarySuggestResponseBodyData struct {
	// ancillary detail list
	SegAncillaryMapList []*AncillarySuggestResponseBodyDataSegAncillaryMapList `json:"seg_ancillary_map_list,omitempty" xml:"seg_ancillary_map_list,omitempty" type:"Repeated"`
	// solution_id, equals to solution_id in request
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s AncillarySuggestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AncillarySuggestResponseBodyData) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyData) GetSegAncillaryMapList() []*AncillarySuggestResponseBodyDataSegAncillaryMapList {
	return s.SegAncillaryMapList
}

func (s *AncillarySuggestResponseBodyData) GetSolutionId() *string {
	return s.SolutionId
}

func (s *AncillarySuggestResponseBodyData) SetSegAncillaryMapList(v []*AncillarySuggestResponseBodyDataSegAncillaryMapList) *AncillarySuggestResponseBodyData {
	s.SegAncillaryMapList = v
	return s
}

func (s *AncillarySuggestResponseBodyData) SetSolutionId(v string) *AncillarySuggestResponseBodyData {
	s.SolutionId = &v
	return s
}

func (s *AncillarySuggestResponseBodyData) Validate() error {
	if s.SegAncillaryMapList != nil {
		for _, item := range s.SegAncillaryMapList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AncillarySuggestResponseBodyDataSegAncillaryMapList struct {
	// Ancillary product
	Ancillary *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary `json:"ancillary,omitempty" xml:"ancillary,omitempty" type:"Struct"`
	// Segment ID list, these segments share the same ancillary
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapList) String() string {
	return dara.Prettify(s)
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapList) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapList) GetAncillary() *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary {
	return s.Ancillary
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapList) SetAncillary(v *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) *AncillarySuggestResponseBodyDataSegAncillaryMapList {
	s.Ancillary = v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapList) SetSegmentIdList(v []*string) *AncillarySuggestResponseBodyDataSegAncillaryMapList {
	s.SegmentIdList = v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapList) Validate() error {
	if s.Ancillary != nil {
		if err := s.Ancillary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary struct {
	// Ancillary product ID
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU
	AncillaryId *string `json:"ancillary_id,omitempty" xml:"ancillary_id,omitempty"`
	// Ancillary product type. currently supports 4: paid luggage
	//
	// example:
	//
	// 4
	AncillaryType *int32 `json:"ancillary_type,omitempty" xml:"ancillary_type,omitempty"`
	// Baggage details
	BaggageAncillary *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary `json:"baggage_ancillary,omitempty" xml:"baggage_ancillary,omitempty" type:"Struct"`
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) String() string {
	return dara.Prettify(s)
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) GetAncillaryId() *string {
	return s.AncillaryId
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) GetAncillaryType() *int32 {
	return s.AncillaryType
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) GetBaggageAncillary() *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	return s.BaggageAncillary
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) SetAncillaryId(v string) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary {
	s.AncillaryId = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) SetAncillaryType(v int32) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary {
	s.AncillaryType = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) SetBaggageAncillary(v *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary {
	s.BaggageAncillary = v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) Validate() error {
	if s.BaggageAncillary != nil {
		if err := s.BaggageAncillary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary struct {
	// baggage quantity, values such as: 3, 2, 1, 0, -2.     -2 indicates weight-based
	//
	// example:
	//
	// 0
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// Baggage weight, 0-50. When isAllWeight=true, it represents the total weight of all baggages.
	//
	// example:
	//
	// 0
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// Unit of baggage weight
	//
	// example:
	//
	// KG
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllWeight *bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// Total price
	//
	// example:
	//
	// 10.0
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) String() string {
	return dara.Prettify(s)
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) GetBaggageAmount() *int32 {
	return s.BaggageAmount
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) GetBaggageWeight() *int32 {
	return s.BaggageWeight
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) GetBaggageWeightUnit() *string {
	return s.BaggageWeightUnit
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) GetIsAllWeight() *bool {
	return s.IsAllWeight
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) GetPrice() *float64 {
	return s.Price
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetBaggageAmount(v int32) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.BaggageAmount = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetBaggageWeight(v int32) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.BaggageWeight = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetBaggageWeightUnit(v string) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.BaggageWeightUnit = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetIsAllWeight(v bool) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.IsAllWeight = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetPrice(v float64) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.Price = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) Validate() error {
	return dara.Validate(s)
}
