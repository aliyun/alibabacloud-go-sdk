// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupDetailedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupId(v int64) *GetApgroupDetailedConfigRequest
	GetApgroupId() *int64
	SetApgroupUuid(v string) *GetApgroupDetailedConfigRequest
	GetApgroupUuid() *string
	SetAppCode(v string) *GetApgroupDetailedConfigRequest
	GetAppCode() *string
	SetAppName(v string) *GetApgroupDetailedConfigRequest
	GetAppName() *string
}

type GetApgroupDetailedConfigRequest struct {
	// This parameter is required.
	ApgroupId *int64 `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
	// This parameter is required.
	ApgroupUuid *string `json:"ApgroupUuid,omitempty" xml:"ApgroupUuid,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetApgroupDetailedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupDetailedConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApgroupDetailedConfigRequest) GetApgroupId() *int64 {
	return s.ApgroupId
}

func (s *GetApgroupDetailedConfigRequest) GetApgroupUuid() *string {
	return s.ApgroupUuid
}

func (s *GetApgroupDetailedConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApgroupDetailedConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApgroupDetailedConfigRequest) SetApgroupId(v int64) *GetApgroupDetailedConfigRequest {
	s.ApgroupId = &v
	return s
}

func (s *GetApgroupDetailedConfigRequest) SetApgroupUuid(v string) *GetApgroupDetailedConfigRequest {
	s.ApgroupUuid = &v
	return s
}

func (s *GetApgroupDetailedConfigRequest) SetAppCode(v string) *GetApgroupDetailedConfigRequest {
	s.AppCode = &v
	return s
}

func (s *GetApgroupDetailedConfigRequest) SetAppName(v string) *GetApgroupDetailedConfigRequest {
	s.AppName = &v
	return s
}

func (s *GetApgroupDetailedConfigRequest) Validate() error {
	return dara.Validate(s)
}
