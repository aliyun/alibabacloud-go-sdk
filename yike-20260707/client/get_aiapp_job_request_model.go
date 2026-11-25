// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIAppJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetAIAppJobRequest
	GetJobId() *string
	SetOnlyOutputs(v bool) *GetAIAppJobRequest
	GetOnlyOutputs() *bool
}

type GetAIAppJobRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Specifies whether to return only the task results. If you set this parameter to true, the response does not include the results of each node.
	//
	// example:
	//
	// true
	OnlyOutputs *bool `json:"OnlyOutputs,omitempty" xml:"OnlyOutputs,omitempty"`
}

func (s GetAIAppJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIAppJobRequest) GoString() string {
	return s.String()
}

func (s *GetAIAppJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetAIAppJobRequest) GetOnlyOutputs() *bool {
	return s.OnlyOutputs
}

func (s *GetAIAppJobRequest) SetJobId(v string) *GetAIAppJobRequest {
	s.JobId = &v
	return s
}

func (s *GetAIAppJobRequest) SetOnlyOutputs(v bool) *GetAIAppJobRequest {
	s.OnlyOutputs = &v
	return s
}

func (s *GetAIAppJobRequest) Validate() error {
	return dara.Validate(s)
}
