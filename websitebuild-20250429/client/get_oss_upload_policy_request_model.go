// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScenario(v string) *GetOssUploadPolicyRequest
	GetScenario() *string
}

type GetOssUploadPolicyRequest struct {
	// example:
	//
	// default
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
}

func (s GetOssUploadPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadPolicyRequest) GetScenario() *string {
	return s.Scenario
}

func (s *GetOssUploadPolicyRequest) SetScenario(v string) *GetOssUploadPolicyRequest {
	s.Scenario = &v
	return s
}

func (s *GetOssUploadPolicyRequest) Validate() error {
	return dara.Validate(s)
}
