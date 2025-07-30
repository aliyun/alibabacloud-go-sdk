// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobSanityCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetJobSanityCheckResultResponseBody
	GetJobId() *string
	SetRequestID(v string) *GetJobSanityCheckResultResponseBody
	GetRequestID() *string
	SetSanityCheckResult(v []*SanityCheckResultItem) *GetJobSanityCheckResultResponseBody
	GetSanityCheckResult() []*SanityCheckResultItem
}

type GetJobSanityCheckResultResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-xxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B3789344-F1xxxBE-5xx2-A04D-xxxxx
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// The job sanity check result.
	SanityCheckResult []*SanityCheckResultItem `json:"SanityCheckResult,omitempty" xml:"SanityCheckResult,omitempty" type:"Repeated"`
}

func (s GetJobSanityCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobSanityCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobSanityCheckResultResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetJobSanityCheckResultResponseBody) GetRequestID() *string {
	return s.RequestID
}

func (s *GetJobSanityCheckResultResponseBody) GetSanityCheckResult() []*SanityCheckResultItem {
	return s.SanityCheckResult
}

func (s *GetJobSanityCheckResultResponseBody) SetJobId(v string) *GetJobSanityCheckResultResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobSanityCheckResultResponseBody) SetRequestID(v string) *GetJobSanityCheckResultResponseBody {
	s.RequestID = &v
	return s
}

func (s *GetJobSanityCheckResultResponseBody) SetSanityCheckResult(v []*SanityCheckResultItem) *GetJobSanityCheckResultResponseBody {
	s.SanityCheckResult = v
	return s
}

func (s *GetJobSanityCheckResultResponseBody) Validate() error {
	return dara.Validate(s)
}
