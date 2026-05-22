// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackEdgeContainerAppVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RollbackEdgeContainerAppVersionRequest
	GetAppId() *string
	SetPercentage(v int32) *RollbackEdgeContainerAppVersionRequest
	GetPercentage() *int32
	SetRemarks(v string) *RollbackEdgeContainerAppVersionRequest
	GetRemarks() *string
	SetUsedPercent(v bool) *RollbackEdgeContainerAppVersionRequest
	GetUsedPercent() *bool
	SetVersionId(v string) *RollbackEdgeContainerAppVersionRequest
	GetVersionId() *string
}

type RollbackEdgeContainerAppVersionRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test rollback app
	Remarks     *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	UsedPercent *bool   `json:"UsedPercent,omitempty" xml:"UsedPercent,omitempty"`
	// The ID of version that you want to roll back.
	//
	// This parameter is required.
	//
	// example:
	//
	// ver-87962637161651****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RollbackEdgeContainerAppVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackEdgeContainerAppVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackEdgeContainerAppVersionRequest) GetAppId() *string {
	return s.AppId
}

func (s *RollbackEdgeContainerAppVersionRequest) GetPercentage() *int32 {
	return s.Percentage
}

func (s *RollbackEdgeContainerAppVersionRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *RollbackEdgeContainerAppVersionRequest) GetUsedPercent() *bool {
	return s.UsedPercent
}

func (s *RollbackEdgeContainerAppVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *RollbackEdgeContainerAppVersionRequest) SetAppId(v string) *RollbackEdgeContainerAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *RollbackEdgeContainerAppVersionRequest) SetPercentage(v int32) *RollbackEdgeContainerAppVersionRequest {
	s.Percentage = &v
	return s
}

func (s *RollbackEdgeContainerAppVersionRequest) SetRemarks(v string) *RollbackEdgeContainerAppVersionRequest {
	s.Remarks = &v
	return s
}

func (s *RollbackEdgeContainerAppVersionRequest) SetUsedPercent(v bool) *RollbackEdgeContainerAppVersionRequest {
	s.UsedPercent = &v
	return s
}

func (s *RollbackEdgeContainerAppVersionRequest) SetVersionId(v string) *RollbackEdgeContainerAppVersionRequest {
	s.VersionId = &v
	return s
}

func (s *RollbackEdgeContainerAppVersionRequest) Validate() error {
	return dara.Validate(s)
}
