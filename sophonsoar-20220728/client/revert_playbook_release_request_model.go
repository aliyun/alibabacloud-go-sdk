// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertPlaybookReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsPublish(v bool) *RevertPlaybookReleaseRequest
	GetIsPublish() *bool
	SetPlayReleaseId(v int32) *RevertPlaybookReleaseRequest
	GetPlayReleaseId() *int32
	SetPlaybookUuid(v string) *RevertPlaybookReleaseRequest
	GetPlaybookUuid() *string
}

type RevertPlaybookReleaseRequest struct {
	// Specifies whether to directly publish the new playbook after the rollback.
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsPublish *bool `json:"IsPublish,omitempty" xml:"IsPublish,omitempty"`
	// The version of the playbook that you want to publish.
	//
	// >  You can call the [DescribePlaybookReleases](~~DescribePlaybookReleases~~) operation to query the playbook version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3f97b56e-064e-47e7-a309-xxxxxxx
	PlayReleaseId *int32 `json:"PlayReleaseId,omitempty" xml:"PlayReleaseId,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 185295a1-c987-4b64-8796-xxxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s RevertPlaybookReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertPlaybookReleaseRequest) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseRequest) GetIsPublish() *bool {
	return s.IsPublish
}

func (s *RevertPlaybookReleaseRequest) GetPlayReleaseId() *int32 {
	return s.PlayReleaseId
}

func (s *RevertPlaybookReleaseRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *RevertPlaybookReleaseRequest) SetIsPublish(v bool) *RevertPlaybookReleaseRequest {
	s.IsPublish = &v
	return s
}

func (s *RevertPlaybookReleaseRequest) SetPlayReleaseId(v int32) *RevertPlaybookReleaseRequest {
	s.PlayReleaseId = &v
	return s
}

func (s *RevertPlaybookReleaseRequest) SetPlaybookUuid(v string) *RevertPlaybookReleaseRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RevertPlaybookReleaseRequest) Validate() error {
	return dara.Validate(s)
}
