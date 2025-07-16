// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipartFileUploadInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v *GetMultipartFileUploadInfosRequestOption) *GetMultipartFileUploadInfosRequest
	GetOption() *GetMultipartFileUploadInfosRequestOption
	SetPartNumbers(v []*int32) *GetMultipartFileUploadInfosRequest
	GetPartNumbers() []*int32
	SetTenantContext(v *GetMultipartFileUploadInfosRequestTenantContext) *GetMultipartFileUploadInfosRequest
	GetTenantContext() *GetMultipartFileUploadInfosRequestTenantContext
	SetUploadKey(v string) *GetMultipartFileUploadInfosRequest
	GetUploadKey() *string
}

type GetMultipartFileUploadInfosRequest struct {
	Option        *GetMultipartFileUploadInfosRequestOption        `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	PartNumbers   []*int32                                         `json:"PartNumbers,omitempty" xml:"PartNumbers,omitempty" type:"Repeated"`
	TenantContext *GetMultipartFileUploadInfosRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// hwHPAAAAAipHxxxxx
	UploadKey *string `json:"UploadKey,omitempty" xml:"UploadKey,omitempty"`
}

func (s GetMultipartFileUploadInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosRequest) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosRequest) GetOption() *GetMultipartFileUploadInfosRequestOption {
	return s.Option
}

func (s *GetMultipartFileUploadInfosRequest) GetPartNumbers() []*int32 {
	return s.PartNumbers
}

func (s *GetMultipartFileUploadInfosRequest) GetTenantContext() *GetMultipartFileUploadInfosRequestTenantContext {
	return s.TenantContext
}

func (s *GetMultipartFileUploadInfosRequest) GetUploadKey() *string {
	return s.UploadKey
}

func (s *GetMultipartFileUploadInfosRequest) SetOption(v *GetMultipartFileUploadInfosRequestOption) *GetMultipartFileUploadInfosRequest {
	s.Option = v
	return s
}

func (s *GetMultipartFileUploadInfosRequest) SetPartNumbers(v []*int32) *GetMultipartFileUploadInfosRequest {
	s.PartNumbers = v
	return s
}

func (s *GetMultipartFileUploadInfosRequest) SetTenantContext(v *GetMultipartFileUploadInfosRequestTenantContext) *GetMultipartFileUploadInfosRequest {
	s.TenantContext = v
	return s
}

func (s *GetMultipartFileUploadInfosRequest) SetUploadKey(v string) *GetMultipartFileUploadInfosRequest {
	s.UploadKey = &v
	return s
}

func (s *GetMultipartFileUploadInfosRequest) Validate() error {
	return dara.Validate(s)
}

type GetMultipartFileUploadInfosRequestOption struct {
	// example:
	//
	// true
	PreferIntranet *bool `json:"PreferIntranet,omitempty" xml:"PreferIntranet,omitempty"`
}

func (s GetMultipartFileUploadInfosRequestOption) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosRequestOption) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosRequestOption) GetPreferIntranet() *bool {
	return s.PreferIntranet
}

func (s *GetMultipartFileUploadInfosRequestOption) SetPreferIntranet(v bool) *GetMultipartFileUploadInfosRequestOption {
	s.PreferIntranet = &v
	return s
}

func (s *GetMultipartFileUploadInfosRequestOption) Validate() error {
	return dara.Validate(s)
}

type GetMultipartFileUploadInfosRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMultipartFileUploadInfosRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMultipartFileUploadInfosRequestTenantContext) SetTenantId(v string) *GetMultipartFileUploadInfosRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetMultipartFileUploadInfosRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
