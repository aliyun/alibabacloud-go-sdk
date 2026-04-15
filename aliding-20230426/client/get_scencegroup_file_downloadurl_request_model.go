// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScencegroupFileDownloadurlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadCode(v string) *GetScencegroupFileDownloadurlRequest
	GetDownloadCode() *string
	SetTenantContext(v *GetScencegroupFileDownloadurlRequestTenantContext) *GetScencegroupFileDownloadurlRequest
	GetTenantContext() *GetScencegroupFileDownloadurlRequestTenantContext
}

type GetScencegroupFileDownloadurlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc*****xyz
	DownloadCode  *string                                            `json:"DownloadCode,omitempty" xml:"DownloadCode,omitempty"`
	TenantContext *GetScencegroupFileDownloadurlRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetScencegroupFileDownloadurlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlRequest) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlRequest) GetDownloadCode() *string {
	return s.DownloadCode
}

func (s *GetScencegroupFileDownloadurlRequest) GetTenantContext() *GetScencegroupFileDownloadurlRequestTenantContext {
	return s.TenantContext
}

func (s *GetScencegroupFileDownloadurlRequest) SetDownloadCode(v string) *GetScencegroupFileDownloadurlRequest {
	s.DownloadCode = &v
	return s
}

func (s *GetScencegroupFileDownloadurlRequest) SetTenantContext(v *GetScencegroupFileDownloadurlRequestTenantContext) *GetScencegroupFileDownloadurlRequest {
	s.TenantContext = v
	return s
}

func (s *GetScencegroupFileDownloadurlRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScencegroupFileDownloadurlRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScencegroupFileDownloadurlRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScencegroupFileDownloadurlRequestTenantContext) SetTenantId(v string) *GetScencegroupFileDownloadurlRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetScencegroupFileDownloadurlRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
