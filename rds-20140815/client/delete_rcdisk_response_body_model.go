// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCDiskResponseBody
	GetRequestId() *string
}

type DeleteRCDiskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8BE834C8-3C25-5AF8-BE3E-C8A690602A7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRCDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCDiskResponseBody) SetRequestId(v string) *DeleteRCDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
