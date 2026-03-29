// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDiagnosisJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiagnosisJobIds(v []*string) *CreateDataDiagnosisJobsResponseBody
	GetDataDiagnosisJobIds() []*string
	SetRequestId(v string) *CreateDataDiagnosisJobsResponseBody
	GetRequestId() *string
}

type CreateDataDiagnosisJobsResponseBody struct {
	DataDiagnosisJobIds []*string `json:"DataDiagnosisJobIds,omitempty" xml:"DataDiagnosisJobIds,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataDiagnosisJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDiagnosisJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataDiagnosisJobsResponseBody) GetDataDiagnosisJobIds() []*string {
	return s.DataDiagnosisJobIds
}

func (s *CreateDataDiagnosisJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataDiagnosisJobsResponseBody) SetDataDiagnosisJobIds(v []*string) *CreateDataDiagnosisJobsResponseBody {
	s.DataDiagnosisJobIds = v
	return s
}

func (s *CreateDataDiagnosisJobsResponseBody) SetRequestId(v string) *CreateDataDiagnosisJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataDiagnosisJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
