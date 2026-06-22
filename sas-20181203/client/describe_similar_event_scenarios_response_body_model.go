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
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// FDF7B8D9-8493-4B90-8D13-E0C1FFCE5F97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of handling scenarios for alerts of the same type.
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
	// The code of the handling scenario. Valid values:
	//
	// - **default**: same alerting type
	//
	// - **same_file_content**: same file content rule
	//
	// - **same_ip**: same IP rule
	//
	// - **same_url**: same URL rule.
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
