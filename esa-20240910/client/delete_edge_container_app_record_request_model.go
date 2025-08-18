// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteEdgeContainerAppRecordRequest
	GetAppId() *string
	SetRecordName(v string) *DeleteEdgeContainerAppRecordRequest
	GetRecordName() *string
	SetSiteId(v int64) *DeleteEdgeContainerAppRecordRequest
	GetSiteId() *int64
}

type DeleteEdgeContainerAppRecordRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The associated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteEdgeContainerAppRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppRecordRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteEdgeContainerAppRecordRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *DeleteEdgeContainerAppRecordRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteEdgeContainerAppRecordRequest) SetAppId(v string) *DeleteEdgeContainerAppRecordRequest {
	s.AppId = &v
	return s
}

func (s *DeleteEdgeContainerAppRecordRequest) SetRecordName(v string) *DeleteEdgeContainerAppRecordRequest {
	s.RecordName = &v
	return s
}

func (s *DeleteEdgeContainerAppRecordRequest) SetSiteId(v int64) *DeleteEdgeContainerAppRecordRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteEdgeContainerAppRecordRequest) Validate() error {
	return dara.Validate(s)
}
