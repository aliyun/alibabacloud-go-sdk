// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamRecordIndexFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLiveStreamRecordIndexFilesResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteLiveStreamRecordIndexFilesResponseBody
	GetMessage() *string
	SetRecordDeleteInfoList(v *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList) *DeleteLiveStreamRecordIndexFilesResponseBody
	GetRecordDeleteInfoList() *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList
	SetRequestId(v string) *DeleteLiveStreamRecordIndexFilesResponseBody
	GetRequestId() *string
}

type DeleteLiveStreamRecordIndexFilesResponseBody struct {
	// The status code. A return value of 500 indicates an error. For details, see the Error codes section of this topic.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status description. A return value of 500 indicates an error. For details, see the Error codes section of this topic.
	//
	// example:
	//
	// OK
	Message              *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RecordDeleteInfoList *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList `json:"RecordDeleteInfoList,omitempty" xml:"RecordDeleteInfoList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 939D19EE-59A0-18E9-B458-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveStreamRecordIndexFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamRecordIndexFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) GetRecordDeleteInfoList() *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList {
	return s.RecordDeleteInfoList
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) SetCode(v string) *DeleteLiveStreamRecordIndexFilesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) SetMessage(v string) *DeleteLiveStreamRecordIndexFilesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) SetRecordDeleteInfoList(v *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList) *DeleteLiveStreamRecordIndexFilesResponseBody {
	s.RecordDeleteInfoList = v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) SetRequestId(v string) *DeleteLiveStreamRecordIndexFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBody) Validate() error {
	if s.RecordDeleteInfoList != nil {
		if err := s.RecordDeleteInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList struct {
	RecordDeleteInfo []*DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo `json:"RecordDeleteInfo,omitempty" xml:"RecordDeleteInfo,omitempty" type:"Repeated"`
}

func (s DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList) GetRecordDeleteInfo() []*DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo {
	return s.RecordDeleteInfo
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList) SetRecordDeleteInfo(v []*DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList {
	s.RecordDeleteInfo = v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoList) Validate() error {
	if s.RecordDeleteInfo != nil {
		for _, item := range s.RecordDeleteInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo struct {
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) GetMessage() *string {
	return s.Message
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) GetRecordId() *string {
	return s.RecordId
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) SetMessage(v string) *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo {
	s.Message = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) SetRecordId(v string) *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo {
	s.RecordId = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponseBodyRecordDeleteInfoListRecordDeleteInfo) Validate() error {
	return dara.Validate(s)
}
