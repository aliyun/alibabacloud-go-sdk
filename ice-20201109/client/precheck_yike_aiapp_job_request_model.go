// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckYikeAIAppJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PrecheckYikeAIAppJobRequest
	GetAppId() *string
	SetAppParams(v string) *PrecheckYikeAIAppJobRequest
	GetAppParams() *string
}

type PrecheckYikeAIAppJobRequest struct {
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {\\"LoadImage.1.TargetImage\\":\\"794da8a01b8c71f1b973e6e7c7586301\\"}
	AppParams *string `json:"AppParams,omitempty" xml:"AppParams,omitempty"`
}

func (s PrecheckYikeAIAppJobRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckYikeAIAppJobRequest) GoString() string {
	return s.String()
}

func (s *PrecheckYikeAIAppJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *PrecheckYikeAIAppJobRequest) GetAppParams() *string {
	return s.AppParams
}

func (s *PrecheckYikeAIAppJobRequest) SetAppId(v string) *PrecheckYikeAIAppJobRequest {
	s.AppId = &v
	return s
}

func (s *PrecheckYikeAIAppJobRequest) SetAppParams(v string) *PrecheckYikeAIAppJobRequest {
	s.AppParams = &v
	return s
}

func (s *PrecheckYikeAIAppJobRequest) Validate() error {
	return dara.Validate(s)
}
