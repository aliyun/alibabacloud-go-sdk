// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamMergeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveStreamMergeRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveStreamMergeRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveStreamMergeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveStreamMergeRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLiveStreamMergeRequest
	GetStreamName() *string
}

type DeleteLiveStreamMergeRequest struct {
	// The name of the application that generates the output stream. You can view the application name on the [Primary/Secondary Stream Mixing](https://help.aliyun.com/document_detail/606583.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the output stream. You can view the stream name on the [Primary/Secondary Stream Mixing](https://help.aliyun.com/document_detail/606583.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveStreamMergeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamMergeRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamMergeRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveStreamMergeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveStreamMergeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamMergeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveStreamMergeRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLiveStreamMergeRequest) SetAppName(v string) *DeleteLiveStreamMergeRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveStreamMergeRequest) SetDomainName(v string) *DeleteLiveStreamMergeRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveStreamMergeRequest) SetOwnerId(v int64) *DeleteLiveStreamMergeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamMergeRequest) SetRegionId(v string) *DeleteLiveStreamMergeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveStreamMergeRequest) SetStreamName(v string) *DeleteLiveStreamMergeRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLiveStreamMergeRequest) Validate() error {
	return dara.Validate(s)
}
