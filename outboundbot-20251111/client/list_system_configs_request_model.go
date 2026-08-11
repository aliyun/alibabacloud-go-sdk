// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListSystemConfigsRequest
	GetName() *string
	SetObjectId(v string) *ListSystemConfigsRequest
	GetObjectId() *string
	SetObjectType(v string) *ListSystemConfigsRequest
	GetObjectType() *string
}

type ListSystemConfigsRequest struct {
	// The system configuration name.\\
	//
	// callableTime: the outbound job window.\\
	//
	// calleeDailyAttemptLimit: the maximum number of daily calls to a single callee number.
	//
	// example:
	//
	// callableTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration type ID.\\
	//
	// If ObjectType is set to INSTANCE, this parameter specifies the instance ID.\\
	//
	// If ObjectType is set to TENANT, this parameter specifies the tenant ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The configuration type.\\
	//
	// INSTANCE: instance-level.\\
	//
	// TENANT: tenant-level.
	//
	// example:
	//
	// INSTANCE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s ListSystemConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListSystemConfigsRequest) GetName() *string {
	return s.Name
}

func (s *ListSystemConfigsRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListSystemConfigsRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListSystemConfigsRequest) SetName(v string) *ListSystemConfigsRequest {
	s.Name = &v
	return s
}

func (s *ListSystemConfigsRequest) SetObjectId(v string) *ListSystemConfigsRequest {
	s.ObjectId = &v
	return s
}

func (s *ListSystemConfigsRequest) SetObjectType(v string) *ListSystemConfigsRequest {
	s.ObjectType = &v
	return s
}

func (s *ListSystemConfigsRequest) Validate() error {
	return dara.Validate(s)
}
