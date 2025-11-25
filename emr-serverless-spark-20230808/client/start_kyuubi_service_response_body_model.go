// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartKyuubiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *StartKyuubiServiceResponseBody
	GetData() interface{}
	SetRequestId(v string) *StartKyuubiServiceResponseBody
	GetRequestId() *string
}

type StartKyuubiServiceResponseBody struct {
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartKyuubiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartKyuubiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartKyuubiServiceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *StartKyuubiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartKyuubiServiceResponseBody) SetData(v interface{}) *StartKyuubiServiceResponseBody {
	s.Data = v
	return s
}

func (s *StartKyuubiServiceResponseBody) SetRequestId(v string) *StartKyuubiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartKyuubiServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
