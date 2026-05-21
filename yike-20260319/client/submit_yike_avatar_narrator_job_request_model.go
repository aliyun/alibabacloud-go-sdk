// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeAvatarNarratorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobParams(v string) *SubmitYikeAvatarNarratorJobRequest
	GetJobParams() *string
	SetUserData(v string) *SubmitYikeAvatarNarratorJobRequest
	GetUserData() *string
}

type SubmitYikeAvatarNarratorJobRequest struct {
	// This parameter is required.
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// example:
	//
	// {\\"newsKey\\":\\"NEWS_20260420_001\\",\\"key1\\":\\"value1\\", \\"NotifyAddress\\":\\"https://cms.example.com/callback/video-task\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitYikeAvatarNarratorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeAvatarNarratorJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitYikeAvatarNarratorJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitYikeAvatarNarratorJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitYikeAvatarNarratorJobRequest) SetJobParams(v string) *SubmitYikeAvatarNarratorJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitYikeAvatarNarratorJobRequest) SetUserData(v string) *SubmitYikeAvatarNarratorJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitYikeAvatarNarratorJobRequest) Validate() error {
	return dara.Validate(s)
}
