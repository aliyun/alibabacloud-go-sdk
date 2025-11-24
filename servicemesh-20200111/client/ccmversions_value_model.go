// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCCMVersionsValue interface {
	dara.Model
	String() string
	GoString() string
	SetQueryState(v string) *CCMVersionsValue
	GetQueryState() *string
	SetVersion(v string) *CCMVersionsValue
	GetVersion() *string
	SetSLBGracefulDrainSupported(v bool) *CCMVersionsValue
	GetSLBGracefulDrainSupported() *bool
	SetClusterId(v string) *CCMVersionsValue
	GetClusterId() *string
	SetMessage(v string) *CCMVersionsValue
	GetMessage() *string
	SetSLBGracefulDrainSupport(v bool) *CCMVersionsValue
	GetSLBGracefulDrainSupport() *bool
}

type CCMVersionsValue struct {
	// The status of the query. Valid values:
	//
	// 	- `time_out`: The query times out.
	//
	// 	- `failed`: The query fails.
	//
	// 	- `Empty string`: The query is successful.
	//
	// example:
	//
	// time_out
	QueryState *string `json:"QueryState,omitempty" xml:"QueryState,omitempty"`
	// The CCM version.
	//
	// example:
	//
	// v2.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// Indicates whether Classic Load Balancer (CLB) graceful shutdown is supported.
	//
	// example:
	//
	// true
	SLBGracefulDrainSupported *bool `json:"SLBGracefulDrainSupported,omitempty" xml:"SLBGracefulDrainSupported,omitempty"`
	// The ID of the cluster instance on the data plane.
	//
	// example:
	//
	// cfbb81b9b51a44b299349xxxxxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The additional information that is returned when the query fails. This parameter is empty if the query is successful.
	//
	// example:
	//
	// timeout error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Deprecated
	SLBGracefulDrainSupport *bool `json:"SLBGracefulDrainSupport,omitempty" xml:"SLBGracefulDrainSupport,omitempty"`
}

func (s CCMVersionsValue) String() string {
	return dara.Prettify(s)
}

func (s CCMVersionsValue) GoString() string {
	return s.String()
}

func (s *CCMVersionsValue) GetQueryState() *string {
	return s.QueryState
}

func (s *CCMVersionsValue) GetVersion() *string {
	return s.Version
}

func (s *CCMVersionsValue) GetSLBGracefulDrainSupported() *bool {
	return s.SLBGracefulDrainSupported
}

func (s *CCMVersionsValue) GetClusterId() *string {
	return s.ClusterId
}

func (s *CCMVersionsValue) GetMessage() *string {
	return s.Message
}

func (s *CCMVersionsValue) GetSLBGracefulDrainSupport() *bool {
	return s.SLBGracefulDrainSupport
}

func (s *CCMVersionsValue) SetQueryState(v string) *CCMVersionsValue {
	s.QueryState = &v
	return s
}

func (s *CCMVersionsValue) SetVersion(v string) *CCMVersionsValue {
	s.Version = &v
	return s
}

func (s *CCMVersionsValue) SetSLBGracefulDrainSupported(v bool) *CCMVersionsValue {
	s.SLBGracefulDrainSupported = &v
	return s
}

func (s *CCMVersionsValue) SetClusterId(v string) *CCMVersionsValue {
	s.ClusterId = &v
	return s
}

func (s *CCMVersionsValue) SetMessage(v string) *CCMVersionsValue {
	s.Message = &v
	return s
}

func (s *CCMVersionsValue) SetSLBGracefulDrainSupport(v bool) *CCMVersionsValue {
	s.SLBGracefulDrainSupport = &v
	return s
}

func (s *CCMVersionsValue) Validate() error {
	return dara.Validate(s)
}
