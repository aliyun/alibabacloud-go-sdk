// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v []*string) *MigrateApplicationRequest
	GetAppIds() []*string
	SetCmd(v string) *MigrateApplicationRequest
	GetCmd() *string
	SetConfig(v string) *MigrateApplicationRequest
	GetConfig() *string
	SetRawData(v string) *MigrateApplicationRequest
	GetRawData() *string
	SetRegionId(v string) *MigrateApplicationRequest
	GetRegionId() *string
}

type MigrateApplicationRequest struct {
	// The list of application IDs.
	AppIds []*string `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	// The operation command. Valid values:
	//
	// - export: Export.
	//
	// - import: Import.
	//
	// example:
	//
	// export
	Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
	// Specifies whether to export the application binary. Default value: false.
	//
	// example:
	//
	// {withBinary:true}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The raw data for the application to be imported, which is sourced from the JSON file of the exported application.
	//
	// example:
	//
	// {"job_id":"b72c0ed4-a69f-4872-b4c6-def5555bfd3e","app_info":"xxxx"
	RawData *string `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// regionId
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s MigrateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateApplicationRequest) GoString() string {
	return s.String()
}

func (s *MigrateApplicationRequest) GetAppIds() []*string {
	return s.AppIds
}

func (s *MigrateApplicationRequest) GetCmd() *string {
	return s.Cmd
}

func (s *MigrateApplicationRequest) GetConfig() *string {
	return s.Config
}

func (s *MigrateApplicationRequest) GetRawData() *string {
	return s.RawData
}

func (s *MigrateApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MigrateApplicationRequest) SetAppIds(v []*string) *MigrateApplicationRequest {
	s.AppIds = v
	return s
}

func (s *MigrateApplicationRequest) SetCmd(v string) *MigrateApplicationRequest {
	s.Cmd = &v
	return s
}

func (s *MigrateApplicationRequest) SetConfig(v string) *MigrateApplicationRequest {
	s.Config = &v
	return s
}

func (s *MigrateApplicationRequest) SetRawData(v string) *MigrateApplicationRequest {
	s.RawData = &v
	return s
}

func (s *MigrateApplicationRequest) SetRegionId(v string) *MigrateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *MigrateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
