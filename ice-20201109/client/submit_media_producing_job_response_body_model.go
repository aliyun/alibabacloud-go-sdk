// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaProducingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitMediaProducingJobResponseBody
	GetJobId() *string
	SetMediaId(v string) *SubmitMediaProducingJobResponseBody
	GetMediaId() *string
	SetProjectId(v string) *SubmitMediaProducingJobResponseBody
	GetProjectId() *string
	SetRequestId(v string) *SubmitMediaProducingJobResponseBody
	GetRequestId() *string
	SetVodMediaId(v string) *SubmitMediaProducingJobResponseBody
	GetVodMediaId() *string
}

type SubmitMediaProducingJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The media asset ID of the output file.
	//
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the editing project.
	//
	// example:
	//
	// ****b4549d46c88681030f6e****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The media asset ID of the output file in ApsaraVideo VOD if the output file is stored in ApsaraVideo VOD.
	//
	// example:
	//
	// ****d8s4h75ci975745c14b****
	VodMediaId *string `json:"VodMediaId,omitempty" xml:"VodMediaId,omitempty"`
}

func (s SubmitMediaProducingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaProducingJobResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitMediaProducingJobResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitMediaProducingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaProducingJobResponseBody) GetVodMediaId() *string {
	return s.VodMediaId
}

func (s *SubmitMediaProducingJobResponseBody) SetJobId(v string) *SubmitMediaProducingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetMediaId(v string) *SubmitMediaProducingJobResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetProjectId(v string) *SubmitMediaProducingJobResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetRequestId(v string) *SubmitMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetVodMediaId(v string) *SubmitMediaProducingJobResponseBody {
	s.VodMediaId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
