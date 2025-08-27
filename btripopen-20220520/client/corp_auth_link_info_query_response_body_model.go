// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCorpAuthLinkInfoQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CorpAuthLinkInfoQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CorpAuthLinkInfoQueryResponseBody
	GetMessage() *string
	SetModule(v *CorpAuthLinkInfoQueryResponseBodyModule) *CorpAuthLinkInfoQueryResponseBody
	GetModule() *CorpAuthLinkInfoQueryResponseBodyModule
	SetRequestId(v string) *CorpAuthLinkInfoQueryResponseBody
	GetRequestId() *string
	SetTraceId(v string) *CorpAuthLinkInfoQueryResponseBody
	GetTraceId() *string
}

type CorpAuthLinkInfoQueryResponseBody struct {
	Code      *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    *CorpAuthLinkInfoQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TraceId   *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CorpAuthLinkInfoQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CorpAuthLinkInfoQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CorpAuthLinkInfoQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CorpAuthLinkInfoQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CorpAuthLinkInfoQueryResponseBody) GetModule() *CorpAuthLinkInfoQueryResponseBodyModule {
	return s.Module
}

func (s *CorpAuthLinkInfoQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CorpAuthLinkInfoQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CorpAuthLinkInfoQueryResponseBody) SetCode(v string) *CorpAuthLinkInfoQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBody) SetMessage(v string) *CorpAuthLinkInfoQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBody) SetModule(v *CorpAuthLinkInfoQueryResponseBodyModule) *CorpAuthLinkInfoQueryResponseBody {
	s.Module = v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBody) SetRequestId(v string) *CorpAuthLinkInfoQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBody) SetTraceId(v string) *CorpAuthLinkInfoQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CorpAuthLinkInfoQueryResponseBodyModule struct {
	LinkCorps []*CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps `json:"link_corps,omitempty" xml:"link_corps,omitempty" type:"Repeated"`
	OrgCorp   *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp     `json:"org_corp,omitempty" xml:"org_corp,omitempty" type:"Struct"`
}

func (s CorpAuthLinkInfoQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CorpAuthLinkInfoQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CorpAuthLinkInfoQueryResponseBodyModule) GetLinkCorps() []*CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps {
	return s.LinkCorps
}

func (s *CorpAuthLinkInfoQueryResponseBodyModule) GetOrgCorp() *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp {
	return s.OrgCorp
}

func (s *CorpAuthLinkInfoQueryResponseBodyModule) SetLinkCorps(v []*CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) *CorpAuthLinkInfoQueryResponseBodyModule {
	s.LinkCorps = v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModule) SetOrgCorp(v *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) *CorpAuthLinkInfoQueryResponseBodyModule {
	s.OrgCorp = v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps struct {
	CorpName   *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	OpenCorpId *string `json:"open_corp_id,omitempty" xml:"open_corp_id,omitempty"`
	TrueCorpId *string `json:"true_corp_id,omitempty" xml:"true_corp_id,omitempty"`
}

func (s CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) String() string {
	return dara.Prettify(s)
}

func (s CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) GoString() string {
	return s.String()
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) GetCorpName() *string {
	return s.CorpName
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) GetOpenCorpId() *string {
	return s.OpenCorpId
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) GetTrueCorpId() *string {
	return s.TrueCorpId
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) SetCorpName(v string) *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps {
	s.CorpName = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) SetOpenCorpId(v string) *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps {
	s.OpenCorpId = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) SetTrueCorpId(v string) *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps {
	s.TrueCorpId = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleLinkCorps) Validate() error {
	return dara.Validate(s)
}

type CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp struct {
	CorpName   *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	OpenCorpId *string `json:"open_corp_id,omitempty" xml:"open_corp_id,omitempty"`
	TrueCorpId *string `json:"true_corp_id,omitempty" xml:"true_corp_id,omitempty"`
}

func (s CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) String() string {
	return dara.Prettify(s)
}

func (s CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) GoString() string {
	return s.String()
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) GetCorpName() *string {
	return s.CorpName
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) GetOpenCorpId() *string {
	return s.OpenCorpId
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) GetTrueCorpId() *string {
	return s.TrueCorpId
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) SetCorpName(v string) *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp {
	s.CorpName = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) SetOpenCorpId(v string) *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp {
	s.OpenCorpId = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) SetTrueCorpId(v string) *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp {
	s.TrueCorpId = &v
	return s
}

func (s *CorpAuthLinkInfoQueryResponseBodyModuleOrgCorp) Validate() error {
	return dara.Validate(s)
}
