// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerhubImageRiskRankInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTypes(v []*string) *GetDockerhubImageRiskRankInfoRequest
	GetTypes() []*string
}

type GetDockerhubImageRiskRankInfoRequest struct {
	// The dimension types.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s GetDockerhubImageRiskRankInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskRankInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskRankInfoRequest) GetTypes() []*string {
	return s.Types
}

func (s *GetDockerhubImageRiskRankInfoRequest) SetTypes(v []*string) *GetDockerhubImageRiskRankInfoRequest {
	s.Types = v
	return s
}

func (s *GetDockerhubImageRiskRankInfoRequest) Validate() error {
	return dara.Validate(s)
}
