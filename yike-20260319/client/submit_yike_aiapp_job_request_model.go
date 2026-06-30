// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeAIAppJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SubmitYikeAIAppJobRequest
	GetAppId() *string
	SetAppParams(v string) *SubmitYikeAIAppJobRequest
	GetAppParams() *string
	SetFolderId(v string) *SubmitYikeAIAppJobRequest
	GetFolderId() *string
	SetProductionId(v string) *SubmitYikeAIAppJobRequest
	GetProductionId() *string
}

type SubmitYikeAIAppJobRequest struct {
	// The AI application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_test
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The runtime parameters of the AI application. The value is a string produced by JSON serialization.
	//
	// example:
	//
	// {"testKey":"testValue"}
	AppParams *string `json:"AppParams,omitempty" xml:"AppParams,omitempty"`
	// The folder ID. This parameter is used to save results to the specified folder.
	//
	// example:
	//
	// fd-cReaEcVqQK
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The project ID. This parameter is used to save results to the specified project.
	//
	// example:
	//
	// ProductionId
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
}

func (s SubmitYikeAIAppJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeAIAppJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitYikeAIAppJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *SubmitYikeAIAppJobRequest) GetAppParams() *string {
	return s.AppParams
}

func (s *SubmitYikeAIAppJobRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitYikeAIAppJobRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *SubmitYikeAIAppJobRequest) SetAppId(v string) *SubmitYikeAIAppJobRequest {
	s.AppId = &v
	return s
}

func (s *SubmitYikeAIAppJobRequest) SetAppParams(v string) *SubmitYikeAIAppJobRequest {
	s.AppParams = &v
	return s
}

func (s *SubmitYikeAIAppJobRequest) SetFolderId(v string) *SubmitYikeAIAppJobRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitYikeAIAppJobRequest) SetProductionId(v string) *SubmitYikeAIAppJobRequest {
	s.ProductionId = &v
	return s
}

func (s *SubmitYikeAIAppJobRequest) Validate() error {
	return dara.Validate(s)
}
