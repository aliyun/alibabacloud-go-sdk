// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpandPolicyConfigs(v bool) *GetHttpApiRequest
	GetExpandPolicyConfigs() *bool
}

type GetHttpApiRequest struct {
	// Specifies whether to expand independent policy configurations. When omitted or set to true, a full compatible view is returned. When set to false, the ModelAPI Token throttling managed by Policy returns policy references and optional read-only plug-in status, and the rule body can be retrieved by calling GetPolicy.
	ExpandPolicyConfigs *bool `json:"expandPolicyConfigs,omitempty" xml:"expandPolicyConfigs,omitempty"`
}

func (s GetHttpApiRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpApiRequest) GoString() string {
	return s.String()
}

func (s *GetHttpApiRequest) GetExpandPolicyConfigs() *bool {
	return s.ExpandPolicyConfigs
}

func (s *GetHttpApiRequest) SetExpandPolicyConfigs(v bool) *GetHttpApiRequest {
	s.ExpandPolicyConfigs = &v
	return s
}

func (s *GetHttpApiRequest) Validate() error {
	return dara.Validate(s)
}
