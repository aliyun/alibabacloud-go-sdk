// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsFreezeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateOssCheckResultsFreezeResponseBodyData) *UpdateOssCheckResultsFreezeResponseBody
	GetData() *UpdateOssCheckResultsFreezeResponseBodyData
	SetRequestId(v string) *UpdateOssCheckResultsFreezeResponseBody
	GetRequestId() *string
}

type UpdateOssCheckResultsFreezeResponseBody struct {
	Data *UpdateOssCheckResultsFreezeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOssCheckResultsFreezeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsFreezeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsFreezeResponseBody) GetData() *UpdateOssCheckResultsFreezeResponseBodyData {
	return s.Data
}

func (s *UpdateOssCheckResultsFreezeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOssCheckResultsFreezeResponseBody) SetData(v *UpdateOssCheckResultsFreezeResponseBodyData) *UpdateOssCheckResultsFreezeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponseBody) SetRequestId(v string) *UpdateOssCheckResultsFreezeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateOssCheckResultsFreezeResponseBodyData struct {
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
	// 5
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s UpdateOssCheckResultsFreezeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsFreezeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) SetInvalidCount(v int32) *UpdateOssCheckResultsFreezeResponseBodyData {
	s.InvalidCount = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) SetRepeatCount(v int32) *UpdateOssCheckResultsFreezeResponseBodyData {
	s.RepeatCount = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) SetSuccessCount(v int32) *UpdateOssCheckResultsFreezeResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) SetTotalCount(v int32) *UpdateOssCheckResultsFreezeResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
