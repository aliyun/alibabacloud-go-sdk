// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgRunSensIdentifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEsMetaParamsShrink(v string) *DsgRunSensIdentifyShrinkRequest
	GetEsMetaParamsShrink() *string
	SetTenantId(v string) *DsgRunSensIdentifyShrinkRequest
	GetTenantId() *string
}

type DsgRunSensIdentifyShrinkRequest struct {
	// The parameters that you need to configure to scan specified metadata.
	EsMetaParamsShrink *string `json:"EsMetaParams,omitempty" xml:"EsMetaParams,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DsgRunSensIdentifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgRunSensIdentifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgRunSensIdentifyShrinkRequest) GetEsMetaParamsShrink() *string {
	return s.EsMetaParamsShrink
}

func (s *DsgRunSensIdentifyShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DsgRunSensIdentifyShrinkRequest) SetEsMetaParamsShrink(v string) *DsgRunSensIdentifyShrinkRequest {
	s.EsMetaParamsShrink = &v
	return s
}

func (s *DsgRunSensIdentifyShrinkRequest) SetTenantId(v string) *DsgRunSensIdentifyShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *DsgRunSensIdentifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
