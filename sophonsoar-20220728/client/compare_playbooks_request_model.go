// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComparePlaybooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ComparePlaybooksRequest
	GetLang() *string
	SetNewPlaybookReleaseId(v int32) *ComparePlaybooksRequest
	GetNewPlaybookReleaseId() *int32
	SetOldPlaybookReleaseId(v int32) *ComparePlaybooksRequest
	GetOldPlaybookReleaseId() *int32
	SetPlaybookUuid(v string) *ComparePlaybooksRequest
	GetPlaybookUuid() *string
}

type ComparePlaybooksRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the second version.
	//
	// >  You can call the [DescribePlaybookReleases](~~DescribePlaybookReleases~~) operation to query the IDs of versions. The system automatically generates IDs for new versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// sfdf2395-e814-459f-9662-xxxxx
	NewPlaybookReleaseId *int32 `json:"NewPlaybookReleaseId,omitempty" xml:"NewPlaybookReleaseId,omitempty"`
	// The ID of the first version.
	//
	// >  You can call the [DescribePlaybookReleases](~~DescribePlaybookReleases~~) operation to query the IDs of versions. The system automatically generates IDs for new versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// sflk23423-e814-459f-9662-xxxxx
	OldPlaybookReleaseId *int32 `json:"OldPlaybookReleaseId,omitempty" xml:"OldPlaybookReleaseId,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ComparePlaybooksRequest) String() string {
	return dara.Prettify(s)
}

func (s ComparePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksRequest) GetLang() *string {
	return s.Lang
}

func (s *ComparePlaybooksRequest) GetNewPlaybookReleaseId() *int32 {
	return s.NewPlaybookReleaseId
}

func (s *ComparePlaybooksRequest) GetOldPlaybookReleaseId() *int32 {
	return s.OldPlaybookReleaseId
}

func (s *ComparePlaybooksRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ComparePlaybooksRequest) SetLang(v string) *ComparePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *ComparePlaybooksRequest) SetNewPlaybookReleaseId(v int32) *ComparePlaybooksRequest {
	s.NewPlaybookReleaseId = &v
	return s
}

func (s *ComparePlaybooksRequest) SetOldPlaybookReleaseId(v int32) *ComparePlaybooksRequest {
	s.OldPlaybookReleaseId = &v
	return s
}

func (s *ComparePlaybooksRequest) SetPlaybookUuid(v string) *ComparePlaybooksRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ComparePlaybooksRequest) Validate() error {
	return dara.Validate(s)
}
