// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHaVipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHaVipsResponseBody
	GetRequestId() *string
}

type DeleteHaVipsResponseBody struct {
	// example:
	//
	// 6666C5A5-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHaVipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHaVipsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHaVipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHaVipsResponseBody) SetRequestId(v string) *DeleteHaVipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHaVipsResponseBody) Validate() error {
	return dara.Validate(s)
}
