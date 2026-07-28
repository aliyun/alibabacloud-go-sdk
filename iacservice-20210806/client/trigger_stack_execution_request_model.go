// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerStackExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *TriggerStackExecutionRequest
	GetAction() *string
	SetChangedFolders(v []*string) *TriggerStackExecutionRequest
	GetChangedFolders() []*string
	SetClientToken(v string) *TriggerStackExecutionRequest
	GetClientToken() *string
	SetCodePackagePath(v string) *TriggerStackExecutionRequest
	GetCodePackagePath() *string
	SetCodeVersionId(v string) *TriggerStackExecutionRequest
	GetCodeVersionId() *string
}

type TriggerStackExecutionRequest struct {
	// Operation to execute
	//
	// - terraform plan
	//
	// - terraform apply
	//
	// - state detect
	//
	// This parameter is required.
	//
	// example:
	//
	// terraform plan
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// Collection of changed files. You can specify only the folders containing changed files.
	//
	// This parameter is required.
	ChangedFolders []*string `json:"changedFolders,omitempty" xml:"changedFolders,omitempty" type:"Repeated"`
	// Idempotent token. Format: [0-9a-zA-Z-]{1,64}. We recommend using a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Stack code path. Currently, only Stacks created through IaC templates are supported.
	//
	// Parameter format: iacservice::{moduleId}
	//
	// This parameter is required.
	//
	// example:
	//
	// iacservice::mod-xxx
	CodePackagePath *string `json:"codePackagePath,omitempty" xml:"codePackagePath,omitempty"`
	// Code version of the Stack to trigger
	//
	// example:
	//
	// v1
	CodeVersionId *string `json:"codeVersionId,omitempty" xml:"codeVersionId,omitempty"`
}

func (s TriggerStackExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerStackExecutionRequest) GoString() string {
	return s.String()
}

func (s *TriggerStackExecutionRequest) GetAction() *string {
	return s.Action
}

func (s *TriggerStackExecutionRequest) GetChangedFolders() []*string {
	return s.ChangedFolders
}

func (s *TriggerStackExecutionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TriggerStackExecutionRequest) GetCodePackagePath() *string {
	return s.CodePackagePath
}

func (s *TriggerStackExecutionRequest) GetCodeVersionId() *string {
	return s.CodeVersionId
}

func (s *TriggerStackExecutionRequest) SetAction(v string) *TriggerStackExecutionRequest {
	s.Action = &v
	return s
}

func (s *TriggerStackExecutionRequest) SetChangedFolders(v []*string) *TriggerStackExecutionRequest {
	s.ChangedFolders = v
	return s
}

func (s *TriggerStackExecutionRequest) SetClientToken(v string) *TriggerStackExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *TriggerStackExecutionRequest) SetCodePackagePath(v string) *TriggerStackExecutionRequest {
	s.CodePackagePath = &v
	return s
}

func (s *TriggerStackExecutionRequest) SetCodeVersionId(v string) *TriggerStackExecutionRequest {
	s.CodeVersionId = &v
	return s
}

func (s *TriggerStackExecutionRequest) Validate() error {
	return dara.Validate(s)
}
