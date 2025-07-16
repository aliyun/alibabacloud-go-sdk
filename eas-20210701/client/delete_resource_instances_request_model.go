// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllFailed(v bool) *DeleteResourceInstancesRequest
	GetAllFailed() *bool
	SetInstanceList(v string) *DeleteResourceInstancesRequest
	GetInstanceList() *string
}

type DeleteResourceInstancesRequest struct {
	// Specifies whether to delete all the instances that fail to be created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllFailed *bool `json:"AllFailed,omitempty" xml:"AllFailed,omitempty"`
	// The instances. Separate multiple instances with commas (,), such as `instanceId1,instanceId2`. For more information about how to query the instances, see [ListResourceInstances](https://help.aliyun.com/document_detail/412129.html).
	//
	// example:
	//
	// eas-i-xxxxxxx,eas-i-xxxxxxx
	InstanceList *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
}

func (s DeleteResourceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesRequest) GetAllFailed() *bool {
	return s.AllFailed
}

func (s *DeleteResourceInstancesRequest) GetInstanceList() *string {
	return s.InstanceList
}

func (s *DeleteResourceInstancesRequest) SetAllFailed(v bool) *DeleteResourceInstancesRequest {
	s.AllFailed = &v
	return s
}

func (s *DeleteResourceInstancesRequest) SetInstanceList(v string) *DeleteResourceInstancesRequest {
	s.InstanceList = &v
	return s
}

func (s *DeleteResourceInstancesRequest) Validate() error {
	return dara.Validate(s)
}
