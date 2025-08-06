// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMeasureReadyFlagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *PutMeasureReadyFlagResponseBody
	GetData() *bool
	SetRequestId(v string) *PutMeasureReadyFlagResponseBody
	GetRequestId() *string
}

type PutMeasureReadyFlagResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutMeasureReadyFlagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutMeasureReadyFlagResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagResponseBody) GetData() *bool {
	return s.Data
}

func (s *PutMeasureReadyFlagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutMeasureReadyFlagResponseBody) SetData(v bool) *PutMeasureReadyFlagResponseBody {
	s.Data = &v
	return s
}

func (s *PutMeasureReadyFlagResponseBody) SetRequestId(v string) *PutMeasureReadyFlagResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMeasureReadyFlagResponseBody) Validate() error {
	return dara.Validate(s)
}
