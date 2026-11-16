// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemakeScriptJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetRemakeScriptJobRequest
	GetJobId() *string
}

type GetRemakeScriptJobRequest struct {
	// The task ID.
	//
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetRemakeScriptJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRemakeScriptJobRequest) GoString() string {
	return s.String()
}

func (s *GetRemakeScriptJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetRemakeScriptJobRequest) SetJobId(v string) *GetRemakeScriptJobRequest {
	s.JobId = &v
	return s
}

func (s *GetRemakeScriptJobRequest) Validate() error {
	return dara.Validate(s)
}
