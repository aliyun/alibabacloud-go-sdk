// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKyuubiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *UpdateKyuubiServiceResponseBody
	GetData() interface{}
	SetRequestId(v string) *UpdateKyuubiServiceResponseBody
	GetRequestId() *string
}

type UpdateKyuubiServiceResponseBody struct {
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateKyuubiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKyuubiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKyuubiServiceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateKyuubiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKyuubiServiceResponseBody) SetData(v interface{}) *UpdateKyuubiServiceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKyuubiServiceResponseBody) SetRequestId(v string) *UpdateKyuubiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKyuubiServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
