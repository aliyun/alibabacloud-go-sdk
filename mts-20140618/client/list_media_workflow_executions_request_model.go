// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaWorkflowExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputFileURL(v string) *ListMediaWorkflowExecutionsRequest
	GetInputFileURL() *string
	SetMaximumPageSize(v int64) *ListMediaWorkflowExecutionsRequest
	GetMaximumPageSize() *int64
	SetMediaWorkflowId(v string) *ListMediaWorkflowExecutionsRequest
	GetMediaWorkflowId() *string
	SetMediaWorkflowName(v string) *ListMediaWorkflowExecutionsRequest
	GetMediaWorkflowName() *string
	SetNextPageToken(v string) *ListMediaWorkflowExecutionsRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *ListMediaWorkflowExecutionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListMediaWorkflowExecutionsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListMediaWorkflowExecutionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMediaWorkflowExecutionsRequest
	GetResourceOwnerId() *int64
}

type ListMediaWorkflowExecutionsRequest struct {
	// The Object Storage Service (OSS) URL of the input file of the media workflow. The URL complies with RFC 3986 and is encoded in UTF-8, with reserved characters being percent-encoded. For more information, see [URL encoding](https://help.aliyun.com/document_detail/423796.html).
	//
	// example:
	//
	// http://example-****.cn-hangzhou.aliyuncs.com/test****.flv
	InputFileURL *string `json:"InputFileURL,omitempty" xml:"InputFileURL,omitempty"`
	// The maximum number of media workflow execution instances to return. Valid values: `[1,100]`. Default value: **10**.
	//
	// example:
	//
	// 1
	MaximumPageSize *int64 `json:"MaximumPageSize,omitempty" xml:"MaximumPageSize,omitempty"`
	// The ID of the media workflow whose execution instances you want to query. To obtain the workflow ID, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Workflow Settings**.
	//
	// example:
	//
	// 43b7335a4b1d4fe883670036affb****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow. To obtain the workflow name, you can log on to the **MPS console*	- and choose **Workflows*	- > **Workflow Settings**.
	//
	// example:
	//
	// example-mediaworkflow-****
	MediaWorkflowName *string `json:"MediaWorkflowName,omitempty" xml:"MediaWorkflowName,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. The value is a UUID that contains 32 characters. When you request the first page of query results, leave the NextPageToken parameter empty. When you request more query results, specify the value of the NextPageToken parameter returned in the query results on the previous page.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	NextPageToken        *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListMediaWorkflowExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsRequest) GetInputFileURL() *string {
	return s.InputFileURL
}

func (s *ListMediaWorkflowExecutionsRequest) GetMaximumPageSize() *int64 {
	return s.MaximumPageSize
}

func (s *ListMediaWorkflowExecutionsRequest) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *ListMediaWorkflowExecutionsRequest) GetMediaWorkflowName() *string {
	return s.MediaWorkflowName
}

func (s *ListMediaWorkflowExecutionsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListMediaWorkflowExecutionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListMediaWorkflowExecutionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMediaWorkflowExecutionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMediaWorkflowExecutionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMediaWorkflowExecutionsRequest) SetInputFileURL(v string) *ListMediaWorkflowExecutionsRequest {
	s.InputFileURL = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetMaximumPageSize(v int64) *ListMediaWorkflowExecutionsRequest {
	s.MaximumPageSize = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetMediaWorkflowId(v string) *ListMediaWorkflowExecutionsRequest {
	s.MediaWorkflowId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetMediaWorkflowName(v string) *ListMediaWorkflowExecutionsRequest {
	s.MediaWorkflowName = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetNextPageToken(v string) *ListMediaWorkflowExecutionsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetOwnerAccount(v string) *ListMediaWorkflowExecutionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetOwnerId(v int64) *ListMediaWorkflowExecutionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetResourceOwnerAccount(v string) *ListMediaWorkflowExecutionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) SetResourceOwnerId(v int64) *ListMediaWorkflowExecutionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
