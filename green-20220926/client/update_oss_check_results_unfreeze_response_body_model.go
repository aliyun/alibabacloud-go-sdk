// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsUnfreezeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateOssCheckResultsUnfreezeResponseBodyData) *UpdateOssCheckResultsUnfreezeResponseBody
	GetData() *UpdateOssCheckResultsUnfreezeResponseBodyData
	SetRequestId(v string) *UpdateOssCheckResultsUnfreezeResponseBody
	GetRequestId() *string
}

type UpdateOssCheckResultsUnfreezeResponseBody struct {
	Data *UpdateOssCheckResultsUnfreezeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOssCheckResultsUnfreezeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsUnfreezeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsUnfreezeResponseBody) GetData() *UpdateOssCheckResultsUnfreezeResponseBodyData {
	return s.Data
}

func (s *UpdateOssCheckResultsUnfreezeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOssCheckResultsUnfreezeResponseBody) SetData(v *UpdateOssCheckResultsUnfreezeResponseBodyData) *UpdateOssCheckResultsUnfreezeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponseBody) SetRequestId(v string) *UpdateOssCheckResultsUnfreezeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateOssCheckResultsUnfreezeResponseBodyData struct {
	// example:
	//
	// 1
	InvalidCount *int32 `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	// example:
	//
	// 1
	RepeatCount *int32 `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s UpdateOssCheckResultsUnfreezeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsUnfreezeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) SetInvalidCount(v int32) *UpdateOssCheckResultsUnfreezeResponseBodyData {
	s.InvalidCount = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) SetRepeatCount(v int32) *UpdateOssCheckResultsUnfreezeResponseBodyData {
	s.RepeatCount = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) SetSuccessCount(v int32) *UpdateOssCheckResultsUnfreezeResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) SetTotalCount(v int32) *UpdateOssCheckResultsUnfreezeResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
