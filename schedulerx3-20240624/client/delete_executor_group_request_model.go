// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExecutorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteExecutorGroupRequest
	GetClusterId() *string
	SetId(v int32) *DeleteExecutorGroupRequest
	GetId() *int32
}

type DeleteExecutorGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 83163
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteExecutorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExecutorGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteExecutorGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteExecutorGroupRequest) GetId() *int32 {
	return s.Id
}

func (s *DeleteExecutorGroupRequest) SetClusterId(v string) *DeleteExecutorGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteExecutorGroupRequest) SetId(v int32) *DeleteExecutorGroupRequest {
	s.Id = &v
	return s
}

func (s *DeleteExecutorGroupRequest) Validate() error {
	return dara.Validate(s)
}
