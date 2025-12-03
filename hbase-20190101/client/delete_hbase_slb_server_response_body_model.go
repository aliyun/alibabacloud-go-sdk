// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHBaseSlbServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHBaseSlbServerResponseBody
	GetRequestId() *string
}

type DeleteHBaseSlbServerResponseBody struct {
	// example:
	//
	// 7242130A-82CF-49BF-AB32-30DCB819EBA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHBaseSlbServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHBaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHBaseSlbServerResponseBody) SetRequestId(v string) *DeleteHBaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHBaseSlbServerResponseBody) Validate() error {
	return dara.Validate(s)
}
