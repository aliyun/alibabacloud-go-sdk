// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunApiTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *RunApiTemplateRequest
	GetApiName() *string
	SetClientToken(v string) *RunApiTemplateRequest
	GetClientToken() *string
	SetRegionId(v string) *RunApiTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *RunApiTemplateRequest
	GetTemplateId() *string
}

type RunApiTemplateRequest struct {
	// 接口名。
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateCluster
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// 幂等客户端TOKEN。
	//
	// example:
	//
	// A7D960FA-6DBA-5E07-8746-A63E3E4D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 地域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 集群模板id。
	//
	// This parameter is required.
	//
	// example:
	//
	// AT-****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s RunApiTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s RunApiTemplateRequest) GoString() string {
	return s.String()
}

func (s *RunApiTemplateRequest) GetApiName() *string {
	return s.ApiName
}

func (s *RunApiTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunApiTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunApiTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *RunApiTemplateRequest) SetApiName(v string) *RunApiTemplateRequest {
	s.ApiName = &v
	return s
}

func (s *RunApiTemplateRequest) SetClientToken(v string) *RunApiTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *RunApiTemplateRequest) SetRegionId(v string) *RunApiTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *RunApiTemplateRequest) SetTemplateId(v string) *RunApiTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *RunApiTemplateRequest) Validate() error {
	return dara.Validate(s)
}
