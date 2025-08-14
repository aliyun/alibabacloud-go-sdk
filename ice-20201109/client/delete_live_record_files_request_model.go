// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordIds(v []*string) *DeleteLiveRecordFilesRequest
	GetRecordIds() []*string
	SetRemoveFile(v bool) *DeleteLiveRecordFilesRequest
	GetRemoveFile() *bool
}

type DeleteLiveRecordFilesRequest struct {
	// The collection of IDs of recording files.
	//
	// This parameter is required.
	RecordIds []*string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
	// Specifies whether to delete the original files in OSS.
	//
	// example:
	//
	// true
	RemoveFile *bool `json:"RemoveFile,omitempty" xml:"RemoveFile,omitempty"`
}

func (s DeleteLiveRecordFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordFilesRequest) GetRecordIds() []*string {
	return s.RecordIds
}

func (s *DeleteLiveRecordFilesRequest) GetRemoveFile() *bool {
	return s.RemoveFile
}

func (s *DeleteLiveRecordFilesRequest) SetRecordIds(v []*string) *DeleteLiveRecordFilesRequest {
	s.RecordIds = v
	return s
}

func (s *DeleteLiveRecordFilesRequest) SetRemoveFile(v bool) *DeleteLiveRecordFilesRequest {
	s.RemoveFile = &v
	return s
}

func (s *DeleteLiveRecordFilesRequest) Validate() error {
	return dara.Validate(s)
}
