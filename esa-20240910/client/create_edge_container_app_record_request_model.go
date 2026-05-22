// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateEdgeContainerAppRecordRequest
	GetAppId() *string
	SetRecordName(v string) *CreateEdgeContainerAppRecordRequest
	GetRecordName() *string
	SetSiteId(v int64) *CreateEdgeContainerAppRecordRequest
	GetSiteId() *int64
}

type CreateEdgeContainerAppRecordRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateEdgeContainerAppRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppRecordRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateEdgeContainerAppRecordRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateEdgeContainerAppRecordRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateEdgeContainerAppRecordRequest) SetAppId(v string) *CreateEdgeContainerAppRecordRequest {
	s.AppId = &v
	return s
}

func (s *CreateEdgeContainerAppRecordRequest) SetRecordName(v string) *CreateEdgeContainerAppRecordRequest {
	s.RecordName = &v
	return s
}

func (s *CreateEdgeContainerAppRecordRequest) SetSiteId(v int64) *CreateEdgeContainerAppRecordRequest {
	s.SiteId = &v
	return s
}

func (s *CreateEdgeContainerAppRecordRequest) Validate() error {
	return dara.Validate(s)
}
