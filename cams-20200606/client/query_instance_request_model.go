// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryInstanceRequest
	GetInstanceId() *string
}

type QueryInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9293938****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstanceRequest) SetInstanceId(v string) *QueryInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceRequest) Validate() error {
	return dara.Validate(s)
}
