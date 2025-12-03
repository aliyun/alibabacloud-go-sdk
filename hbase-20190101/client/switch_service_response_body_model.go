// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchServiceResponseBody
	GetRequestId() *string
}

type SwitchServiceResponseBody struct {
	// example:
	//
	// F1005DE4-D981-559F-9E37-5172DXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchServiceResponseBody) SetRequestId(v string) *SwitchServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
