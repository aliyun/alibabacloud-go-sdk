// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVulWhitelistId(v int64) *GetVulWhitelistRequest
	GetVulWhitelistId() *int64
}

type GetVulWhitelistRequest struct {
	// The ID of the whitelist.
	//
	// example:
	//
	// 1275
	VulWhitelistId *int64 `json:"VulWhitelistId,omitempty" xml:"VulWhitelistId,omitempty"`
}

func (s GetVulWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *GetVulWhitelistRequest) GetVulWhitelistId() *int64 {
	return s.VulWhitelistId
}

func (s *GetVulWhitelistRequest) SetVulWhitelistId(v int64) *GetVulWhitelistRequest {
	s.VulWhitelistId = &v
	return s
}

func (s *GetVulWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
