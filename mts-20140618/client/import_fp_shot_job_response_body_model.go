// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportFpShotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ImportFpShotJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *ImportFpShotJobResponseBody
	GetRequestId() *string
}

type ImportFpShotJobResponseBody struct {
	// The ID of the import job. We recommend that you save this ID for subsequent operations.
	//
	// example:
	//
	// c074b118ace44395a02063a5ab94****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportFpShotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportFpShotJobResponseBody) GoString() string {
	return s.String()
}

func (s *ImportFpShotJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ImportFpShotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportFpShotJobResponseBody) SetJobId(v string) *ImportFpShotJobResponseBody {
	s.JobId = &v
	return s
}

func (s *ImportFpShotJobResponseBody) SetRequestId(v string) *ImportFpShotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportFpShotJobResponseBody) Validate() error {
	return dara.Validate(s)
}
