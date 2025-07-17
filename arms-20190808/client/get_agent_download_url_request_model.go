// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetAgentDownloadUrlRequest
	GetRegionId() *string
}

type GetAgentDownloadUrlRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAgentDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAgentDownloadUrlRequest) SetRegionId(v string) *GetAgentDownloadUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetAgentDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
