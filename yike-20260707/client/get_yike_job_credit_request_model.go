// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeJobCreditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikeJobCreditRequest
	GetJobId() *string
}

type GetYikeJobCreditRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ag_12412424****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetYikeJobCreditRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeJobCreditRequest) GoString() string {
	return s.String()
}

func (s *GetYikeJobCreditRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeJobCreditRequest) SetJobId(v string) *GetYikeJobCreditRequest {
	s.JobId = &v
	return s
}

func (s *GetYikeJobCreditRequest) Validate() error {
	return dara.Validate(s)
}
