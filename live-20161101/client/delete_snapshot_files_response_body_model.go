// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailureCount(v int32) *DeleteSnapshotFilesResponseBody
	GetFailureCount() *int32
	SetRequestId(v string) *DeleteSnapshotFilesResponseBody
	GetRequestId() *string
	SetSnapshotDeleteInfoList(v *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList) *DeleteSnapshotFilesResponseBody
	GetSnapshotDeleteInfoList() *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList
	SetSuccessCount(v int32) *DeleteSnapshotFilesResponseBody
	GetSuccessCount() *int32
}

type DeleteSnapshotFilesResponseBody struct {
	// The number of snapshots that failed to be deleted.
	//
	// example:
	//
	// 1
	FailureCount *int32 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90F60327-ABEC-5A93-BF1F-****
	RequestId              *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotDeleteInfoList *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList `json:"SnapshotDeleteInfoList,omitempty" xml:"SnapshotDeleteInfoList,omitempty" type:"Struct"`
	// The number of successful screenshot deletions.
	//
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DeleteSnapshotFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotFilesResponseBody) GetFailureCount() *int32 {
	return s.FailureCount
}

func (s *DeleteSnapshotFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnapshotFilesResponseBody) GetSnapshotDeleteInfoList() *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList {
	return s.SnapshotDeleteInfoList
}

func (s *DeleteSnapshotFilesResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DeleteSnapshotFilesResponseBody) SetFailureCount(v int32) *DeleteSnapshotFilesResponseBody {
	s.FailureCount = &v
	return s
}

func (s *DeleteSnapshotFilesResponseBody) SetRequestId(v string) *DeleteSnapshotFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotFilesResponseBody) SetSnapshotDeleteInfoList(v *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList) *DeleteSnapshotFilesResponseBody {
	s.SnapshotDeleteInfoList = v
	return s
}

func (s *DeleteSnapshotFilesResponseBody) SetSuccessCount(v int32) *DeleteSnapshotFilesResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DeleteSnapshotFilesResponseBody) Validate() error {
	if s.SnapshotDeleteInfoList != nil {
		if err := s.SnapshotDeleteInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList struct {
	SnapshotDeleteInfo []*DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo `json:"SnapshotDeleteInfo,omitempty" xml:"SnapshotDeleteInfo,omitempty" type:"Repeated"`
}

func (s DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList) GetSnapshotDeleteInfo() []*DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo {
	return s.SnapshotDeleteInfo
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList) SetSnapshotDeleteInfo(v []*DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList {
	s.SnapshotDeleteInfo = v
	return s
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoList) Validate() error {
	if s.SnapshotDeleteInfo != nil {
		for _, item := range s.SnapshotDeleteInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo struct {
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) GetMessage() *string {
	return s.Message
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) SetCreateTimestamp(v int64) *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) SetMessage(v string) *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo {
	s.Message = &v
	return s
}

func (s *DeleteSnapshotFilesResponseBodySnapshotDeleteInfoListSnapshotDeleteInfo) Validate() error {
	return dara.Validate(s)
}
