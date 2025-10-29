// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAICLoginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ManageAICLoginResponseBody
	GetRequestId() *string
}

type ManageAICLoginResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManageAICLoginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManageAICLoginResponseBody) GoString() string {
	return s.String()
}

func (s *ManageAICLoginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManageAICLoginResponseBody) SetRequestId(v string) *ManageAICLoginResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManageAICLoginResponseBody) Validate() error {
	return dara.Validate(s)
}
