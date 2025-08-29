// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentGroupId(v string) *CreateExperimentGroupResponseBody
	GetExperimentGroupId() *string
	SetRequestId(v string) *CreateExperimentGroupResponseBody
	GetRequestId() *string
}

type CreateExperimentGroupResponseBody struct {
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentGroupResponseBody) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *CreateExperimentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExperimentGroupResponseBody) SetExperimentGroupId(v string) *CreateExperimentGroupResponseBody {
	s.ExperimentGroupId = &v
	return s
}

func (s *CreateExperimentGroupResponseBody) SetRequestId(v string) *CreateExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
