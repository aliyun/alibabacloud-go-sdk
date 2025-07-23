// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *CreateDeploymentJobResponseBody
	GetJobId() *int64
	SetRequestId(v string) *CreateDeploymentJobResponseBody
	GetRequestId() *string
}

type CreateDeploymentJobResponseBody struct {
	// The ID of the deployment task.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeploymentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentJobResponseBody) GetJobId() *int64 {
	return s.JobId
}

func (s *CreateDeploymentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeploymentJobResponseBody) SetJobId(v int64) *CreateDeploymentJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDeploymentJobResponseBody) SetRequestId(v string) *CreateDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentJobResponseBody) Validate() error {
	return dara.Validate(s)
}
