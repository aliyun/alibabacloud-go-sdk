// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSqlOptimizeStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceSqlOptimizeStatisticResponseBody
	GetCode() *string
	SetData(v *GetInstanceSqlOptimizeStatisticResponseBodyData) *GetInstanceSqlOptimizeStatisticResponseBody
	GetData() *GetInstanceSqlOptimizeStatisticResponseBodyData
	SetMessage(v string) *GetInstanceSqlOptimizeStatisticResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceSqlOptimizeStatisticResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetInstanceSqlOptimizeStatisticResponseBody
	GetSuccess() *string
}

type GetInstanceSqlOptimizeStatisticResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the automatic SQL optimization events.
	Data *GetInstanceSqlOptimizeStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceSqlOptimizeStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSqlOptimizeStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) GetData() *GetInstanceSqlOptimizeStatisticResponseBodyData {
	return s.Data
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) SetCode(v string) *GetInstanceSqlOptimizeStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) SetData(v *GetInstanceSqlOptimizeStatisticResponseBodyData) *GetInstanceSqlOptimizeStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) SetMessage(v string) *GetInstanceSqlOptimizeStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) SetRequestId(v string) *GetInstanceSqlOptimizeStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) SetSuccess(v string) *GetInstanceSqlOptimizeStatisticResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceSqlOptimizeStatisticResponseBodyData struct {
	// The total number of automatic SQL optimization events.
	//
	// example:
	//
	// 16
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The multiple of the maximum improvement for returned automatic SQL optimization events.
	//
	// example:
	//
	// 1003
	Improvement *float64 `json:"improvement,omitempty" xml:"improvement,omitempty"`
}

func (s GetInstanceSqlOptimizeStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSqlOptimizeStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceSqlOptimizeStatisticResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetInstanceSqlOptimizeStatisticResponseBodyData) GetImprovement() *float64 {
	return s.Improvement
}

func (s *GetInstanceSqlOptimizeStatisticResponseBodyData) SetCount(v int32) *GetInstanceSqlOptimizeStatisticResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponseBodyData) SetImprovement(v float64) *GetInstanceSqlOptimizeStatisticResponseBodyData {
	s.Improvement = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
