// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUpgradeToHasForHbaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUpgradeToHasForHbaseResponseBody
	GetRequestId() *string
}

type ModifyUpgradeToHasForHbaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUpgradeToHasForHbaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUpgradeToHasForHbaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUpgradeToHasForHbaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUpgradeToHasForHbaseResponseBody) SetRequestId(v string) *ModifyUpgradeToHasForHbaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseResponseBody) Validate() error {
	return dara.Validate(s)
}
