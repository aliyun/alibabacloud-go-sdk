// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterNodePoolResponseBody
	GetRequestId() *string
}

type ModifyClusterNodePoolResponseBody struct {
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterNodePoolResponseBody) SetRequestId(v string) *ModifyClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
