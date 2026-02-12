// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsInstanceDeleteRequest
	GetInstanceId() *string
}

type OnsInstanceDeleteRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsInstanceDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceDeleteRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsInstanceDeleteRequest) SetInstanceId(v string) *OnsInstanceDeleteRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceDeleteRequest) Validate() error {
	return dara.Validate(s)
}
