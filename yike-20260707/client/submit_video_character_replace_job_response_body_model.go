// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCharacterReplaceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitVideoCharacterReplaceJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoCharacterReplaceJobResponseBody
	GetRequestId() *string
}

type SubmitVideoCharacterReplaceJobResponseBody struct {
	// The task ID returned by SubmitVideoCharacterReplaceJob.
	//
	// example:
	//
	// ag_83baeea38b534ee1937b8a4acd08****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoCharacterReplaceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCharacterReplaceJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoCharacterReplaceJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoCharacterReplaceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoCharacterReplaceJobResponseBody) SetJobId(v string) *SubmitVideoCharacterReplaceJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoCharacterReplaceJobResponseBody) SetRequestId(v string) *SubmitVideoCharacterReplaceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoCharacterReplaceJobResponseBody) Validate() error {
	return dara.Validate(s)
}
