// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKBSyncLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *CreateKBSyncLinkResponseBody
	GetClientId() *string
	SetCreationTime(v string) *CreateKBSyncLinkResponseBody
	GetCreationTime() *string
	SetDescription(v string) *CreateKBSyncLinkResponseBody
	GetDescription() *string
	SetImPlatform(v string) *CreateKBSyncLinkResponseBody
	GetImPlatform() *string
	SetLinkId(v string) *CreateKBSyncLinkResponseBody
	GetLinkId() *string
	SetLinkName(v string) *CreateKBSyncLinkResponseBody
	GetLinkName() *string
	SetRequestId(v string) *CreateKBSyncLinkResponseBody
	GetRequestId() *string
	SetSourceDir(v string) *CreateKBSyncLinkResponseBody
	GetSourceDir() *string
	SetSyncIntervalMinutes(v int32) *CreateKBSyncLinkResponseBody
	GetSyncIntervalMinutes() *int32
	SetSyncStatus(v string) *CreateKBSyncLinkResponseBody
	GetSyncStatus() *string
}

type CreateKBSyncLinkResponseBody struct {
	// The client ID.
	//
	// example:
	//
	// cli_xxxxxxbe8
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-14T05:58:42Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the synchronization link.
	//
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The source channel of the synchronization link. Valid values:
	//
	// - FEISHU
	//
	// - SHAREPOINT
	//
	// example:
	//
	// FEISHU
	ImPlatform *string `json:"ImPlatform,omitempty" xml:"ImPlatform,omitempty"`
	// The synchronization link ID.
	//
	// example:
	//
	// pkbl-xxxxxx
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// The name of the synchronization link.
	//
	// example:
	//
	// testName
	LinkName *string `json:"LinkName,omitempty" xml:"LinkName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// CREATING
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s CreateKBSyncLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKBSyncLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKBSyncLinkResponseBody) GetClientId() *string {
	return s.ClientId
}

func (s *CreateKBSyncLinkResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *CreateKBSyncLinkResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateKBSyncLinkResponseBody) GetImPlatform() *string {
	return s.ImPlatform
}

func (s *CreateKBSyncLinkResponseBody) GetLinkId() *string {
	return s.LinkId
}

func (s *CreateKBSyncLinkResponseBody) GetLinkName() *string {
	return s.LinkName
}

func (s *CreateKBSyncLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKBSyncLinkResponseBody) GetSourceDir() *string {
	return s.SourceDir
}

func (s *CreateKBSyncLinkResponseBody) GetSyncIntervalMinutes() *int32 {
	return s.SyncIntervalMinutes
}

func (s *CreateKBSyncLinkResponseBody) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *CreateKBSyncLinkResponseBody) SetClientId(v string) *CreateKBSyncLinkResponseBody {
	s.ClientId = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetCreationTime(v string) *CreateKBSyncLinkResponseBody {
	s.CreationTime = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetDescription(v string) *CreateKBSyncLinkResponseBody {
	s.Description = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetImPlatform(v string) *CreateKBSyncLinkResponseBody {
	s.ImPlatform = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetLinkId(v string) *CreateKBSyncLinkResponseBody {
	s.LinkId = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetLinkName(v string) *CreateKBSyncLinkResponseBody {
	s.LinkName = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetRequestId(v string) *CreateKBSyncLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetSourceDir(v string) *CreateKBSyncLinkResponseBody {
	s.SourceDir = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetSyncIntervalMinutes(v int32) *CreateKBSyncLinkResponseBody {
	s.SyncIntervalMinutes = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) SetSyncStatus(v string) *CreateKBSyncLinkResponseBody {
	s.SyncStatus = &v
	return s
}

func (s *CreateKBSyncLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
