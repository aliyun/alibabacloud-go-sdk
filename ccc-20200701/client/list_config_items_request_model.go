// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListConfigItemsRequest
	GetInstanceId() *string
	SetObjectId(v string) *ListConfigItemsRequest
	GetObjectId() *string
	SetObjectType(v string) *ListConfigItemsRequest
	GetObjectType() *string
}

type ListConfigItemsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s ListConfigItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigItemsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigItemsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConfigItemsRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListConfigItemsRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListConfigItemsRequest) SetInstanceId(v string) *ListConfigItemsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConfigItemsRequest) SetObjectId(v string) *ListConfigItemsRequest {
	s.ObjectId = &v
	return s
}

func (s *ListConfigItemsRequest) SetObjectType(v string) *ListConfigItemsRequest {
	s.ObjectType = &v
	return s
}

func (s *ListConfigItemsRequest) Validate() error {
	return dara.Validate(s)
}
