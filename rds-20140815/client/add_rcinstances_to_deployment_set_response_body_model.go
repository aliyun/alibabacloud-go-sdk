// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRCInstancesToDeploymentSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddRCInstancesToDeploymentSetResponseBody
	GetRequestId() *string
	SetResults(v []*AddRCInstancesToDeploymentSetResponseBodyResults) *AddRCInstancesToDeploymentSetResponseBody
	GetResults() []*AddRCInstancesToDeploymentSetResponseBodyResults
}

type AddRCInstancesToDeploymentSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The inspection results.
	Results []*AddRCInstancesToDeploymentSetResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AddRCInstancesToDeploymentSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRCInstancesToDeploymentSetResponseBody) GoString() string {
	return s.String()
}

func (s *AddRCInstancesToDeploymentSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRCInstancesToDeploymentSetResponseBody) GetResults() []*AddRCInstancesToDeploymentSetResponseBodyResults {
	return s.Results
}

func (s *AddRCInstancesToDeploymentSetResponseBody) SetRequestId(v string) *AddRCInstancesToDeploymentSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponseBody) SetResults(v []*AddRCInstancesToDeploymentSetResponseBodyResults) *AddRCInstancesToDeploymentSetResponseBody {
	s.Results = v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponseBody) Validate() error {
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

type AddRCInstancesToDeploymentSetResponseBodyResults struct {
	// The node status. Valid values:
	//
	// 	- **activation**: Running.
	//
	// 	- **creating**: Being created.
	//
	// example:
	//
	// completed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-aaaa
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

func (s AddRCInstancesToDeploymentSetResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AddRCInstancesToDeploymentSetResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddRCInstancesToDeploymentSetResponseBodyResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddRCInstancesToDeploymentSetResponseBodyResults) GetRCInstanceId() *string {
	return s.RCInstanceId
}

func (s *AddRCInstancesToDeploymentSetResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *AddRCInstancesToDeploymentSetResponseBodyResults) SetErrorMessage(v string) *AddRCInstancesToDeploymentSetResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponseBodyResults) SetRCInstanceId(v string) *AddRCInstancesToDeploymentSetResponseBodyResults {
	s.RCInstanceId = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponseBodyResults) SetStatus(v string) *AddRCInstancesToDeploymentSetResponseBodyResults {
	s.Status = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
