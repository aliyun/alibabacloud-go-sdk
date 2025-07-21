// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAccountsRequest
	GetInstanceId() *string
}

type ListAccountsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance for which you want to query the static username and password.
	//
	// example:
	//
	// amqp-cn-20p****04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAccountsRequest) SetInstanceId(v string) *ListAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAccountsRequest) Validate() error {
	return dara.Validate(s)
}
