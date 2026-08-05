// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAgentJobEstimatedCreditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobAction(v string) *GetYikeAgentJobEstimatedCreditRequest
	GetJobAction() *string
	SetJobParams(v string) *GetYikeAgentJobEstimatedCreditRequest
	GetJobParams() *string
}

type GetYikeAgentJobEstimatedCreditRequest struct {
	// The name of the task submission operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// SubmitYikeAvatarNarratorJob
	JobAction *string `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// The task request content. This is a JSON string and uses the same JobParams parameter as the task submission operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"TextType\\":2,\\"TextContent\\":\\"Today, Beijing held a press conference to announce plans to further optimize the city\\"s transportation network, including adding three new subway lines within the next three years....\\",\\"AspectRatio\\":\\"16:9\\", \\"Resolution\\":\\"720P\\", \\"OutputLanguages\\":[\\"CN\\",\\"YUE\\"]"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
}

func (s GetYikeAgentJobEstimatedCreditRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAgentJobEstimatedCreditRequest) GoString() string {
	return s.String()
}

func (s *GetYikeAgentJobEstimatedCreditRequest) GetJobAction() *string {
	return s.JobAction
}

func (s *GetYikeAgentJobEstimatedCreditRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *GetYikeAgentJobEstimatedCreditRequest) SetJobAction(v string) *GetYikeAgentJobEstimatedCreditRequest {
	s.JobAction = &v
	return s
}

func (s *GetYikeAgentJobEstimatedCreditRequest) SetJobParams(v string) *GetYikeAgentJobEstimatedCreditRequest {
	s.JobParams = &v
	return s
}

func (s *GetYikeAgentJobEstimatedCreditRequest) Validate() error {
	return dara.Validate(s)
}
