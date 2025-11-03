// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoConfirmIPList(v string) *CreateOriginProtectionRequest
	GetAutoConfirmIPList() *string
	SetSiteId(v int64) *CreateOriginProtectionRequest
	GetSiteId() *int64
}

type CreateOriginProtectionRequest struct {
	// example:
	//
	// off
	AutoConfirmIPList *string `json:"AutoConfirmIPList,omitempty" xml:"AutoConfirmIPList,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateOriginProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginProtectionRequest) GoString() string {
	return s.String()
}

func (s *CreateOriginProtectionRequest) GetAutoConfirmIPList() *string {
	return s.AutoConfirmIPList
}

func (s *CreateOriginProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateOriginProtectionRequest) SetAutoConfirmIPList(v string) *CreateOriginProtectionRequest {
	s.AutoConfirmIPList = &v
	return s
}

func (s *CreateOriginProtectionRequest) SetSiteId(v int64) *CreateOriginProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *CreateOriginProtectionRequest) Validate() error {
	return dara.Validate(s)
}
