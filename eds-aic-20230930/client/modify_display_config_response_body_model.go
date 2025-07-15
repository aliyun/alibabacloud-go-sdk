// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDisplayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDisplayConfigResponseBody
	GetRequestId() *string
}

type ModifyDisplayConfigResponseBody struct {
	// example:
	//
	// A578AD3A-8E7C-54FE-A09F-B060941*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDisplayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDisplayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDisplayConfigResponseBody) SetRequestId(v string) *ModifyDisplayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDisplayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
