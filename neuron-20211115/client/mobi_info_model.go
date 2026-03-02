// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobiInfo interface {
	dara.Model
	String() string
	GoString() string
	SetMobiCommitId(v string) *MobiInfo
	GetMobiCommitId() *string
	SetMobiId(v int64) *MobiInfo
	GetMobiId() *int64
	SetMobiModuleId(v string) *MobiInfo
	GetMobiModuleId() *string
	SetStoreUrl(v string) *MobiInfo
	GetStoreUrl() *string
}

type MobiInfo struct {
	// example:
	//
	// commit_pckesd7d_20230227215101
	MobiCommitId *string `json:"mobiCommitId,omitempty" xml:"mobiCommitId,omitempty"`
	// example:
	//
	// 1
	MobiId *int64 `json:"mobiId,omitempty" xml:"mobiId,omitempty"`
	// example:
	//
	// module_tvtpydeq
	MobiModuleId *string `json:"mobiModuleId,omitempty" xml:"mobiModuleId,omitempty"`
	// example:
	//
	// sda
	StoreUrl *string `json:"storeUrl,omitempty" xml:"storeUrl,omitempty"`
}

func (s MobiInfo) String() string {
	return dara.Prettify(s)
}

func (s MobiInfo) GoString() string {
	return s.String()
}

func (s *MobiInfo) GetMobiCommitId() *string {
	return s.MobiCommitId
}

func (s *MobiInfo) GetMobiId() *int64 {
	return s.MobiId
}

func (s *MobiInfo) GetMobiModuleId() *string {
	return s.MobiModuleId
}

func (s *MobiInfo) GetStoreUrl() *string {
	return s.StoreUrl
}

func (s *MobiInfo) SetMobiCommitId(v string) *MobiInfo {
	s.MobiCommitId = &v
	return s
}

func (s *MobiInfo) SetMobiId(v int64) *MobiInfo {
	s.MobiId = &v
	return s
}

func (s *MobiInfo) SetMobiModuleId(v string) *MobiInfo {
	s.MobiModuleId = &v
	return s
}

func (s *MobiInfo) SetStoreUrl(v string) *MobiInfo {
	s.StoreUrl = &v
	return s
}

func (s *MobiInfo) Validate() error {
	return dara.Validate(s)
}
