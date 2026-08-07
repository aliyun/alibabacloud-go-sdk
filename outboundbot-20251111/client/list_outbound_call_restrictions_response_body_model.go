// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundCallRestrictionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOutboundCallRestrictionsResponseBody
	GetCode() *string
	SetData(v *ListOutboundCallRestrictionsResponseBodyData) *ListOutboundCallRestrictionsResponseBody
	GetData() *ListOutboundCallRestrictionsResponseBodyData
	SetHttpStatusCode(v int32) *ListOutboundCallRestrictionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListOutboundCallRestrictionsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListOutboundCallRestrictionsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListOutboundCallRestrictionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOutboundCallRestrictionsResponseBody
	GetSuccess() *bool
}

type ListOutboundCallRestrictionsResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ListOutboundCallRestrictionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=392db13c-8901-4a25-b566-91d0d8114cec
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOutboundCallRestrictionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallRestrictionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundCallRestrictionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOutboundCallRestrictionsResponseBody) GetData() *ListOutboundCallRestrictionsResponseBodyData {
	return s.Data
}

func (s *ListOutboundCallRestrictionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOutboundCallRestrictionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOutboundCallRestrictionsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListOutboundCallRestrictionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOutboundCallRestrictionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOutboundCallRestrictionsResponseBody) SetCode(v string) *ListOutboundCallRestrictionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBody) SetData(v *ListOutboundCallRestrictionsResponseBodyData) *ListOutboundCallRestrictionsResponseBody {
	s.Data = v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBody) SetHttpStatusCode(v int32) *ListOutboundCallRestrictionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBody) SetMessage(v string) *ListOutboundCallRestrictionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBody) SetParams(v []*string) *ListOutboundCallRestrictionsResponseBody {
	s.Params = v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBody) SetRequestId(v string) *ListOutboundCallRestrictionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBody) SetSuccess(v bool) *ListOutboundCallRestrictionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOutboundCallRestrictionsResponseBodyData struct {
	// The data list.
	OutboundCallRestrictions []*ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions `json:"OutboundCallRestrictions,omitempty" xml:"OutboundCallRestrictions,omitempty" type:"Repeated"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the conditions.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOutboundCallRestrictionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallRestrictionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOutboundCallRestrictionsResponseBodyData) GetOutboundCallRestrictions() []*ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions {
	return s.OutboundCallRestrictions
}

func (s *ListOutboundCallRestrictionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOutboundCallRestrictionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOutboundCallRestrictionsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOutboundCallRestrictionsResponseBodyData) SetOutboundCallRestrictions(v []*ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) *ListOutboundCallRestrictionsResponseBodyData {
	s.OutboundCallRestrictions = v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyData) SetPageNumber(v int32) *ListOutboundCallRestrictionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyData) SetPageSize(v int32) *ListOutboundCallRestrictionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyData) SetTotalCount(v int32) *ListOutboundCallRestrictionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyData) Validate() error {
	if s.OutboundCallRestrictions != nil {
		for _, item := range s.OutboundCallRestrictions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions struct {
	// The creation time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1735660800000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 示例值
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 示例值
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The policy. Valid values:
	//
	// 0: blacklist.
	//
	// 1: whitelist.
	//
	// example:
	//
	// 0
	Policy *int32 `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The remark.
	//
	// example:
	//
	// 示例值
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The outbound call restriction ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RestrictionId *string `json:"RestrictionId,omitempty" xml:"RestrictionId,omitempty"`
}

func (s ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) GoString() string {
	return s.String()
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) GetCreator() *string {
	return s.Creator
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) GetNumber() *string {
	return s.Number
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) GetPolicy() *int32 {
	return s.Policy
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) GetRemark() *string {
	return s.Remark
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) GetRestrictionId() *string {
	return s.RestrictionId
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) SetCreatedTime(v int64) *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions {
	s.CreatedTime = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) SetCreator(v string) *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions {
	s.Creator = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) SetNumber(v string) *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions {
	s.Number = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) SetPolicy(v int32) *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions {
	s.Policy = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) SetRemark(v string) *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions {
	s.Remark = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) SetRestrictionId(v string) *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions {
	s.RestrictionId = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponseBodyDataOutboundCallRestrictions) Validate() error {
	return dara.Validate(s)
}
