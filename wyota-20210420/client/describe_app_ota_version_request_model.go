// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppOtaVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *DescribeAppOtaVersionRequest
	GetAppVersion() *string
	SetChannel(v string) *DescribeAppOtaVersionRequest
	GetChannel() *string
	SetClientType(v int32) *DescribeAppOtaVersionRequest
	GetClientType() *int32
	SetCreator(v string) *DescribeAppOtaVersionRequest
	GetCreator() *string
	SetNullChannel(v bool) *DescribeAppOtaVersionRequest
	GetNullChannel() *bool
	SetOtaType(v int32) *DescribeAppOtaVersionRequest
	GetOtaType() *int32
	SetProject(v string) *DescribeAppOtaVersionRequest
	GetProject() *string
	SetStatus(v int32) *DescribeAppOtaVersionRequest
	GetStatus() *int32
	SetVersionUid(v string) *DescribeAppOtaVersionRequest
	GetVersionUid() *string
}

type DescribeAppOtaVersionRequest struct {
	AppVersion  *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Channel     *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientType  *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	NullChannel *bool   `json:"NullChannel,omitempty" xml:"NullChannel,omitempty"`
	OtaType     *int32  `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionUid  *string `json:"VersionUid,omitempty" xml:"VersionUid,omitempty"`
}

func (s DescribeAppOtaVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppOtaVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeAppOtaVersionRequest) GetChannel() *string {
	return s.Channel
}

func (s *DescribeAppOtaVersionRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *DescribeAppOtaVersionRequest) GetCreator() *string {
	return s.Creator
}

func (s *DescribeAppOtaVersionRequest) GetNullChannel() *bool {
	return s.NullChannel
}

func (s *DescribeAppOtaVersionRequest) GetOtaType() *int32 {
	return s.OtaType
}

func (s *DescribeAppOtaVersionRequest) GetProject() *string {
	return s.Project
}

func (s *DescribeAppOtaVersionRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAppOtaVersionRequest) GetVersionUid() *string {
	return s.VersionUid
}

func (s *DescribeAppOtaVersionRequest) SetAppVersion(v string) *DescribeAppOtaVersionRequest {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetChannel(v string) *DescribeAppOtaVersionRequest {
	s.Channel = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetClientType(v int32) *DescribeAppOtaVersionRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetCreator(v string) *DescribeAppOtaVersionRequest {
	s.Creator = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetNullChannel(v bool) *DescribeAppOtaVersionRequest {
	s.NullChannel = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetOtaType(v int32) *DescribeAppOtaVersionRequest {
	s.OtaType = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetProject(v string) *DescribeAppOtaVersionRequest {
	s.Project = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetStatus(v int32) *DescribeAppOtaVersionRequest {
	s.Status = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) SetVersionUid(v string) *DescribeAppOtaVersionRequest {
	s.VersionUid = &v
	return s
}

func (s *DescribeAppOtaVersionRequest) Validate() error {
	return dara.Validate(s)
}
