// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryFeaturesRequest
	GetInstanceId() *string
	SetResourceName(v string) *QueryFeaturesRequest
	GetResourceName() *string
	SetTargetScope(v string) *QueryFeaturesRequest
	GetTargetScope() *string
}

type QueryFeaturesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// at-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// worker-a
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s QueryFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFeaturesRequest) GoString() string {
	return s.String()
}

func (s *QueryFeaturesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryFeaturesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *QueryFeaturesRequest) GetTargetScope() *string {
	return s.TargetScope
}

func (s *QueryFeaturesRequest) SetInstanceId(v string) *QueryFeaturesRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryFeaturesRequest) SetResourceName(v string) *QueryFeaturesRequest {
	s.ResourceName = &v
	return s
}

func (s *QueryFeaturesRequest) SetTargetScope(v string) *QueryFeaturesRequest {
	s.TargetScope = &v
	return s
}

func (s *QueryFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
