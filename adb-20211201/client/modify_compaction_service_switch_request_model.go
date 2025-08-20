// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCompactionServiceSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyCompactionServiceSwitchRequest
	GetDBClusterId() *string
	SetEnableCompactionService(v bool) *ModifyCompactionServiceSwitchRequest
	GetEnableCompactionService() *bool
}

type ModifyCompactionServiceSwitchRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp14t95lun0w****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the remote build feature.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableCompactionService *bool `json:"EnableCompactionService,omitempty" xml:"EnableCompactionService,omitempty"`
}

func (s ModifyCompactionServiceSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCompactionServiceSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyCompactionServiceSwitchRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyCompactionServiceSwitchRequest) GetEnableCompactionService() *bool {
	return s.EnableCompactionService
}

func (s *ModifyCompactionServiceSwitchRequest) SetDBClusterId(v string) *ModifyCompactionServiceSwitchRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyCompactionServiceSwitchRequest) SetEnableCompactionService(v bool) *ModifyCompactionServiceSwitchRequest {
	s.EnableCompactionService = &v
	return s
}

func (s *ModifyCompactionServiceSwitchRequest) Validate() error {
	return dara.Validate(s)
}
