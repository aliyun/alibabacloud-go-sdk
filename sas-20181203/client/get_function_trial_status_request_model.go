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
	// The name of the feature module. Valid values:
	//
	// - trail_file_detect_api_reward: File detection API trial.
	//
	// - trail_honeypot_reward: Honeypot trial.
	//
	// > This parameter is required. If it is not provided, the API returns an HTTP 400 error with the error code -101.
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
