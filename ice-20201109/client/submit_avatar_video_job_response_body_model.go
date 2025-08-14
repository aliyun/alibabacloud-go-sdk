// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAvatarVideoJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAvatarVideoJobResponseBody
	GetJobId() *string
	SetMediaId(v string) *SubmitAvatarVideoJobResponseBody
	GetMediaId() *string
	SetRequestId(v string) *SubmitAvatarVideoJobResponseBody
	GetRequestId() *string
}

type SubmitAvatarVideoJobResponseBody struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******70dcc471edaf00e6f6f4******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAvatarVideoJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAvatarVideoJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAvatarVideoJobResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAvatarVideoJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAvatarVideoJobResponseBody) SetJobId(v string) *SubmitAvatarVideoJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAvatarVideoJobResponseBody) SetMediaId(v string) *SubmitAvatarVideoJobResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitAvatarVideoJobResponseBody) SetRequestId(v string) *SubmitAvatarVideoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAvatarVideoJobResponseBody) Validate() error {
	return dara.Validate(s)
}
