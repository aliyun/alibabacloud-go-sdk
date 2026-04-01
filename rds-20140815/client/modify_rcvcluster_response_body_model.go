// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCVClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCVClusterResponseBody
	GetRequestId() *string
}

type ModifyRCVClusterResponseBody struct {
	// example:
	//
	// 2553A660-E4EB-4AF4-A402-8AFF70A49143
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCVClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCVClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCVClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCVClusterResponseBody) SetRequestId(v string) *ModifyRCVClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCVClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
