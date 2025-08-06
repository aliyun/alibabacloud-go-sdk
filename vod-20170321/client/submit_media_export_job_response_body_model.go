// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *SubmitMediaExportJobResponseBody
	GetFileURL() *string
	SetJobId(v string) *SubmitMediaExportJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitMediaExportJobResponseBody
	GetRequestId() *string
	SetTotal(v int64) *SubmitMediaExportJobResponseBody
	GetTotal() *int64
}

type SubmitMediaExportJobResponseBody struct {
	FileURL   *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SubmitMediaExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaExportJobResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *SubmitMediaExportJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaExportJobResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SubmitMediaExportJobResponseBody) SetFileURL(v string) *SubmitMediaExportJobResponseBody {
	s.FileURL = &v
	return s
}

func (s *SubmitMediaExportJobResponseBody) SetJobId(v string) *SubmitMediaExportJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMediaExportJobResponseBody) SetRequestId(v string) *SubmitMediaExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaExportJobResponseBody) SetTotal(v int64) *SubmitMediaExportJobResponseBody {
	s.Total = &v
	return s
}

func (s *SubmitMediaExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
