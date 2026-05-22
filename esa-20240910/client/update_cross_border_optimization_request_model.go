// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrossBorderOptimizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *UpdateCrossBorderOptimizationRequest
	GetEnable() *string
	SetSiteId(v int64) *UpdateCrossBorderOptimizationRequest
	GetSiteId() *int64
}

type UpdateCrossBorderOptimizationRequest struct {
	// Whether to enable Chinese mainland network access optimization. By default, it is disabled. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateCrossBorderOptimizationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossBorderOptimizationRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrossBorderOptimizationRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateCrossBorderOptimizationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateCrossBorderOptimizationRequest) SetEnable(v string) *UpdateCrossBorderOptimizationRequest {
	s.Enable = &v
	return s
}

func (s *UpdateCrossBorderOptimizationRequest) SetSiteId(v int64) *UpdateCrossBorderOptimizationRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateCrossBorderOptimizationRequest) Validate() error {
	return dara.Validate(s)
}
