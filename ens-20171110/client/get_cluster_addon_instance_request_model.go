// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAddonInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterAddonInstanceRequest
	GetClusterId() *string
	SetInstanceName(v string) *GetClusterAddonInstanceRequest
	GetInstanceName() *string
}

type GetClusterAddonInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// edge-csi-lite
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GetClusterAddonInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAddonInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetClusterAddonInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterAddonInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetClusterAddonInstanceRequest) SetClusterId(v string) *GetClusterAddonInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterAddonInstanceRequest) SetInstanceName(v string) *GetClusterAddonInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *GetClusterAddonInstanceRequest) Validate() error {
	return dara.Validate(s)
}
