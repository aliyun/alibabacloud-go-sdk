// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySensitiveSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySensitiveSwitchResponseBody
	GetRequestId() *string
}

type ModifySensitiveSwitchResponseBody struct {
	// example:
	//
	// 6169C0A4-B91A-5D48-AE4D-B9432D15****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySensitiveSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySensitiveSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySensitiveSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySensitiveSwitchResponseBody) SetRequestId(v string) *ModifySensitiveSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySensitiveSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
