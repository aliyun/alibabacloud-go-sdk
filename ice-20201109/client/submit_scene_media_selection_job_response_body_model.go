// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneMediaSelectionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitSceneMediaSelectionJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitSceneMediaSelectionJobResponseBody
	GetRequestId() *string
}

type SubmitSceneMediaSelectionJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSceneMediaSelectionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneMediaSelectionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSceneMediaSelectionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSceneMediaSelectionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSceneMediaSelectionJobResponseBody) SetJobId(v string) *SubmitSceneMediaSelectionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobResponseBody) SetRequestId(v string) *SubmitSceneMediaSelectionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSceneMediaSelectionJobResponseBody) Validate() error {
	return dara.Validate(s)
}
