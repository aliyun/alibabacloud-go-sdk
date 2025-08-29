// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneExperimentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentGroupId(v string) *CloneExperimentGroupResponseBody
	GetExperimentGroupId() *string
	SetRequestId(v string) *CloneExperimentGroupResponseBody
	GetRequestId() *string
}

type CloneExperimentGroupResponseBody struct {
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 12A65C6C-AFA1-59B2-9A66-A9E0BB73F0E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneExperimentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloneExperimentGroupResponseBody) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *CloneExperimentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneExperimentGroupResponseBody) SetExperimentGroupId(v string) *CloneExperimentGroupResponseBody {
	s.ExperimentGroupId = &v
	return s
}

func (s *CloneExperimentGroupResponseBody) SetRequestId(v string) *CloneExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneExperimentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
