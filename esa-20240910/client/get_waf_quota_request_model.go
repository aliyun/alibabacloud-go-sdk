// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPaths(v string) *GetWafQuotaRequest
	GetPaths() *string
}

type GetWafQuotaRequest struct {
	// The path of the WAF quota. Valid values:
	//
	// - managed_rules_group: The quota for managed rule groups.
	//
	// - list: The quota for custom lists.
	//
	// - page: The quota for custom response pages.
	//
	// - scene_policy: The quota for scene-based protection.
	//
	// example:
	//
	// page
	Paths *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
}

func (s GetWafQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetWafQuotaRequest) GetPaths() *string {
	return s.Paths
}

func (s *GetWafQuotaRequest) SetPaths(v string) *GetWafQuotaRequest {
	s.Paths = &v
	return s
}

func (s *GetWafQuotaRequest) Validate() error {
	return dara.Validate(s)
}
