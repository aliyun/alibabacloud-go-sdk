// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetJobId(v string) *CreateDatasetJobResponseBody
	GetDatasetJobId() *string
	SetRequestId(v string) *CreateDatasetJobResponseBody
	GetRequestId() *string
}

type CreateDatasetJobResponseBody struct {
	// The ID of the dataset job.
	//
	// example:
	//
	// dsjob-9jx1******uj9e
	DatasetJobId *string `json:"DatasetJobId,omitempty" xml:"DatasetJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 99341606-****-0757724D97EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetJobResponseBody) GetDatasetJobId() *string {
	return s.DatasetJobId
}

func (s *CreateDatasetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetJobResponseBody) SetDatasetJobId(v string) *CreateDatasetJobResponseBody {
	s.DatasetJobId = &v
	return s
}

func (s *CreateDatasetJobResponseBody) SetRequestId(v string) *CreateDatasetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetJobResponseBody) Validate() error {
	return dara.Validate(s)
}
