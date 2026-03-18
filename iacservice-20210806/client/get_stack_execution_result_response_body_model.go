// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackExecutionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackExecutionResultResponseBody
	GetRequestId() *string
	SetStackResults(v []*GetStackExecutionResultResponseBodyStackResults) *GetStackExecutionResultResponseBody
	GetStackResults() []*GetStackExecutionResultResponseBodyStackResults
	SetTriggerId(v string) *GetStackExecutionResultResponseBody
	GetTriggerId() *string
}

type GetStackExecutionResultResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F2D40488-3F74-568B-87EC-1C04D098DF8B
	RequestId    *string                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	StackResults []*GetStackExecutionResultResponseBodyStackResults `json:"stackResults,omitempty" xml:"stackResults,omitempty" type:"Repeated"`
	// example:
	//
	// event-xxx
	TriggerId *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
}

func (s GetStackExecutionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackExecutionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackExecutionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackExecutionResultResponseBody) GetStackResults() []*GetStackExecutionResultResponseBodyStackResults {
	return s.StackResults
}

func (s *GetStackExecutionResultResponseBody) GetTriggerId() *string {
	return s.TriggerId
}

func (s *GetStackExecutionResultResponseBody) SetRequestId(v string) *GetStackExecutionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackExecutionResultResponseBody) SetStackResults(v []*GetStackExecutionResultResponseBodyStackResults) *GetStackExecutionResultResponseBody {
	s.StackResults = v
	return s
}

func (s *GetStackExecutionResultResponseBody) SetTriggerId(v string) *GetStackExecutionResultResponseBody {
	s.TriggerId = &v
	return s
}

func (s *GetStackExecutionResultResponseBody) Validate() error {
	if s.StackResults != nil {
		for _, item := range s.StackResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStackExecutionResultResponseBodyStackResults struct {
	Deployments []*GetStackExecutionResultResponseBodyStackResultsDeployments `json:"deployments,omitempty" xml:"deployments,omitempty" type:"Repeated"`
	// example:
	//
	// No corresponding Stack found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// stack-al181av2bloah5s53hacbp4
	StackId *string `json:"stackId,omitempty" xml:"stackId,omitempty"`
	// example:
	//
	// stack-demo
	StackName *string `json:"stackName,omitempty" xml:"stackName,omitempty"`
	// example:
	//
	// Deployed
	StackStatus *string `json:"stackStatus,omitempty" xml:"stackStatus,omitempty"`
}

func (s GetStackExecutionResultResponseBodyStackResults) String() string {
	return dara.Prettify(s)
}

func (s GetStackExecutionResultResponseBodyStackResults) GoString() string {
	return s.String()
}

func (s *GetStackExecutionResultResponseBodyStackResults) GetDeployments() []*GetStackExecutionResultResponseBodyStackResultsDeployments {
	return s.Deployments
}

func (s *GetStackExecutionResultResponseBodyStackResults) GetMessage() *string {
	return s.Message
}

func (s *GetStackExecutionResultResponseBodyStackResults) GetStackId() *string {
	return s.StackId
}

func (s *GetStackExecutionResultResponseBodyStackResults) GetStackName() *string {
	return s.StackName
}

func (s *GetStackExecutionResultResponseBodyStackResults) GetStackStatus() *string {
	return s.StackStatus
}

func (s *GetStackExecutionResultResponseBodyStackResults) SetDeployments(v []*GetStackExecutionResultResponseBodyStackResultsDeployments) *GetStackExecutionResultResponseBodyStackResults {
	s.Deployments = v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResults) SetMessage(v string) *GetStackExecutionResultResponseBodyStackResults {
	s.Message = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResults) SetStackId(v string) *GetStackExecutionResultResponseBodyStackResults {
	s.StackId = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResults) SetStackName(v string) *GetStackExecutionResultResponseBodyStackResults {
	s.StackName = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResults) SetStackStatus(v string) *GetStackExecutionResultResponseBodyStackResults {
	s.StackStatus = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResults) Validate() error {
	if s.Deployments != nil {
		for _, item := range s.Deployments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStackExecutionResultResponseBodyStackResultsDeployments struct {
	// example:
	//
	// prod
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// example:
	//
	// Service returned null result
	JobResult *string `json:"jobResult,omitempty" xml:"jobResult,omitempty"`
	// example:
	//
	// Applied
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// https://iacnext.console.aliyun.com/stack/stack-al181av2bloah5s53hacbp4/details?deploymentName=production&deploymentNo=6&configVersion=v1
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetStackExecutionResultResponseBodyStackResultsDeployments) String() string {
	return dara.Prettify(s)
}

func (s GetStackExecutionResultResponseBodyStackResultsDeployments) GoString() string {
	return s.String()
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) GetJobResult() *string {
	return s.JobResult
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) GetStatus() *string {
	return s.Status
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) GetUrl() *string {
	return s.Url
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) SetDeploymentName(v string) *GetStackExecutionResultResponseBodyStackResultsDeployments {
	s.DeploymentName = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) SetJobResult(v string) *GetStackExecutionResultResponseBodyStackResultsDeployments {
	s.JobResult = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) SetStatus(v string) *GetStackExecutionResultResponseBodyStackResultsDeployments {
	s.Status = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) SetUrl(v string) *GetStackExecutionResultResponseBodyStackResultsDeployments {
	s.Url = &v
	return s
}

func (s *GetStackExecutionResultResponseBodyStackResultsDeployments) Validate() error {
	return dara.Validate(s)
}
