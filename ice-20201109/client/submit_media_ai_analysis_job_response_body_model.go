// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaAiAnalysisJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *SubmitMediaAiAnalysisJobResponseBody
	GetMediaId() *string
	SetRequestId(v string) *SubmitMediaAiAnalysisJobResponseBody
	GetRequestId() *string
}

type SubmitMediaAiAnalysisJobResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaAiAnalysisJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaAiAnalysisJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaAiAnalysisJobResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitMediaAiAnalysisJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaAiAnalysisJobResponseBody) SetMediaId(v string) *SubmitMediaAiAnalysisJobResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitMediaAiAnalysisJobResponseBody) SetRequestId(v string) *SubmitMediaAiAnalysisJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaAiAnalysisJobResponseBody) Validate() error {
	return dara.Validate(s)
}
