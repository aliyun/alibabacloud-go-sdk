// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoTranslationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitVideoTranslationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoTranslationJobResponseBody
	GetRequestId() *string
}

type SubmitVideoTranslationJobResponseBody struct {
	// `data.JobId`
	//
	// example:
	//
	// vtj_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// `requestId`
	//
	// example:
	//
	// request-id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoTranslationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoTranslationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoTranslationJobResponseBody) SetJobId(v string) *SubmitVideoTranslationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoTranslationJobResponseBody) SetRequestId(v string) *SubmitVideoTranslationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoTranslationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
