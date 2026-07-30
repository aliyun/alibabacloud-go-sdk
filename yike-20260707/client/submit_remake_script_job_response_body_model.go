// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRemakeScriptJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitRemakeScriptJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *SubmitRemakeScriptJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitRemakeScriptJobResponseBody
	GetRequestId() *string
}

type SubmitRemakeScriptJobResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 68ca759e798b40b4903b255********
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitRemakeScriptJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitRemakeScriptJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitRemakeScriptJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitRemakeScriptJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitRemakeScriptJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitRemakeScriptJobResponseBody) SetErrorCode(v string) *SubmitRemakeScriptJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitRemakeScriptJobResponseBody) SetJobId(v string) *SubmitRemakeScriptJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitRemakeScriptJobResponseBody) SetRequestId(v string) *SubmitRemakeScriptJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitRemakeScriptJobResponseBody) Validate() error {
	return dara.Validate(s)
}
