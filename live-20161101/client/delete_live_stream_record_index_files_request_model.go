// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamRecordIndexFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveStreamRecordIndexFilesRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveStreamRecordIndexFilesRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveStreamRecordIndexFilesRequest
	GetOwnerId() *int64
	SetRecordId(v []*string) *DeleteLiveStreamRecordIndexFilesRequest
	GetRecordId() []*string
	SetRegionId(v string) *DeleteLiveStreamRecordIndexFilesRequest
	GetRegionId() *string
	SetRemoveFile(v string) *DeleteLiveStreamRecordIndexFilesRequest
	GetRemoveFile() *string
	SetStreamName(v string) *DeleteLiveStreamRecordIndexFilesRequest
	GetStreamName() *string
}

type DeleteLiveStreamRecordIndexFilesRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The index file IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c4d7f0a4-b506-43f9-8de3-07732c3f****
	RecordId []*string `json:"RecordId,omitempty" xml:"RecordId,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to delete the corresponding file in Object Storage Service (OSS) synchronously. Valid values:
	//
	// 	- **true**: The corresponding file in OSS is deleted.
	//
	// 	- **false**: The corresponding file in OSS is not deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	RemoveFile *string `json:"RemoveFile,omitempty" xml:"RemoveFile,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveStreamRecordIndexFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamRecordIndexFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) GetRecordId() []*string {
	return s.RecordId
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) GetRemoveFile() *string {
	return s.RemoveFile
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) SetAppName(v string) *DeleteLiveStreamRecordIndexFilesRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) SetDomainName(v string) *DeleteLiveStreamRecordIndexFilesRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) SetOwnerId(v int64) *DeleteLiveStreamRecordIndexFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) SetRecordId(v []*string) *DeleteLiveStreamRecordIndexFilesRequest {
	s.RecordId = v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) SetRegionId(v string) *DeleteLiveStreamRecordIndexFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) SetRemoveFile(v string) *DeleteLiveStreamRecordIndexFilesRequest {
	s.RemoveFile = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) SetStreamName(v string) *DeleteLiveStreamRecordIndexFilesRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesRequest) Validate() error {
	return dara.Validate(s)
}
