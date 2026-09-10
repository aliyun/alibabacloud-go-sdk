// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStepResultOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResultId(v string) *GetStepResultOverviewRequest
	GetResultId() *string
}

type GetStepResultOverviewRequest struct {
	// The unique ID of the validation result.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
}

func (s GetStepResultOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStepResultOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetStepResultOverviewRequest) GetResultId() *string {
	return s.ResultId
}

func (s *GetStepResultOverviewRequest) SetResultId(v string) *GetStepResultOverviewRequest {
	s.ResultId = &v
	return s
}

func (s *GetStepResultOverviewRequest) Validate() error {
	return dara.Validate(s)
}
