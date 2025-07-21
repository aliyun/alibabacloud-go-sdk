// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppOtaLatestVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseVersion(v string) *GetAppOtaLatestVersionRequest
	GetBaseVersion() *string
	SetClientType(v int32) *GetAppOtaLatestVersionRequest
	GetClientType() *int32
	SetClientUid(v string) *GetAppOtaLatestVersionRequest
	GetClientUid() *string
	SetOsType(v string) *GetAppOtaLatestVersionRequest
	GetOsType() *string
	SetProject(v string) *GetAppOtaLatestVersionRequest
	GetProject() *string
}

type GetAppOtaLatestVersionRequest struct {
	// This parameter is required.
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	ClientType  *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientUid   *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	// This parameter is required.
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s GetAppOtaLatestVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppOtaLatestVersionRequest) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionRequest) GetBaseVersion() *string {
	return s.BaseVersion
}

func (s *GetAppOtaLatestVersionRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetAppOtaLatestVersionRequest) GetClientUid() *string {
	return s.ClientUid
}

func (s *GetAppOtaLatestVersionRequest) GetOsType() *string {
	return s.OsType
}

func (s *GetAppOtaLatestVersionRequest) GetProject() *string {
	return s.Project
}

func (s *GetAppOtaLatestVersionRequest) SetBaseVersion(v string) *GetAppOtaLatestVersionRequest {
	s.BaseVersion = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetClientType(v int32) *GetAppOtaLatestVersionRequest {
	s.ClientType = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetClientUid(v string) *GetAppOtaLatestVersionRequest {
	s.ClientUid = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetOsType(v string) *GetAppOtaLatestVersionRequest {
	s.OsType = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) SetProject(v string) *GetAppOtaLatestVersionRequest {
	s.Project = &v
	return s
}

func (s *GetAppOtaLatestVersionRequest) Validate() error {
	return dara.Validate(s)
}
