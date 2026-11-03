// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentId(v string) *CreateExperimentResponseBody
	GetExperimentId() *string
	SetRequestId(v string) *CreateExperimentResponseBody
	GetRequestId() *string
}

type CreateExperimentResponseBody struct {
	// The Job ID returned if the creation API is implemented asynchronously.
	//
	// example:
	//
	// ID of the experiment created by the API request
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// Request ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponseBody) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *CreateExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExperimentResponseBody) SetExperimentId(v string) *CreateExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CreateExperimentResponseBody) SetRequestId(v string) *CreateExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
