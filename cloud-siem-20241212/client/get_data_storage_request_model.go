// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetDataStorageRequest
	GetLang() *string
	SetRegionId(v string) *GetDataStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetDataStorageRequest
	GetRoleFor() *int64
	SetWithAutoKeyDetect(v bool) *GetDataStorageRequest
	GetWithAutoKeyDetect() *bool
}

type GetDataStorageRequest struct {
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required in actual API calls. If it is not provided, the service returns a 400 error with the message "The specified Lang can not be empty." Explicitly pass the language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the data management center for threat analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// Specifies whether to read the auto-discovered index configuration of each Logstore. This operation requires accessing Simple Log Service (SLS) for each Logstore. Default value: false. Set this parameter to true only on the storage management page.
	WithAutoKeyDetect *bool `json:"WithAutoKeyDetect,omitempty" xml:"WithAutoKeyDetect,omitempty"`
}

func (s GetDataStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageRequest) GoString() string {
	return s.String()
}

func (s *GetDataStorageRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDataStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDataStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetDataStorageRequest) GetWithAutoKeyDetect() *bool {
	return s.WithAutoKeyDetect
}

func (s *GetDataStorageRequest) SetLang(v string) *GetDataStorageRequest {
	s.Lang = &v
	return s
}

func (s *GetDataStorageRequest) SetRegionId(v string) *GetDataStorageRequest {
	s.RegionId = &v
	return s
}

func (s *GetDataStorageRequest) SetRoleFor(v int64) *GetDataStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *GetDataStorageRequest) SetWithAutoKeyDetect(v bool) *GetDataStorageRequest {
	s.WithAutoKeyDetect = &v
	return s
}

func (s *GetDataStorageRequest) Validate() error {
	return dara.Validate(s)
}
