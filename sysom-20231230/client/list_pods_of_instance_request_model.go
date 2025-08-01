// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPodsOfInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListPodsOfInstanceRequest
	GetClusterId() *string
	SetCurrent(v int64) *ListPodsOfInstanceRequest
	GetCurrent() *int64
	SetInstance(v string) *ListPodsOfInstanceRequest
	GetInstance() *string
	SetPageSize(v int64) *ListPodsOfInstanceRequest
	GetPageSize() *int64
}

type ListPodsOfInstanceRequest struct {
	// example:
	//
	// c96e34d74eb6748f3b2a46552d5d653f6
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPodsOfInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPodsOfInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPodsOfInstanceRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListPodsOfInstanceRequest) GetInstance() *string {
	return s.Instance
}

func (s *ListPodsOfInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPodsOfInstanceRequest) SetClusterId(v string) *ListPodsOfInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetCurrent(v int64) *ListPodsOfInstanceRequest {
	s.Current = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetInstance(v string) *ListPodsOfInstanceRequest {
	s.Instance = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetPageSize(v int64) *ListPodsOfInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListPodsOfInstanceRequest) Validate() error {
	return dara.Validate(s)
}
