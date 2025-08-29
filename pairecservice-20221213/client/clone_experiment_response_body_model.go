// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentId(v string) *CloneExperimentResponseBody
	GetExperimentId() *string
	SetRequestId(v string) *CloneExperimentResponseBody
	GetRequestId() *string
}

type CloneExperimentResponseBody struct {
	// example:
	//
	// 3
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CloneExperimentResponseBody) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *CloneExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneExperimentResponseBody) SetExperimentId(v string) *CloneExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CloneExperimentResponseBody) SetRequestId(v string) *CloneExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
