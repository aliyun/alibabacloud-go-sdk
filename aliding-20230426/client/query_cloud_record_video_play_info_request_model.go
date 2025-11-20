// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoPlayInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConferenceId(v string) *QueryCloudRecordVideoPlayInfoRequest
	GetConferenceId() *string
	SetMediaId(v string) *QueryCloudRecordVideoPlayInfoRequest
	GetMediaId() *string
	SetRegionId(v string) *QueryCloudRecordVideoPlayInfoRequest
	GetRegionId() *string
	SetTenantContext(v *QueryCloudRecordVideoPlayInfoRequestTenantContext) *QueryCloudRecordVideoPlayInfoRequest
	GetTenantContext() *QueryCloudRecordVideoPlayInfoRequestTenantContext
}

type QueryCloudRecordVideoPlayInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6139b4xxx
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 44444444
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId      *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantContext *QueryCloudRecordVideoPlayInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryCloudRecordVideoPlayInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryCloudRecordVideoPlayInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryCloudRecordVideoPlayInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCloudRecordVideoPlayInfoRequest) GetTenantContext() *QueryCloudRecordVideoPlayInfoRequestTenantContext {
	return s.TenantContext
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetConferenceId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetMediaId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.MediaId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetRegionId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetTenantContext(v *QueryCloudRecordVideoPlayInfoRequestTenantContext) *QueryCloudRecordVideoPlayInfoRequest {
	s.TenantContext = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCloudRecordVideoPlayInfoRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryCloudRecordVideoPlayInfoRequestTenantContext) SetTenantId(v string) *QueryCloudRecordVideoPlayInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
