// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIProductionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *QueryIProductionJobRequest
	GetClientToken() *string
	SetJobId(v string) *QueryIProductionJobRequest
	GetJobId() *string
}

type QueryIProductionJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the intelligent production job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryIProductionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobRequest) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *QueryIProductionJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryIProductionJobRequest) SetClientToken(v string) *QueryIProductionJobRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryIProductionJobRequest) SetJobId(v string) *QueryIProductionJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryIProductionJobRequest) Validate() error {
	return dara.Validate(s)
}
