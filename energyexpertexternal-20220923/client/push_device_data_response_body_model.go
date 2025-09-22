// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDeviceDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *PushDeviceDataResponseBody
	GetData() *string
	SetRequestId(v string) *PushDeviceDataResponseBody
	GetRequestId() *string
}

type PushDeviceDataResponseBody struct {
	// Whether the data is pushed successfully. Success is returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PushDeviceDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushDeviceDataResponseBody) GetData() *string {
	return s.Data
}

func (s *PushDeviceDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushDeviceDataResponseBody) SetData(v string) *PushDeviceDataResponseBody {
	s.Data = &v
	return s
}

func (s *PushDeviceDataResponseBody) SetRequestId(v string) *PushDeviceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushDeviceDataResponseBody) Validate() error {
	return dara.Validate(s)
}
