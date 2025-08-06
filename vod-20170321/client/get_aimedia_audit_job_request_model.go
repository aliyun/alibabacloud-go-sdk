// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIMediaAuditJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetAIMediaAuditJobRequest
	GetJobId() *string
}

type GetAIMediaAuditJobRequest struct {
	// The ID of the intelligent review job.
	//
	// This parameter is required.
	//
	// example:
	//
	// bdbc266af6894*****943a70176d92e9
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAIMediaAuditJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobRequest) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetAIMediaAuditJobRequest) SetJobId(v string) *GetAIMediaAuditJobRequest {
	s.JobId = &v
	return s
}

func (s *GetAIMediaAuditJobRequest) Validate() error {
	return dara.Validate(s)
}
