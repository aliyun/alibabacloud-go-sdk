// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBNodeDescriptionResponseBody
	GetRequestId() *string
}

type ModifyDBNodeDescriptionResponseBody struct {
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodeDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodeDescriptionResponseBody) SetRequestId(v string) *ModifyDBNodeDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
