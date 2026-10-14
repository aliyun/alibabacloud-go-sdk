// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKBSyncLinkAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetClientId() *string
	SetConfig(v *DescribeKBSyncLinkAttributeResponseBodyConfig) *DescribeKBSyncLinkAttributeResponseBody
	GetConfig() *DescribeKBSyncLinkAttributeResponseBodyConfig
	SetCreationTime(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetCreationTime() *string
	SetDescription(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetDescription() *string
	SetImPlatform(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetImPlatform() *string
	SetLinkId(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetLinkId() *string
	SetLinkName(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetLinkName() *string
	SetRequestId(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetRequestId() *string
	SetSourceDir(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetSourceDir() *string
	SetSyncIntervalMinutes(v int32) *DescribeKBSyncLinkAttributeResponseBody
	GetSyncIntervalMinutes() *int32
	SetSyncStatus(v string) *DescribeKBSyncLinkAttributeResponseBody
	GetSyncStatus() *string
}

type DescribeKBSyncLinkAttributeResponseBody struct {
	// The client ID of the source application.
	//
	// example:
	//
	// cli_a123456
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The synchronization configuration whitelist returned based on the source type. For DINGTALK, UserId can be returned. For FEISHU, FeishuChatMemberStaleTtlSeconds can be returned.
	Config *DescribeKBSyncLinkAttributeResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The creation time of the synchronization link.
	//
	// example:
	//
	// 2026-09-22T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the synchronization link.
	//
	// example:
	//
	// Synchronize Lark knowledge base documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The source channel of the synchronization link. Valid values: OSS, DINGTALK, FEISHU, and SHAREPOINT.
	//
	// example:
	//
	// FEISHU
	ImPlatform *string `json:"ImPlatform,omitempty" xml:"ImPlatform,omitempty"`
	// The unique identifier of the synchronization link.
	//
	// example:
	//
	// pkbl-123456
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// The name of the synchronization link.
	//
	// example:
	//
	// feishu-doc-sync
	LinkName *string `json:"LinkName,omitempty" xml:"LinkName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-1234567890ab
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source directory.
	//
	// example:
	//
	// /documents
	SourceDir *string `json:"SourceDir,omitempty" xml:"SourceDir,omitempty"`
	// The automatic synchronization interval, in minutes.
	//
	// example:
	//
	// 60
	SyncIntervalMinutes *int32 `json:"SyncIntervalMinutes,omitempty" xml:"SyncIntervalMinutes,omitempty"`
	// The status of the synchronization link.
	//
	// example:
	//
	// RUNNING
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s DescribeKBSyncLinkAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinkAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetConfig() *DescribeKBSyncLinkAttributeResponseBodyConfig {
	return s.Config
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetImPlatform() *string {
	return s.ImPlatform
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetLinkId() *string {
	return s.LinkId
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetLinkName() *string {
	return s.LinkName
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetSourceDir() *string {
	return s.SourceDir
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetSyncIntervalMinutes() *int32 {
	return s.SyncIntervalMinutes
}

func (s *DescribeKBSyncLinkAttributeResponseBody) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetClientId(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.ClientId = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetConfig(v *DescribeKBSyncLinkAttributeResponseBodyConfig) *DescribeKBSyncLinkAttributeResponseBody {
	s.Config = v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetCreationTime(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetDescription(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetImPlatform(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.ImPlatform = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetLinkId(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.LinkId = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetLinkName(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.LinkName = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetRequestId(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetSourceDir(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.SourceDir = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetSyncIntervalMinutes(v int32) *DescribeKBSyncLinkAttributeResponseBody {
	s.SyncIntervalMinutes = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) SetSyncStatus(v string) *DescribeKBSyncLinkAttributeResponseBody {
	s.SyncStatus = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKBSyncLinkAttributeResponseBodyConfig struct {
	// The validity period of the successful snapshot of Lark group members for fallback reuse, in seconds. This parameter is returned only for FEISHU synchronization links. A value of -1 indicates that the snapshot never expires. A value of 0 indicates that snapshot fallback reuse is disabled. A positive integer indicates the maximum age of the snapshot allowed for reuse. The valid value configured for the link is returned. If no value is configured, the default value -1 is returned.
	//
	// example:
	//
	// -1
	FeishuChatMemberStaleTtlSeconds *int32 `json:"FeishuChatMemberStaleTtlSeconds,omitempty" xml:"FeishuChatMemberStaleTtlSeconds,omitempty"`
	// The ID of the DingTalk user who performs the operation. This parameter is returned only for DINGTALK synchronization links.
	//
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeKBSyncLinkAttributeResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinkAttributeResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinkAttributeResponseBodyConfig) GetFeishuChatMemberStaleTtlSeconds() *int32 {
	return s.FeishuChatMemberStaleTtlSeconds
}

func (s *DescribeKBSyncLinkAttributeResponseBodyConfig) GetUserId() *string {
	return s.UserId
}

func (s *DescribeKBSyncLinkAttributeResponseBodyConfig) SetFeishuChatMemberStaleTtlSeconds(v int32) *DescribeKBSyncLinkAttributeResponseBodyConfig {
	s.FeishuChatMemberStaleTtlSeconds = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBodyConfig) SetUserId(v string) *DescribeKBSyncLinkAttributeResponseBodyConfig {
	s.UserId = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
