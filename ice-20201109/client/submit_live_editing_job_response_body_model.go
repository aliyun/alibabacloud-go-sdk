// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveEditingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitLiveEditingJobResponseBody
	GetJobId() *string
	SetMediaId(v string) *SubmitLiveEditingJobResponseBody
	GetMediaId() *string
	SetMediaURL(v string) *SubmitLiveEditingJobResponseBody
	GetMediaURL() *string
	SetProjectId(v string) *SubmitLiveEditingJobResponseBody
	GetProjectId() *string
	SetRequestId(v string) *SubmitLiveEditingJobResponseBody
	GetRequestId() *string
	SetVodMediaId(v string) *SubmitLiveEditingJobResponseBody
	GetVodMediaId() *string
}

type SubmitLiveEditingJobResponseBody struct {
	// The ID of the live editing job.
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
	// The URL of the output file.
	//
	// example:
	//
	// http://test-bucket.cn-shanghai.aliyuncs.com/test.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The ID of the live editing project.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
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
	// ****d7578s4h75ci945c14b****
	VodMediaId *string `json:"VodMediaId,omitempty" xml:"VodMediaId,omitempty"`
}

func (s SubmitLiveEditingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveEditingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitLiveEditingJobResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitLiveEditingJobResponseBody) GetMediaURL() *string {
	return s.MediaURL
}

func (s *SubmitLiveEditingJobResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitLiveEditingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitLiveEditingJobResponseBody) GetVodMediaId() *string {
	return s.VodMediaId
}

func (s *SubmitLiveEditingJobResponseBody) SetJobId(v string) *SubmitLiveEditingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitLiveEditingJobResponseBody) SetMediaId(v string) *SubmitLiveEditingJobResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitLiveEditingJobResponseBody) SetMediaURL(v string) *SubmitLiveEditingJobResponseBody {
	s.MediaURL = &v
	return s
}

func (s *SubmitLiveEditingJobResponseBody) SetProjectId(v string) *SubmitLiveEditingJobResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SubmitLiveEditingJobResponseBody) SetRequestId(v string) *SubmitLiveEditingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLiveEditingJobResponseBody) SetVodMediaId(v string) *SubmitLiveEditingJobResponseBody {
	s.VodMediaId = &v
	return s
}

func (s *SubmitLiveEditingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
