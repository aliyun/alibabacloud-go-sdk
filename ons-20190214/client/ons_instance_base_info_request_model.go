// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsInstanceBaseInfoRequest
	GetInstanceId() *string
}

type OnsInstanceBaseInfoRequest struct {
	// The ID of the instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_138015630679****_BAAy1Hac
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsInstanceBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsInstanceBaseInfoRequest) SetInstanceId(v string) *OnsInstanceBaseInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
