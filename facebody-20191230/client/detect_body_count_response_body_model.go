// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectBodyCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectBodyCountResponseBodyData) *DetectBodyCountResponseBody
	GetData() *DetectBodyCountResponseBodyData
	SetRequestId(v string) *DetectBodyCountResponseBody
	GetRequestId() *string
}

type DetectBodyCountResponseBody struct {
	Data *DetectBodyCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1C709078-8886-4C91-AEDE-4E04EED0A689
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectBodyCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectBodyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponseBody) GetData() *DetectBodyCountResponseBodyData {
	return s.Data
}

func (s *DetectBodyCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectBodyCountResponseBody) SetData(v *DetectBodyCountResponseBodyData) *DetectBodyCountResponseBody {
	s.Data = v
	return s
}

func (s *DetectBodyCountResponseBody) SetRequestId(v string) *DetectBodyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectBodyCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectBodyCountResponseBodyData struct {
	// example:
	//
	// 6
	PersonNumber *int32 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
}

func (s DetectBodyCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectBodyCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponseBodyData) GetPersonNumber() *int32 {
	return s.PersonNumber
}

func (s *DetectBodyCountResponseBodyData) SetPersonNumber(v int32) *DetectBodyCountResponseBodyData {
	s.PersonNumber = &v
	return s
}

func (s *DetectBodyCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
