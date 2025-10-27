// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarEventScenariosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSimilarEventScenariosResponseBody
	GetRequestId() *string
	SetScenarios(v []*DescribeSimilarEventScenariosResponseBodyScenarios) *DescribeSimilarEventScenariosResponseBody
	GetScenarios() []*DescribeSimilarEventScenariosResponseBodyScenarios
}

type DescribeSimilarEventScenariosResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FDF7B8D9-8493-4B90-8D13-E0C1FFCE5F97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scenarios in which alerts triggered by the same rule or rules of the same type are handled.
	Scenarios []*DescribeSimilarEventScenariosResponseBodyScenarios `json:"Scenarios,omitempty" xml:"Scenarios,omitempty" type:"Repeated"`
}

func (s DescribeSimilarEventScenariosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarEventScenariosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSimilarEventScenariosResponseBody) GetScenarios() []*DescribeSimilarEventScenariosResponseBodyScenarios {
	return s.Scenarios
}

func (s *DescribeSimilarEventScenariosResponseBody) SetRequestId(v string) *DescribeSimilarEventScenariosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimilarEventScenariosResponseBody) SetScenarios(v []*DescribeSimilarEventScenariosResponseBodyScenarios) *DescribeSimilarEventScenariosResponseBody {
	s.Scenarios = v
	return s
}

func (s *DescribeSimilarEventScenariosResponseBody) Validate() error {
	if s.Scenarios != nil {
		for _, item := range s.Scenarios {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSimilarEventScenariosResponseBodyScenarios struct {
	// The code of the scenario. Valid values:
	//
	// 	- **default**: the same alert type
	//
	// 	- **same_file_content**: the same file content rule.
	//
	// 	- **same_ip**: the same IP address rule.
	//
	// 	- **same_url**: the same URL rule.
	//
	// example:
	//
	// same_url
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeSimilarEventScenariosResponseBodyScenarios) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarEventScenariosResponseBodyScenarios) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosResponseBodyScenarios) GetCode() *string {
	return s.Code
}

func (s *DescribeSimilarEventScenariosResponseBodyScenarios) SetCode(v string) *DescribeSimilarEventScenariosResponseBodyScenarios {
	s.Code = &v
	return s
}

func (s *DescribeSimilarEventScenariosResponseBodyScenarios) Validate() error {
	return dara.Validate(s)
}
