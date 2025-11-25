// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKyuubiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteKyuubiServiceResponseBody
	GetData() interface{}
	SetRequestId(v string) *DeleteKyuubiServiceResponseBody
	GetRequestId() *string
}

type DeleteKyuubiServiceResponseBody struct {
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteKyuubiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKyuubiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKyuubiServiceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteKyuubiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKyuubiServiceResponseBody) SetData(v interface{}) *DeleteKyuubiServiceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteKyuubiServiceResponseBody) SetRequestId(v string) *DeleteKyuubiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKyuubiServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
