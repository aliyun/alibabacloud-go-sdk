// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerhubImageRiskStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTypes(v []*string) *GetDockerhubImageRiskStatisticRequest
	GetTypes() []*string
}

type GetDockerhubImageRiskStatisticRequest struct {
	// The types of image risks to be queried.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s GetDockerhubImageRiskStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskStatisticRequest) GetTypes() []*string {
	return s.Types
}

func (s *GetDockerhubImageRiskStatisticRequest) SetTypes(v []*string) *GetDockerhubImageRiskStatisticRequest {
	s.Types = v
	return s
}

func (s *GetDockerhubImageRiskStatisticRequest) Validate() error {
	return dara.Validate(s)
}
