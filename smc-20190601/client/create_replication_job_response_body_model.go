// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateReplicationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateReplicationJobResponseBody
	GetRequestId() *string
}

type CreateReplicationJobResponseBody struct {
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
	// C8B26B44-0189-443E-9816-D951F59623A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateReplicationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateReplicationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReplicationJobResponseBody) SetJobId(v string) *CreateReplicationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateReplicationJobResponseBody) SetRequestId(v string) *CreateReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReplicationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
