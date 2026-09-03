// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKBSyncLinksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeKBSyncLinksResponseBodyItems) *DescribeKBSyncLinksResponseBody
	GetItems() []*DescribeKBSyncLinksResponseBodyItems
	SetRequestId(v string) *DescribeKBSyncLinksResponseBody
	GetRequestId() *string
}

type DescribeKBSyncLinksResponseBody struct {
	// The list of synchronization links.
	Items []*DescribeKBSyncLinksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EB07CFF0-D8A4-5C76-AED7-D00E26FC2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKBSyncLinksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinksResponseBody) GetItems() []*DescribeKBSyncLinksResponseBodyItems {
	return s.Items
}

func (s *DescribeKBSyncLinksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKBSyncLinksResponseBody) SetItems(v []*DescribeKBSyncLinksResponseBodyItems) *DescribeKBSyncLinksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeKBSyncLinksResponseBody) SetRequestId(v string) *DescribeKBSyncLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKBSyncLinksResponseBodyItems struct {
	// The client identifier.
	//
	// example:
	//
	// cli_xxxxxxbe8
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-11T09:55:19Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the synchronization link.
	//
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The source channel of the synchronization link.
	//
	// example:
	//
	// FEISHU
	ImPlatform *string `json:"ImPlatform,omitempty" xml:"ImPlatform,omitempty"`
	// The synchronization link ID.
	//
	// example:
	//
	// pkbl-xxxxx
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// The link name.
	//
	// example:
	//
	// testName
	LinkName *string `json:"LinkName,omitempty" xml:"LinkName,omitempty"`
	// The source directory address for synchronization.
	//
	// example:
	//
	// https://example.feishu.cn/wiki/space/xxxxxx
	SourceDir *string `json:"SourceDir,omitempty" xml:"SourceDir,omitempty"`
	// The synchronization interval. Unit: minutes.
	//
	// example:
	//
	// 30
	SyncIntervalMinutes *int32 `json:"SyncIntervalMinutes,omitempty" xml:"SyncIntervalMinutes,omitempty"`
	// The synchronization status. Valid values:
	//
	// - CREATING
	//
	// - RUNNING
	//
	// - PAUSED
	//
	// - DELETING
	//
	// example:
	//
	// RUNNING
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s DescribeKBSyncLinksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetImPlatform() *string {
	return s.ImPlatform
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetLinkId() *string {
	return s.LinkId
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetLinkName() *string {
	return s.LinkName
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetSourceDir() *string {
	return s.SourceDir
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetSyncIntervalMinutes() *int32 {
	return s.SyncIntervalMinutes
}

func (s *DescribeKBSyncLinksResponseBodyItems) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetClientId(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.ClientId = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetCreationTime(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetDescription(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetImPlatform(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.ImPlatform = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetLinkId(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.LinkId = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetLinkName(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.LinkName = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetSourceDir(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.SourceDir = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetSyncIntervalMinutes(v int32) *DescribeKBSyncLinksResponseBodyItems {
	s.SyncIntervalMinutes = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) SetSyncStatus(v string) *DescribeKBSyncLinksResponseBodyItems {
	s.SyncStatus = &v
	return s
}

func (s *DescribeKBSyncLinksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
