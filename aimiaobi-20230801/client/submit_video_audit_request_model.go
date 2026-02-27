// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v string) *SubmitVideoAuditRequest
	GetExt() *string
	SetFileKey(v string) *SubmitVideoAuditRequest
	GetFileKey() *string
	SetSnapshotInterval(v float64) *SubmitVideoAuditRequest
	GetSnapshotInterval() *float64
	SetUrl(v string) *SubmitVideoAuditRequest
	GetUrl() *string
	SetWorkspaceId(v string) *SubmitVideoAuditRequest
	GetWorkspaceId() *string
}

type SubmitVideoAuditRequest struct {
	// 扩展参数JSON字符串
	//
	// example:
	//
	// {}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// OSS文件Key，与url参数二选一
	//
	// example:
	//
	// video/test.mp4
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// 抽帧间隔时间（秒）
	//
	// example:
	//
	// 1.0
	SnapshotInterval *float64 `json:"SnapshotInterval,omitempty" xml:"SnapshotInterval,omitempty"`
	// 视频URL地址，与fileKey参数二选一
	//
	// example:
	//
	// https://example.com/video.mp4
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitVideoAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAuditRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoAuditRequest) GetExt() *string {
	return s.Ext
}

func (s *SubmitVideoAuditRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitVideoAuditRequest) GetSnapshotInterval() *float64 {
	return s.SnapshotInterval
}

func (s *SubmitVideoAuditRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitVideoAuditRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitVideoAuditRequest) SetExt(v string) *SubmitVideoAuditRequest {
	s.Ext = &v
	return s
}

func (s *SubmitVideoAuditRequest) SetFileKey(v string) *SubmitVideoAuditRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitVideoAuditRequest) SetSnapshotInterval(v float64) *SubmitVideoAuditRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *SubmitVideoAuditRequest) SetUrl(v string) *SubmitVideoAuditRequest {
	s.Url = &v
	return s
}

func (s *SubmitVideoAuditRequest) SetWorkspaceId(v string) *SubmitVideoAuditRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitVideoAuditRequest) Validate() error {
	return dara.Validate(s)
}
