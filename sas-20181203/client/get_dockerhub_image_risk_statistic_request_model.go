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
	// The image risk types to query. This parameter is required. You must specify at least one valid value. If this parameter is not specified or all specified values are invalid, the API returns HTTP 400 and the error code -101.
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
