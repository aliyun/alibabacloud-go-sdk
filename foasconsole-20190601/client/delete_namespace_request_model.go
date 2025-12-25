// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteNamespaceRequest(v *DeleteNamespaceRequestDeleteNamespaceRequest) *DeleteNamespaceRequest
	GetDeleteNamespaceRequest() *DeleteNamespaceRequestDeleteNamespaceRequest
}

type DeleteNamespaceRequest struct {
	// This parameter is required.
	DeleteNamespaceRequest *DeleteNamespaceRequestDeleteNamespaceRequest `json:"DeleteNamespaceRequest,omitempty" xml:"DeleteNamespaceRequest,omitempty" type:"Struct"`
}

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetDeleteNamespaceRequest() *DeleteNamespaceRequestDeleteNamespaceRequest {
	return s.DeleteNamespaceRequest
}

func (s *DeleteNamespaceRequest) SetDeleteNamespaceRequest(v *DeleteNamespaceRequestDeleteNamespaceRequest) *DeleteNamespaceRequest {
	s.DeleteNamespaceRequest = v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	if s.DeleteNamespaceRequest != nil {
		if err := s.DeleteNamespaceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteNamespaceRequestDeleteNamespaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverlesspost_public_cn-0ju2bj2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ns-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteNamespaceRequestDeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequestDeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) SetInstanceId(v string) *DeleteNamespaceRequestDeleteNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequestDeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) SetRegion(v string) *DeleteNamespaceRequestDeleteNamespaceRequest {
	s.Region = &v
	return s
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
