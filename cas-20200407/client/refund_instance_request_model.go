// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RefundInstanceRequest
	GetInstanceId() *string
}

type RefundInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cas-ivauto-hqito6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RefundInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundInstanceRequest) GoString() string {
	return s.String()
}

func (s *RefundInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefundInstanceRequest) SetInstanceId(v string) *RefundInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundInstanceRequest) Validate() error {
	return dara.Validate(s)
}
