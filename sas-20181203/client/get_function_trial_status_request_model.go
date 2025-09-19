// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionTrialStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *GetFunctionTrialStatusRequest
	GetFunctionName() *string
}

type GetFunctionTrialStatusRequest struct {
	// The name of the function module.
	//
	// example:
	//
	// trail_file_detect_api_reward
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s GetFunctionTrialStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionTrialStatusRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionTrialStatusRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionTrialStatusRequest) SetFunctionName(v string) *GetFunctionTrialStatusRequest {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionTrialStatusRequest) Validate() error {
	return dara.Validate(s)
}
