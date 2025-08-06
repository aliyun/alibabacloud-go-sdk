// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveEditingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *SubmitLiveEditingResponseBody
	GetMediaId() *string
	SetProjectId(v string) *SubmitLiveEditingResponseBody
	GetProjectId() *string
	SetRequestId(v string) *SubmitLiveEditingResponseBody
	GetRequestId() *string
}

type SubmitLiveEditingResponseBody struct {
	MediaId   *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitLiveEditingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveEditingResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitLiveEditingResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitLiveEditingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitLiveEditingResponseBody) SetMediaId(v string) *SubmitLiveEditingResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitLiveEditingResponseBody) SetProjectId(v string) *SubmitLiveEditingResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SubmitLiveEditingResponseBody) SetRequestId(v string) *SubmitLiveEditingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLiveEditingResponseBody) Validate() error {
	return dara.Validate(s)
}
