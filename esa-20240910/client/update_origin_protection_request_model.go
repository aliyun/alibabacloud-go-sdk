// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfirmIPList(v string) *UpdateOriginProtectionRequest
	GetAutoConfirmIPList() *string
	SetOriginConverge(v string) *UpdateOriginProtectionRequest
	GetOriginConverge() *string
	SetSiteId(v int64) *UpdateOriginProtectionRequest
	GetSiteId() *int64
}

type UpdateOriginProtectionRequest struct {
	// example:
	//
	// off
	AutoConfirmIPList *string `json:"AutoConfirmIPList,omitempty" xml:"AutoConfirmIPList,omitempty"`
	// The IP convergence status.
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
	OriginConverge *string `json:"OriginConverge,omitempty" xml:"OriginConverge,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateOriginProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginProtectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOriginProtectionRequest) GetAutoConfirmIPList() *string {
	return s.AutoConfirmIPList
}

func (s *UpdateOriginProtectionRequest) GetOriginConverge() *string {
	return s.OriginConverge
}

func (s *UpdateOriginProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateOriginProtectionRequest) SetAutoConfirmIPList(v string) *UpdateOriginProtectionRequest {
	s.AutoConfirmIPList = &v
	return s
}

func (s *UpdateOriginProtectionRequest) SetOriginConverge(v string) *UpdateOriginProtectionRequest {
	s.OriginConverge = &v
	return s
}

func (s *UpdateOriginProtectionRequest) SetSiteId(v int64) *UpdateOriginProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateOriginProtectionRequest) Validate() error {
	return dara.Validate(s)
}
