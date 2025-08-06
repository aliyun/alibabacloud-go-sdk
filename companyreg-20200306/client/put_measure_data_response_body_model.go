// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMeasureDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *PutMeasureDataResponseBody
	GetData() *bool
	SetRequestId(v string) *PutMeasureDataResponseBody
	GetRequestId() *string
}

type PutMeasureDataResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 6A603AA0-73BA-52B3-AC7D-0F846ECF7A9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutMeasureDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutMeasureDataResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponseBody) GetData() *bool {
	return s.Data
}

func (s *PutMeasureDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutMeasureDataResponseBody) SetData(v bool) *PutMeasureDataResponseBody {
	s.Data = &v
	return s
}

func (s *PutMeasureDataResponseBody) SetRequestId(v string) *PutMeasureDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMeasureDataResponseBody) Validate() error {
	return dara.Validate(s)
}
