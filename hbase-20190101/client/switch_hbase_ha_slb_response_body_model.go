// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchHbaseHaSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchHbaseHaSlbResponseBody
	GetRequestId() *string
}

type SwitchHbaseHaSlbResponseBody struct {
	// example:
	//
	// C9D568D9-A59C-4AF2-8FBB-F086A841D58E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchHbaseHaSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchHbaseHaSlbResponseBody) SetRequestId(v string) *SwitchHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchHbaseHaSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
