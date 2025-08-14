// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelIProductionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelIProductionJobRequest
	GetClientToken() *string
	SetJobId(v string) *CancelIProductionJobRequest
	GetJobId() *string
}

type CancelIProductionJobRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelIProductionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelIProductionJobRequest) GoString() string {
	return s.String()
}

func (s *CancelIProductionJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelIProductionJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelIProductionJobRequest) SetClientToken(v string) *CancelIProductionJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelIProductionJobRequest) SetJobId(v string) *CancelIProductionJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelIProductionJobRequest) Validate() error {
	return dara.Validate(s)
}
