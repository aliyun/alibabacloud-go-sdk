// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetadataAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetMetadataAmountRequest
	GetInstanceId() *string
}

type GetMetadataAmountRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// amqp-cn-v0h1kb9n***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMetadataAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetadataAmountRequest) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMetadataAmountRequest) SetInstanceId(v string) *GetMetadataAmountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMetadataAmountRequest) Validate() error {
	return dara.Validate(s)
}
