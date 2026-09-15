// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelGenerationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelGenerationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CancelGenerationJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *CancelGenerationJobResponseBody
	GetStatus() *string
}

type CancelGenerationJobResponseBody struct {
	// example:
	//
	// ag_77b5f78b***
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Canceled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelGenerationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelGenerationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelGenerationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CancelGenerationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelGenerationJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelGenerationJobResponseBody) SetJobId(v string) *CancelGenerationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CancelGenerationJobResponseBody) SetRequestId(v string) *CancelGenerationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelGenerationJobResponseBody) SetStatus(v string) *CancelGenerationJobResponseBody {
	s.Status = &v
	return s
}

func (s *CancelGenerationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
