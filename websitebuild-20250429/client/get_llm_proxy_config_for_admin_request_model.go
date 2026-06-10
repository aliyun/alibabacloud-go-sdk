// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmProxyConfigForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetLlmProxyConfigForAdminRequest
	GetBizId() *string
	SetCapability(v string) *GetLlmProxyConfigForAdminRequest
	GetCapability() *string
}

type GetLlmProxyConfigForAdminRequest struct {
	// Business ID
	//
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Capability Type: llm, image, video
	//
	// example:
	//
	// understand
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
}

func (s GetLlmProxyConfigForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLlmProxyConfigForAdminRequest) GoString() string {
	return s.String()
}

func (s *GetLlmProxyConfigForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetLlmProxyConfigForAdminRequest) GetCapability() *string {
	return s.Capability
}

func (s *GetLlmProxyConfigForAdminRequest) SetBizId(v string) *GetLlmProxyConfigForAdminRequest {
	s.BizId = &v
	return s
}

func (s *GetLlmProxyConfigForAdminRequest) SetCapability(v string) *GetLlmProxyConfigForAdminRequest {
	s.Capability = &v
	return s
}

func (s *GetLlmProxyConfigForAdminRequest) Validate() error {
	return dara.Validate(s)
}
