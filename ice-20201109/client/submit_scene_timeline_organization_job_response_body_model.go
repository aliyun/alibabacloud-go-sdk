// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneTimelineOrganizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitSceneTimelineOrganizationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitSceneTimelineOrganizationJobResponseBody
	GetRequestId() *string
}

type SubmitSceneTimelineOrganizationJobResponseBody struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSceneTimelineOrganizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneTimelineOrganizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSceneTimelineOrganizationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSceneTimelineOrganizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSceneTimelineOrganizationJobResponseBody) SetJobId(v string) *SubmitSceneTimelineOrganizationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobResponseBody) SetRequestId(v string) *SubmitSceneTimelineOrganizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
