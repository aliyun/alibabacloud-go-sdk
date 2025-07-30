// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarningStrategyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWarningStrategyConfigResponseBody
	GetCode() *string
	SetCount(v int32) *ListWarningStrategyConfigResponseBody
	GetCount() *int32
	SetData(v *ListWarningStrategyConfigResponseBodyData) *ListWarningStrategyConfigResponseBody
	GetData() *ListWarningStrategyConfigResponseBodyData
	SetMessage(v string) *ListWarningStrategyConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWarningStrategyConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWarningStrategyConfigResponseBody
	GetSuccess() *bool
}

type ListWarningStrategyConfigResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      *ListWarningStrategyConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWarningStrategyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWarningStrategyConfigResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListWarningStrategyConfigResponseBody) GetData() *ListWarningStrategyConfigResponseBodyData {
	return s.Data
}

func (s *ListWarningStrategyConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWarningStrategyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWarningStrategyConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWarningStrategyConfigResponseBody) SetCode(v string) *ListWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetCount(v int32) *ListWarningStrategyConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetData(v *ListWarningStrategyConfigResponseBodyData) *ListWarningStrategyConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetMessage(v string) *ListWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetRequestId(v string) *ListWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetSuccess(v bool) *ListWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWarningStrategyConfigResponseBodyData struct {
	Data []*ListWarningStrategyConfigResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s ListWarningStrategyConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBodyData) GetData() []*ListWarningStrategyConfigResponseBodyDataData {
	return s.Data
}

func (s *ListWarningStrategyConfigResponseBodyData) SetData(v []*ListWarningStrategyConfigResponseBodyDataData) *ListWarningStrategyConfigResponseBodyData {
	s.Data = v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListWarningStrategyConfigResponseBodyDataData struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IntervalTime *int64  `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	Lambda       *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Level        *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	MaxNumber    *int64  `json:"MaxNumber,omitempty" xml:"MaxNumber,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListWarningStrategyConfigResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBodyDataData) GetId() *int64 {
	return s.Id
}

func (s *ListWarningStrategyConfigResponseBodyDataData) GetIntervalTime() *int64 {
	return s.IntervalTime
}

func (s *ListWarningStrategyConfigResponseBodyDataData) GetLambda() *string {
	return s.Lambda
}

func (s *ListWarningStrategyConfigResponseBodyDataData) GetLevel() *int64 {
	return s.Level
}

func (s *ListWarningStrategyConfigResponseBodyDataData) GetMaxNumber() *int64 {
	return s.MaxNumber
}

func (s *ListWarningStrategyConfigResponseBodyDataData) GetName() *string {
	return s.Name
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetId(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetIntervalTime(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.IntervalTime = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetLambda(v string) *ListWarningStrategyConfigResponseBodyDataData {
	s.Lambda = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetLevel(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.Level = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetMaxNumber(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.MaxNumber = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetName(v string) *ListWarningStrategyConfigResponseBodyDataData {
	s.Name = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
