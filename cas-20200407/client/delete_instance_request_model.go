// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteInstanceRequest
	GetInstanceId() *string
}

type DeleteInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas-cn-68n1mm16****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequest) SetClientToken(v string) *DeleteInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
