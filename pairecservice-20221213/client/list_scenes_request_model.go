// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScenesRequest
	GetInstanceId() *string
	SetName(v string) *ListScenesRequest
	GetName() *string
}

type ListScenesRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// scene1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScenesRequest) GoString() string {
	return s.String()
}

func (s *ListScenesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScenesRequest) GetName() *string {
	return s.Name
}

func (s *ListScenesRequest) SetInstanceId(v string) *ListScenesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScenesRequest) SetName(v string) *ListScenesRequest {
	s.Name = &v
	return s
}

func (s *ListScenesRequest) Validate() error {
	return dara.Validate(s)
}
