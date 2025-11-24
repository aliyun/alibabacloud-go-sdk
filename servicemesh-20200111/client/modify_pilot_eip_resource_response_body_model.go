// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPilotEipResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPilotEipResourceResponseBody
	GetRequestId() *string
}

type ModifyPilotEipResourceResponseBody struct {
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPilotEipResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPilotEipResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPilotEipResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPilotEipResourceResponseBody) SetRequestId(v string) *ModifyPilotEipResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPilotEipResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
