// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableFromAuthorizedOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOssBucket(v string) *UpdateTableFromAuthorizedOssRequest
	GetOssBucket() *string
	SetOssKey(v string) *UpdateTableFromAuthorizedOssRequest
	GetOssKey() *string
	SetOssRegionId(v string) *UpdateTableFromAuthorizedOssRequest
	GetOssRegionId() *string
	SetUpdateMode(v string) *UpdateTableFromAuthorizedOssRequest
	GetUpdateMode() *string
}

type UpdateTableFromAuthorizedOssRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yinghuo-ai
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a0deedbce4a8162b8d66c63ace28330c
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OVERWRITE
	UpdateMode *string `json:"UpdateMode,omitempty" xml:"UpdateMode,omitempty"`
}

func (s UpdateTableFromAuthorizedOssRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableFromAuthorizedOssRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableFromAuthorizedOssRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *UpdateTableFromAuthorizedOssRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *UpdateTableFromAuthorizedOssRequest) GetOssRegionId() *string {
	return s.OssRegionId
}

func (s *UpdateTableFromAuthorizedOssRequest) GetUpdateMode() *string {
	return s.UpdateMode
}

func (s *UpdateTableFromAuthorizedOssRequest) SetOssBucket(v string) *UpdateTableFromAuthorizedOssRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssRequest) SetOssKey(v string) *UpdateTableFromAuthorizedOssRequest {
	s.OssKey = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssRequest) SetOssRegionId(v string) *UpdateTableFromAuthorizedOssRequest {
	s.OssRegionId = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssRequest) SetUpdateMode(v string) *UpdateTableFromAuthorizedOssRequest {
	s.UpdateMode = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssRequest) Validate() error {
	return dara.Validate(s)
}
