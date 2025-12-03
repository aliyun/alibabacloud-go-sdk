// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteInstanceRequest
	GetClusterId() *string
	SetImmediateDeleteFlag(v bool) *DeleteInstanceRequest
	GetImmediateDeleteFlag() *bool
}

type DeleteInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	ImmediateDeleteFlag *bool `json:"ImmediateDeleteFlag,omitempty" xml:"ImmediateDeleteFlag,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteInstanceRequest) GetImmediateDeleteFlag() *bool {
	return s.ImmediateDeleteFlag
}

func (s *DeleteInstanceRequest) SetClusterId(v string) *DeleteInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteInstanceRequest) SetImmediateDeleteFlag(v bool) *DeleteInstanceRequest {
	s.ImmediateDeleteFlag = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
