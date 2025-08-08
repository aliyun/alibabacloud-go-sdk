// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebhookCodeContext interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *WebhookCodeContext
	GetBranch() *string
	SetCommitID(v string) *WebhookCodeContext
	GetCommitID() *string
	SetDescription(v string) *WebhookCodeContext
	GetDescription() *string
	SetEventType(v string) *WebhookCodeContext
	GetEventType() *string
	SetMessage(v string) *WebhookCodeContext
	GetMessage() *string
	SetPrType(v string) *WebhookCodeContext
	GetPrType() *string
	SetRepoUrl(v string) *WebhookCodeContext
	GetRepoUrl() *string
	SetSourceBranch(v string) *WebhookCodeContext
	GetSourceBranch() *string
	SetTag(v string) *WebhookCodeContext
	GetTag() *string
	SetTargetBranch(v string) *WebhookCodeContext
	GetTargetBranch() *string
	SetTitle(v string) *WebhookCodeContext
	GetTitle() *string
}

type WebhookCodeContext struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// b1dd9ba168dfef1cb3a1dd608b6054c771a93959
	CommitID *string `json:"commitID,omitempty" xml:"commitID,omitempty"`
	// example:
	//
	// my PR decscription
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// PUSH
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// example:
	//
	// commit message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// merged
	PrType *string `json:"prType,omitempty" xml:"prType,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/my-namespace/my-repo.git
	RepoUrl *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
	// example:
	//
	// master
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// example:
	//
	// release-0.0.1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// master
	TargetBranch *string `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	// example:
	//
	// # FIX
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s WebhookCodeContext) String() string {
	return dara.Prettify(s)
}

func (s WebhookCodeContext) GoString() string {
	return s.String()
}

func (s *WebhookCodeContext) GetBranch() *string {
	return s.Branch
}

func (s *WebhookCodeContext) GetCommitID() *string {
	return s.CommitID
}

func (s *WebhookCodeContext) GetDescription() *string {
	return s.Description
}

func (s *WebhookCodeContext) GetEventType() *string {
	return s.EventType
}

func (s *WebhookCodeContext) GetMessage() *string {
	return s.Message
}

func (s *WebhookCodeContext) GetPrType() *string {
	return s.PrType
}

func (s *WebhookCodeContext) GetRepoUrl() *string {
	return s.RepoUrl
}

func (s *WebhookCodeContext) GetSourceBranch() *string {
	return s.SourceBranch
}

func (s *WebhookCodeContext) GetTag() *string {
	return s.Tag
}

func (s *WebhookCodeContext) GetTargetBranch() *string {
	return s.TargetBranch
}

func (s *WebhookCodeContext) GetTitle() *string {
	return s.Title
}

func (s *WebhookCodeContext) SetBranch(v string) *WebhookCodeContext {
	s.Branch = &v
	return s
}

func (s *WebhookCodeContext) SetCommitID(v string) *WebhookCodeContext {
	s.CommitID = &v
	return s
}

func (s *WebhookCodeContext) SetDescription(v string) *WebhookCodeContext {
	s.Description = &v
	return s
}

func (s *WebhookCodeContext) SetEventType(v string) *WebhookCodeContext {
	s.EventType = &v
	return s
}

func (s *WebhookCodeContext) SetMessage(v string) *WebhookCodeContext {
	s.Message = &v
	return s
}

func (s *WebhookCodeContext) SetPrType(v string) *WebhookCodeContext {
	s.PrType = &v
	return s
}

func (s *WebhookCodeContext) SetRepoUrl(v string) *WebhookCodeContext {
	s.RepoUrl = &v
	return s
}

func (s *WebhookCodeContext) SetSourceBranch(v string) *WebhookCodeContext {
	s.SourceBranch = &v
	return s
}

func (s *WebhookCodeContext) SetTag(v string) *WebhookCodeContext {
	s.Tag = &v
	return s
}

func (s *WebhookCodeContext) SetTargetBranch(v string) *WebhookCodeContext {
	s.TargetBranch = &v
	return s
}

func (s *WebhookCodeContext) SetTitle(v string) *WebhookCodeContext {
	s.Title = &v
	return s
}

func (s *WebhookCodeContext) Validate() error {
	return dara.Validate(s)
}
