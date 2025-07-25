// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrainingJobResponseBody
	GetRequestId() *string
	SetTrainingJobId(v string) *CreateTrainingJobResponseBody
	GetTrainingJobId() *string
}

type CreateTrainingJobResponseBody struct {
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// traineyfz0m2hsfv
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
}

func (s CreateTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrainingJobResponseBody) GetTrainingJobId() *string {
	return s.TrainingJobId
}

func (s *CreateTrainingJobResponseBody) SetRequestId(v string) *CreateTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrainingJobResponseBody) SetTrainingJobId(v string) *CreateTrainingJobResponseBody {
	s.TrainingJobId = &v
	return s
}

func (s *CreateTrainingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
