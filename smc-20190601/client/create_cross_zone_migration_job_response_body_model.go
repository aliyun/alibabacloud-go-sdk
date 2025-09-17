// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrossZoneMigrationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateCrossZoneMigrationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateCrossZoneMigrationJobResponseBody
	GetRequestId() *string
}

type CreateCrossZoneMigrationJobResponseBody struct {
	// The ID of the migration job.
	//
	// example:
	//
	// j-bp17bclvg344jlyt****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A9DBD2F8-DE5A-5844-BA6F-957A996CBD78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCrossZoneMigrationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossZoneMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCrossZoneMigrationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateCrossZoneMigrationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCrossZoneMigrationJobResponseBody) SetJobId(v string) *CreateCrossZoneMigrationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobResponseBody) SetRequestId(v string) *CreateCrossZoneMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
