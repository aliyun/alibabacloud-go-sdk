// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpgradeMinorVersionRequest
	GetClusterId() *string
	SetComponents(v string) *UpgradeMinorVersionRequest
	GetComponents() *string
}

type UpgradeMinorVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// HADOOP
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeMinorVersionRequest) GetComponents() *string {
	return s.Components
}

func (s *UpgradeMinorVersionRequest) SetClusterId(v string) *UpgradeMinorVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetComponents(v string) *UpgradeMinorVersionRequest {
	s.Components = &v
	return s
}

func (s *UpgradeMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
