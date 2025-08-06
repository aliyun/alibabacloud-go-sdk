// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDNAInitializationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitDNAInitializationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitDNAInitializationJobResponseBody
	GetRequestId() *string
}

type SubmitDNAInitializationJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDNAInitializationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAInitializationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDNAInitializationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitDNAInitializationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDNAInitializationJobResponseBody) SetJobId(v string) *SubmitDNAInitializationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitDNAInitializationJobResponseBody) SetRequestId(v string) *SubmitDNAInitializationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDNAInitializationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
