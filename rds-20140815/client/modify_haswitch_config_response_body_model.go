// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHASwitchConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHASwitchConfigResponseBody
	GetRequestId() *string
}

type ModifyHASwitchConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B6AE1448-D846-4831-B1C7-CFF3E99D5470
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHASwitchConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHASwitchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHASwitchConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHASwitchConfigResponseBody) SetRequestId(v string) *ModifyHASwitchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHASwitchConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
