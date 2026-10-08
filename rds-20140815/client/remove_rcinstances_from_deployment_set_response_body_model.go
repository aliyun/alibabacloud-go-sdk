// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRCInstancesFromDeploymentSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveRCInstancesFromDeploymentSetResponseBody
	GetRequestId() *string
	SetResults(v []*RemoveRCInstancesFromDeploymentSetResponseBodyResults) *RemoveRCInstancesFromDeploymentSetResponseBody
	GetResults() []*RemoveRCInstancesFromDeploymentSetResponseBodyResults
}

type RemoveRCInstancesFromDeploymentSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C816A4BF-A6EC-4722-95F9-2055859CCFD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The call results of the operation.
	Results []*RemoveRCInstancesFromDeploymentSetResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RemoveRCInstancesFromDeploymentSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveRCInstancesFromDeploymentSetResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBody) GetResults() []*RemoveRCInstancesFromDeploymentSetResponseBodyResults {
	return s.Results
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBody) SetRequestId(v string) *RemoveRCInstancesFromDeploymentSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBody) SetResults(v []*RemoveRCInstancesFromDeploymentSetResponseBodyResults) *RemoveRCInstancesFromDeploymentSetResponseBody {
	s.Results = v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RemoveRCInstancesFromDeploymentSetResponseBodyResults struct {
	// The instance ID.
	//
	// example:
	//
	// rc-w9htiydssds
	RCInstanceId *string `json:"RCInstanceId,omitempty" xml:"RCInstanceId,omitempty"`
	// The node status. Valid values:
	//
	// 	- **Success**: Succeeded.
	//
	// 	- **Failed**: Failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RemoveRCInstancesFromDeploymentSetResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s RemoveRCInstancesFromDeploymentSetResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBodyResults) GetRCInstanceId() *string {
	return s.RCInstanceId
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBodyResults) SetRCInstanceId(v string) *RemoveRCInstancesFromDeploymentSetResponseBodyResults {
	s.RCInstanceId = &v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBodyResults) SetStatus(v string) *RemoveRCInstancesFromDeploymentSetResponseBodyResults {
	s.Status = &v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
