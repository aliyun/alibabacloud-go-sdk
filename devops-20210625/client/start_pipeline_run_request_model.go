// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *StartPipelineRunRequest
	GetParams() *string
}

type StartPipelineRunRequest struct {
	// example:
	//
	// {     "branchModeBranchs":[         "branch1",         "branch2"     ],     "envs":{         "k1":"v1",         "k2":"v2",         "k3":"v3"     },     "runningBranchs":{         "https://codeup.aliyun.com/60c1abb32c5969c370c5fcd0/Codeup-Demo.git":"master1"     },     "runningTags":{         "https://codeup.aliyun.com/60c1abb32c5969c370c5fcd0/Codeup-Demo.git":"1.0"     } }
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s StartPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *StartPipelineRunRequest) GetParams() *string {
	return s.Params
}

func (s *StartPipelineRunRequest) SetParams(v string) *StartPipelineRunRequest {
	s.Params = &v
	return s
}

func (s *StartPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
