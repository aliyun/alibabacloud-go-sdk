// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGranularityBySummaryIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGranularityBySummaryIdResponseBody
	GetCode() *string
	SetData(v *GetGranularityBySummaryIdResponseBodyData) *GetGranularityBySummaryIdResponseBody
	GetData() *GetGranularityBySummaryIdResponseBodyData
	SetHttpStatusCode(v int32) *GetGranularityBySummaryIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGranularityBySummaryIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGranularityBySummaryIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGranularityBySummaryIdResponseBody
	GetSuccess() *bool
}

type GetGranularityBySummaryIdResponseBody struct {
	// The error code. OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data
	Data *GetGranularityBySummaryIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGranularityBySummaryIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGranularityBySummaryIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetGranularityBySummaryIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGranularityBySummaryIdResponseBody) GetData() *GetGranularityBySummaryIdResponseBodyData {
	return s.Data
}

func (s *GetGranularityBySummaryIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGranularityBySummaryIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGranularityBySummaryIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGranularityBySummaryIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGranularityBySummaryIdResponseBody) SetCode(v string) *GetGranularityBySummaryIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetGranularityBySummaryIdResponseBody) SetData(v *GetGranularityBySummaryIdResponseBodyData) *GetGranularityBySummaryIdResponseBody {
	s.Data = v
	return s
}

func (s *GetGranularityBySummaryIdResponseBody) SetHttpStatusCode(v int32) *GetGranularityBySummaryIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGranularityBySummaryIdResponseBody) SetMessage(v string) *GetGranularityBySummaryIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetGranularityBySummaryIdResponseBody) SetRequestId(v string) *GetGranularityBySummaryIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGranularityBySummaryIdResponseBody) SetSuccess(v bool) *GetGranularityBySummaryIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetGranularityBySummaryIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGranularityBySummaryIdResponseBodyData struct {
	// DimIds
	DimIds []*int64 `json:"DimIds,omitempty" xml:"DimIds,omitempty" type:"Repeated"`
	// GranularityId
	//
	// example:
	//
	// 122
	GranularityId *int64 `json:"GranularityId,omitempty" xml:"GranularityId,omitempty"`
	// SummaryId
	//
	// example:
	//
	// 123
	SummaryId *int64 `json:"SummaryId,omitempty" xml:"SummaryId,omitempty"`
}

func (s GetGranularityBySummaryIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGranularityBySummaryIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGranularityBySummaryIdResponseBodyData) GetDimIds() []*int64 {
	return s.DimIds
}

func (s *GetGranularityBySummaryIdResponseBodyData) GetGranularityId() *int64 {
	return s.GranularityId
}

func (s *GetGranularityBySummaryIdResponseBodyData) GetSummaryId() *int64 {
	return s.SummaryId
}

func (s *GetGranularityBySummaryIdResponseBodyData) SetDimIds(v []*int64) *GetGranularityBySummaryIdResponseBodyData {
	s.DimIds = v
	return s
}

func (s *GetGranularityBySummaryIdResponseBodyData) SetGranularityId(v int64) *GetGranularityBySummaryIdResponseBodyData {
	s.GranularityId = &v
	return s
}

func (s *GetGranularityBySummaryIdResponseBodyData) SetSummaryId(v int64) *GetGranularityBySummaryIdResponseBodyData {
	s.SummaryId = &v
	return s
}

func (s *GetGranularityBySummaryIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
