// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartSysAvatarModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *ListSmartSysAvatarModelsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListSmartSysAvatarModelsRequest
	GetPageSize() *int64
	SetSdkVersion(v string) *ListSmartSysAvatarModelsRequest
	GetSdkVersion() *string
}

type ListSmartSysAvatarModelsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s ListSmartSysAvatarModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmartSysAvatarModelsRequest) GoString() string {
	return s.String()
}

func (s *ListSmartSysAvatarModelsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListSmartSysAvatarModelsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSmartSysAvatarModelsRequest) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *ListSmartSysAvatarModelsRequest) SetPageNo(v int64) *ListSmartSysAvatarModelsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSmartSysAvatarModelsRequest) SetPageSize(v int64) *ListSmartSysAvatarModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmartSysAvatarModelsRequest) SetSdkVersion(v string) *ListSmartSysAvatarModelsRequest {
	s.SdkVersion = &v
	return s
}

func (s *ListSmartSysAvatarModelsRequest) Validate() error {
	return dara.Validate(s)
}
