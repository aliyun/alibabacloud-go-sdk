// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupConfigByIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupId(v int64) *GetApgroupConfigByIdentityRequest
	GetApgroupId() *int64
	SetApgroupUuid(v string) *GetApgroupConfigByIdentityRequest
	GetApgroupUuid() *string
	SetAppCode(v string) *GetApgroupConfigByIdentityRequest
	GetAppCode() *string
	SetAppName(v string) *GetApgroupConfigByIdentityRequest
	GetAppName() *string
}

type GetApgroupConfigByIdentityRequest struct {
	// This parameter is required.
	ApgroupId *int64 `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
	// This parameter is required.
	ApgroupUuid *string `json:"ApgroupUuid,omitempty" xml:"ApgroupUuid,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetApgroupConfigByIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupConfigByIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetApgroupConfigByIdentityRequest) GetApgroupId() *int64 {
	return s.ApgroupId
}

func (s *GetApgroupConfigByIdentityRequest) GetApgroupUuid() *string {
	return s.ApgroupUuid
}

func (s *GetApgroupConfigByIdentityRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApgroupConfigByIdentityRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApgroupConfigByIdentityRequest) SetApgroupId(v int64) *GetApgroupConfigByIdentityRequest {
	s.ApgroupId = &v
	return s
}

func (s *GetApgroupConfigByIdentityRequest) SetApgroupUuid(v string) *GetApgroupConfigByIdentityRequest {
	s.ApgroupUuid = &v
	return s
}

func (s *GetApgroupConfigByIdentityRequest) SetAppCode(v string) *GetApgroupConfigByIdentityRequest {
	s.AppCode = &v
	return s
}

func (s *GetApgroupConfigByIdentityRequest) SetAppName(v string) *GetApgroupConfigByIdentityRequest {
	s.AppName = &v
	return s
}

func (s *GetApgroupConfigByIdentityRequest) Validate() error {
	return dara.Validate(s)
}
