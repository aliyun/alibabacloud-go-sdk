// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *GetResourcesRequest
	GetCluster() *string
	SetInstance(v string) *GetResourcesRequest
	GetInstance() *string
	SetType(v string) *GetResourcesRequest
	GetType() *string
}

type GetResourcesRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// mem
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetResourcesRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetResourcesRequest) GetType() *string {
	return s.Type
}

func (s *GetResourcesRequest) SetCluster(v string) *GetResourcesRequest {
	s.Cluster = &v
	return s
}

func (s *GetResourcesRequest) SetInstance(v string) *GetResourcesRequest {
	s.Instance = &v
	return s
}

func (s *GetResourcesRequest) SetType(v string) *GetResourcesRequest {
	s.Type = &v
	return s
}

func (s *GetResourcesRequest) Validate() error {
	return dara.Validate(s)
}
