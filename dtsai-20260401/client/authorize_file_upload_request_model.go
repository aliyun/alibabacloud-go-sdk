// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *AuthorizeFileUploadRequest
	GetAgentName() *string
	SetBatchSize(v string) *AuthorizeFileUploadRequest
	GetBatchSize() *string
	SetFileFormat(v string) *AuthorizeFileUploadRequest
	GetFileFormat() *string
	SetRegionId(v string) *AuthorizeFileUploadRequest
	GetRegionId() *string
}

type AuthorizeFileUploadRequest struct {
	// The source of the agent or client that initiates the call, such as codex, cursor, or openapi. The maximum length is 32 characters. This parameter is used only for statistics and does not participate in authentication, throttling, quota management, or billing.
	//
	// example:
	//
	// codex
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The number of OSS PostObject direct upload authorizations returned in this request.
	//
	// example:
	//
	// 10
	BatchSize *string `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// The format of the file to be uploaded.
	//
	// example:
	//
	// pdf
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// The region ID, such as cn-beijing.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AuthorizeFileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeFileUploadRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeFileUploadRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *AuthorizeFileUploadRequest) GetBatchSize() *string {
	return s.BatchSize
}

func (s *AuthorizeFileUploadRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *AuthorizeFileUploadRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeFileUploadRequest) SetAgentName(v string) *AuthorizeFileUploadRequest {
	s.AgentName = &v
	return s
}

func (s *AuthorizeFileUploadRequest) SetBatchSize(v string) *AuthorizeFileUploadRequest {
	s.BatchSize = &v
	return s
}

func (s *AuthorizeFileUploadRequest) SetFileFormat(v string) *AuthorizeFileUploadRequest {
	s.FileFormat = &v
	return s
}

func (s *AuthorizeFileUploadRequest) SetRegionId(v string) *AuthorizeFileUploadRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeFileUploadRequest) Validate() error {
	return dara.Validate(s)
}
