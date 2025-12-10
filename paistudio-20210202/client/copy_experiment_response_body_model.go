// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentId(v string) *CopyExperimentResponseBody
	GetExperimentId() *string
	SetRequestId(v string) *CopyExperimentResponseBody
	GetRequestId() *string
}

type CopyExperimentResponseBody struct {
	// example:
	//
	// draft-rbvg5wzljzjhc9ks92
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 6D161C28-9CB1-584E-8DD5-64441E32A5B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CopyExperimentResponseBody) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *CopyExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyExperimentResponseBody) SetExperimentId(v string) *CopyExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CopyExperimentResponseBody) SetRequestId(v string) *CopyExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
