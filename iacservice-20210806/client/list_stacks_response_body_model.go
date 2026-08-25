// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListStacksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListStacksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListStacksResponseBody
	GetRequestId() *string
	SetStacks(v []*ListStacksResponseBodyStacks) *ListStacksResponseBody
	GetStacks() []*ListStacksResponseBodyStacks
	SetTotalCount(v int32) *ListStacksResponseBody
	GetTotalCount() *int32
}

type ListStacksResponseBody struct {
	// The maximum number of results returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. This parameter is empty if no more pages are available.
	//
	// example:
	//
	// LC4NJL3Ru2bIiRdnbADPQp4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E49127A-BB65-5CCD-AB93-0EC0A43E5446
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of stacks.
	Stacks []*ListStacksResponseBodyStacks `json:"stacks,omitempty" xml:"stacks,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListStacksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStacksResponseBody) GetStacks() []*ListStacksResponseBodyStacks {
	return s.Stacks
}

func (s *ListStacksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStacksResponseBody) SetMaxResults(v int32) *ListStacksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListStacksResponseBody) SetNextToken(v string) *ListStacksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStacksResponseBody) SetRequestId(v string) *ListStacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStacksResponseBody) SetStacks(v []*ListStacksResponseBodyStacks) *ListStacksResponseBody {
	s.Stacks = v
	return s
}

func (s *ListStacksResponseBody) SetTotalCount(v int32) *ListStacksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStacksResponseBody) Validate() error {
	if s.Stacks != nil {
		for _, item := range s.Stacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStacksResponseBodyStacks struct {
	// The creation time in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-05-07T02:21:28Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the stack.
	//
	// example:
	//
	// description of stack
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The stack name.
	//
	// example:
	//
	// stack-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The creation source. Valid values:
	//
	// - OSS: a template stored in Object Storage Service (OSS).
	//
	// - IAC_SERVICE_MODULE: a template created in the automation service console.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path value of the configuration source. The value cannot exceed 1000 characters.
	//
	// - If the source is OSS, the value is in the format oss::<file link> and must be a zip file. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	//
	// - If the source is IAC_SERVICE_MODULE, the value is a template ID. Example: mod-xxxxx
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The description of the stack.
	//
	// example:
	//
	// description of stack
	StackDescription *string `json:"stackDescription,omitempty" xml:"stackDescription,omitempty"`
	// The stack ID, which is the unique identifier generated after the stack is created.
	//
	// example:
	//
	// stack-as1d4vld898ppnqxxxxxx
	StackId *string `json:"stackId,omitempty" xml:"stackId,omitempty"`
	// The stack name (deprecated). Use name instead.
	//
	// example:
	//
	// stack-test
	StackName *string `json:"stackName,omitempty" xml:"stackName,omitempty"`
	// The stack status.
	//
	// | Name | Description |
	//
	// |------|------|
	//
	// | Creating | Being created. |
	//
	// | Created | Created. |
	//
	// | Waiting | Waiting for deployment. |
	//
	// | Deploying | Being deployed. |
	//
	// | Deployed | Deployed. |
	//
	// | Errored | Deployment failed. |
	//
	// | Deleting | Being deleted. |
	//
	// | Deleted | Deleted. |
	//
	// | DeleteFailed | Deletion failed. |
	//
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListStacksResponseBodyStacks) String() string {
	return dara.Prettify(s)
}

func (s ListStacksResponseBodyStacks) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStacksResponseBodyStacks) GetDescription() *string {
	return s.Description
}

func (s *ListStacksResponseBodyStacks) GetName() *string {
	return s.Name
}

func (s *ListStacksResponseBodyStacks) GetSource() *string {
	return s.Source
}

func (s *ListStacksResponseBodyStacks) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListStacksResponseBodyStacks) GetStackDescription() *string {
	return s.StackDescription
}

func (s *ListStacksResponseBodyStacks) GetStackId() *string {
	return s.StackId
}

func (s *ListStacksResponseBodyStacks) GetStackName() *string {
	return s.StackName
}

func (s *ListStacksResponseBodyStacks) GetStatus() *string {
	return s.Status
}

func (s *ListStacksResponseBodyStacks) SetCreateTime(v string) *ListStacksResponseBodyStacks {
	s.CreateTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDescription(v string) *ListStacksResponseBodyStacks {
	s.Description = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetName(v string) *ListStacksResponseBodyStacks {
	s.Name = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetSource(v string) *ListStacksResponseBodyStacks {
	s.Source = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetSourcePath(v string) *ListStacksResponseBodyStacks {
	s.SourcePath = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackDescription(v string) *ListStacksResponseBodyStacks {
	s.StackDescription = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackId(v string) *ListStacksResponseBodyStacks {
	s.StackId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackName(v string) *ListStacksResponseBodyStacks {
	s.StackName = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStatus(v string) *ListStacksResponseBodyStacks {
	s.Status = &v
	return s
}

func (s *ListStacksResponseBodyStacks) Validate() error {
	return dara.Validate(s)
}
