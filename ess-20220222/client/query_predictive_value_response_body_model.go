// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPredictiveValues(v *QueryPredictiveValueResponseBodyPredictiveValues) *QueryPredictiveValueResponseBody
	GetPredictiveValues() *QueryPredictiveValueResponseBodyPredictiveValues
	SetRequestId(v string) *QueryPredictiveValueResponseBody
	GetRequestId() *string
}

type QueryPredictiveValueResponseBody struct {
	PredictiveValues *QueryPredictiveValueResponseBodyPredictiveValues `json:"PredictiveValues,omitempty" xml:"PredictiveValues,omitempty" type:"Struct"`
	// example:
	//
	// FA5F448E-DC84-5EBC-B6D5-F659ADxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPredictiveValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveValueResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPredictiveValueResponseBody) GetPredictiveValues() *QueryPredictiveValueResponseBodyPredictiveValues {
	return s.PredictiveValues
}

func (s *QueryPredictiveValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPredictiveValueResponseBody) SetPredictiveValues(v *QueryPredictiveValueResponseBodyPredictiveValues) *QueryPredictiveValueResponseBody {
	s.PredictiveValues = v
	return s
}

func (s *QueryPredictiveValueResponseBody) SetRequestId(v string) *QueryPredictiveValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPredictiveValueResponseBody) Validate() error {
	if s.PredictiveValues != nil {
		if err := s.PredictiveValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPredictiveValueResponseBodyPredictiveValues struct {
	PredictiveValue []*QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue `json:"PredictiveValue,omitempty" xml:"PredictiveValue,omitempty" type:"Repeated"`
}

func (s QueryPredictiveValueResponseBodyPredictiveValues) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveValueResponseBodyPredictiveValues) GoString() string {
	return s.String()
}

func (s *QueryPredictiveValueResponseBodyPredictiveValues) GetPredictiveValue() []*QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue {
	return s.PredictiveValue
}

func (s *QueryPredictiveValueResponseBodyPredictiveValues) SetPredictiveValue(v []*QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) *QueryPredictiveValueResponseBodyPredictiveValues {
	s.PredictiveValue = v
	return s
}

func (s *QueryPredictiveValueResponseBodyPredictiveValues) Validate() error {
	if s.PredictiveValue != nil {
		for _, item := range s.PredictiveValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue struct {
	// example:
	//
	// 2025-12-17T10:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 3
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) GoString() string {
	return s.String()
}

func (s *QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) GetTime() *string {
	return s.Time
}

func (s *QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) GetValue() *int32 {
	return s.Value
}

func (s *QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) SetTime(v string) *QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue {
	s.Time = &v
	return s
}

func (s *QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) SetValue(v int32) *QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue {
	s.Value = &v
	return s
}

func (s *QueryPredictiveValueResponseBodyPredictiveValuesPredictiveValue) Validate() error {
	return dara.Validate(s)
}
