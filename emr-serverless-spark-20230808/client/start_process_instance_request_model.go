// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartProcessInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *StartProcessInstanceRequest
	GetAction() *string
	SetComments(v string) *StartProcessInstanceRequest
	GetComments() *string
	SetEmail(v string) *StartProcessInstanceRequest
	GetEmail() *string
	SetInterval(v string) *StartProcessInstanceRequest
	GetInterval() *string
	SetIsProd(v bool) *StartProcessInstanceRequest
	GetIsProd() *bool
	SetProcessDefinitionCode(v int64) *StartProcessInstanceRequest
	GetProcessDefinitionCode() *int64
	SetProductNamespace(v string) *StartProcessInstanceRequest
	GetProductNamespace() *string
	SetRegionId(v string) *StartProcessInstanceRequest
	GetRegionId() *string
	SetRuntimeQueue(v string) *StartProcessInstanceRequest
	GetRuntimeQueue() *string
	SetVersionHashCode(v string) *StartProcessInstanceRequest
	GetVersionHashCode() *string
	SetVersionNumber(v int32) *StartProcessInstanceRequest
	GetVersionNumber() *int32
}

type StartProcessInstanceRequest struct {
	Action   *string `json:"action,omitempty" xml:"action,omitempty"`
	Comments *string `json:"comments,omitempty" xml:"comments,omitempty"`
	Email    *string `json:"email,omitempty" xml:"email,omitempty"`
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
	// Specifies whether to run the workflow in the production environment.
	//
	// example:
	//
	// false
	IsProd *bool `json:"isProd,omitempty" xml:"isProd,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***********
	ProcessDefinitionCode *int64 `json:"processDefinitionCode,omitempty" xml:"processDefinitionCode,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The queue on which the workflow runs.
	//
	// example:
	//
	// root_queue
	RuntimeQueue *string `json:"runtimeQueue,omitempty" xml:"runtimeQueue,omitempty"`
	// The hash code of the version.
	//
	// example:
	//
	// dh*********
	VersionHashCode *string `json:"versionHashCode,omitempty" xml:"versionHashCode,omitempty"`
	// The version number of the workflow.
	//
	// example:
	//
	// 1
	VersionNumber *int32 `json:"versionNumber,omitempty" xml:"versionNumber,omitempty"`
}

func (s StartProcessInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequest) GetAction() *string {
	return s.Action
}

func (s *StartProcessInstanceRequest) GetComments() *string {
	return s.Comments
}

func (s *StartProcessInstanceRequest) GetEmail() *string {
	return s.Email
}

func (s *StartProcessInstanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *StartProcessInstanceRequest) GetIsProd() *bool {
	return s.IsProd
}

func (s *StartProcessInstanceRequest) GetProcessDefinitionCode() *int64 {
	return s.ProcessDefinitionCode
}

func (s *StartProcessInstanceRequest) GetProductNamespace() *string {
	return s.ProductNamespace
}

func (s *StartProcessInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartProcessInstanceRequest) GetRuntimeQueue() *string {
	return s.RuntimeQueue
}

func (s *StartProcessInstanceRequest) GetVersionHashCode() *string {
	return s.VersionHashCode
}

func (s *StartProcessInstanceRequest) GetVersionNumber() *int32 {
	return s.VersionNumber
}

func (s *StartProcessInstanceRequest) SetAction(v string) *StartProcessInstanceRequest {
	s.Action = &v
	return s
}

func (s *StartProcessInstanceRequest) SetComments(v string) *StartProcessInstanceRequest {
	s.Comments = &v
	return s
}

func (s *StartProcessInstanceRequest) SetEmail(v string) *StartProcessInstanceRequest {
	s.Email = &v
	return s
}

func (s *StartProcessInstanceRequest) SetInterval(v string) *StartProcessInstanceRequest {
	s.Interval = &v
	return s
}

func (s *StartProcessInstanceRequest) SetIsProd(v bool) *StartProcessInstanceRequest {
	s.IsProd = &v
	return s
}

func (s *StartProcessInstanceRequest) SetProcessDefinitionCode(v int64) *StartProcessInstanceRequest {
	s.ProcessDefinitionCode = &v
	return s
}

func (s *StartProcessInstanceRequest) SetProductNamespace(v string) *StartProcessInstanceRequest {
	s.ProductNamespace = &v
	return s
}

func (s *StartProcessInstanceRequest) SetRegionId(v string) *StartProcessInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetRuntimeQueue(v string) *StartProcessInstanceRequest {
	s.RuntimeQueue = &v
	return s
}

func (s *StartProcessInstanceRequest) SetVersionHashCode(v string) *StartProcessInstanceRequest {
	s.VersionHashCode = &v
	return s
}

func (s *StartProcessInstanceRequest) SetVersionNumber(v int32) *StartProcessInstanceRequest {
	s.VersionNumber = &v
	return s
}

func (s *StartProcessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
