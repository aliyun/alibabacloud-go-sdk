// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoGroupingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableExistingResourcesTransfer(v bool) *UpdateAutoGroupingConfigRequest
	GetEnableExistingResourcesTransfer() *bool
}

type UpdateAutoGroupingConfigRequest struct {
	// Specifies whether to enable existing associated resources to follow the group transfer task. Valid values:
	//
	// - false: Existing associated resources do not follow the group transfer task.
	//
	// - true: Existing associated resources follow the group transfer task.
	//
	// example:
	//
	// false
	EnableExistingResourcesTransfer *bool `json:"EnableExistingResourcesTransfer,omitempty" xml:"EnableExistingResourcesTransfer,omitempty"`
}

func (s UpdateAutoGroupingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoGroupingConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoGroupingConfigRequest) GetEnableExistingResourcesTransfer() *bool {
	return s.EnableExistingResourcesTransfer
}

func (s *UpdateAutoGroupingConfigRequest) SetEnableExistingResourcesTransfer(v bool) *UpdateAutoGroupingConfigRequest {
	s.EnableExistingResourcesTransfer = &v
	return s
}

func (s *UpdateAutoGroupingConfigRequest) Validate() error {
	return dara.Validate(s)
}
