// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopKyuubiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *StopKyuubiServiceResponseBody
	GetData() interface{}
	SetRequestId(v string) *StopKyuubiServiceResponseBody
	GetRequestId() *string
}

type StopKyuubiServiceResponseBody struct {
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopKyuubiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopKyuubiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopKyuubiServiceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *StopKyuubiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopKyuubiServiceResponseBody) SetData(v interface{}) *StopKyuubiServiceResponseBody {
	s.Data = v
	return s
}

func (s *StopKyuubiServiceResponseBody) SetRequestId(v string) *StopKyuubiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopKyuubiServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
