// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientType(v int32) *ListVersionDistributionRequest
	GetClientType() *int32
	SetInManage(v bool) *ListVersionDistributionRequest
	GetInManage() *bool
	SetMainBizType(v string) *ListVersionDistributionRequest
	GetMainBizType() *string
	SetModel(v string) *ListVersionDistributionRequest
	GetModel() *string
	SetVersionType(v string) *ListVersionDistributionRequest
	GetVersionType() *string
}

type ListVersionDistributionRequest struct {
	// The terminal type. Valid values:
	//
	// - 1: hardware terminal.
	//
	// - 2: software terminal.
	//
	// - 3: secure browser plugin.
	//
	// - 4: GuestOS application.
	//
	// - 5: DingTalk Wuying plugin.
	//
	// - 6: cloud application component.
	//
	// - 7: Cloud Hub.
	//
	// - 8: H5.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The management status. A value of true indicates managed, and a value of false indicates unmanaged. If this parameter is not specified, all terminals are queried.
	InManage *bool `json:"InManage,omitempty" xml:"InManage,omitempty"`
	// The business type. Default value: enterprise.
	//
	// example:
	//
	// enterprise
	MainBizType *string `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
	// The terminal model.
	//
	// This parameter is required.
	//
	// example:
	//
	// AS05-2DCXG
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The version type. Valid values:
	//
	// - SYS: system version.
	//
	// - APP: application version.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYS
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s ListVersionDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVersionDistributionRequest) GoString() string {
	return s.String()
}

func (s *ListVersionDistributionRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *ListVersionDistributionRequest) GetInManage() *bool {
	return s.InManage
}

func (s *ListVersionDistributionRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *ListVersionDistributionRequest) GetModel() *string {
	return s.Model
}

func (s *ListVersionDistributionRequest) GetVersionType() *string {
	return s.VersionType
}

func (s *ListVersionDistributionRequest) SetClientType(v int32) *ListVersionDistributionRequest {
	s.ClientType = &v
	return s
}

func (s *ListVersionDistributionRequest) SetInManage(v bool) *ListVersionDistributionRequest {
	s.InManage = &v
	return s
}

func (s *ListVersionDistributionRequest) SetMainBizType(v string) *ListVersionDistributionRequest {
	s.MainBizType = &v
	return s
}

func (s *ListVersionDistributionRequest) SetModel(v string) *ListVersionDistributionRequest {
	s.Model = &v
	return s
}

func (s *ListVersionDistributionRequest) SetVersionType(v string) *ListVersionDistributionRequest {
	s.VersionType = &v
	return s
}

func (s *ListVersionDistributionRequest) Validate() error {
	return dara.Validate(s)
}
