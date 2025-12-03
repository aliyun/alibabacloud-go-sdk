// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHbaseHaSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHbaseHaSlbResponseBody
	GetRequestId() *string
}

type DeleteHbaseHaSlbResponseBody struct {
	// example:
	//
	// C9D568D9-A59C-4AF2-8FBB-F086A841D58E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHbaseHaSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHbaseHaSlbResponseBody) SetRequestId(v string) *DeleteHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHbaseHaSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
