// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUIAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUIAccountPasswordResponseBody
	GetRequestId() *string
}

type ModifyUIAccountPasswordResponseBody struct {
	// example:
	//
	// BED4ADEB-4EA9-507E-892C-84112D6AC7C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUIAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUIAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUIAccountPasswordResponseBody) SetRequestId(v string) *ModifyUIAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUIAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
