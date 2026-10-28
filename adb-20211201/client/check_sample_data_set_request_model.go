// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSampleDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckSampleDataSetRequest
	GetDBClusterId() *string
}

type CheckSampleDataSetRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9r8f67h4cqz41u
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s CheckSampleDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSampleDataSetRequest) GoString() string {
	return s.String()
}

func (s *CheckSampleDataSetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckSampleDataSetRequest) SetDBClusterId(v string) *CheckSampleDataSetRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckSampleDataSetRequest) Validate() error {
	return dara.Validate(s)
}
