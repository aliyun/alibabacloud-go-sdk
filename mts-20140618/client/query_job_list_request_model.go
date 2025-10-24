// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *QueryJobListRequest
	GetJobIds() *string
	SetOwnerAccount(v string) *QueryJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryJobListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryJobListRequest
	GetResourceOwnerId() *int64
}

type QueryJobListRequest struct {
	// The IDs of transcoding jobs. Separate multiple IDs with commas (,). You can query a maximum of 10 transcoding jobs at a time. You can log on to the [ApsaraVideo Media Processing (MPS) console](https://mps.console.aliyun.com/overview) and click **Tasks*	- in the left-side navigation pane to obtain job IDs. Alternatively, you can obtain job IDs from the response to the [SubmitJobs](https://help.aliyun.com/document_detail/29226.html) operation.
	//
	// >  If you do not set the JobIds parameter, the `InvalidParameter` error code is returned.
	//
	// example:
	//
	// bb558c1cc25b45309aab5be44d19****,d1ce4d3efcb549419193f50f1fcd****
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryJobListRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *QueryJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryJobListRequest) SetJobIds(v string) *QueryJobListRequest {
	s.JobIds = &v
	return s
}

func (s *QueryJobListRequest) SetOwnerAccount(v string) *QueryJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryJobListRequest) SetOwnerId(v int64) *QueryJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryJobListRequest) SetResourceOwnerAccount(v string) *QueryJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryJobListRequest) SetResourceOwnerId(v int64) *QueryJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryJobListRequest) Validate() error {
	return dara.Validate(s)
}
