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
	// This parameter is required.
	//
	// example:
	//
	// sys-ao-smith-1
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// sys-ao-smith-1
	AppParams *string `json:"AppParams,omitempty" xml:"AppParams,omitempty"`
	// example:
	//
	// fd-cReaEcVqQK
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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
