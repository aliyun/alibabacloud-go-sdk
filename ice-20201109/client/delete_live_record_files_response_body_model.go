// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteFileInfoList(v []*DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) *DeleteLiveRecordFilesResponseBody
	GetDeleteFileInfoList() []*DeleteLiveRecordFilesResponseBodyDeleteFileInfoList
	SetMessage(v string) *DeleteLiveRecordFilesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLiveRecordFilesResponseBody
	GetRequestId() *string
}

type DeleteLiveRecordFilesResponseBody struct {
	// The list of files deleted.
	DeleteFileInfoList []*DeleteLiveRecordFilesResponseBodyDeleteFileInfoList `json:"DeleteFileInfoList,omitempty" xml:"DeleteFileInfoList,omitempty" type:"Repeated"`
	// The description of the state returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 13cbb83e-043c-4728-ac35-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRecordFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordFilesResponseBody) GetDeleteFileInfoList() []*DeleteLiveRecordFilesResponseBodyDeleteFileInfoList {
	return s.DeleteFileInfoList
}

func (s *DeleteLiveRecordFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLiveRecordFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveRecordFilesResponseBody) SetDeleteFileInfoList(v []*DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) *DeleteLiveRecordFilesResponseBody {
	s.DeleteFileInfoList = v
	return s
}

func (s *DeleteLiveRecordFilesResponseBody) SetMessage(v string) *DeleteLiveRecordFilesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLiveRecordFilesResponseBody) SetRequestId(v string) *DeleteLiveRecordFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveRecordFilesResponseBody) Validate() error {
	if s.DeleteFileInfoList != nil {
		for _, item := range s.DeleteFileInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteLiveRecordFilesResponseBodyDeleteFileInfoList struct {
	// The code that identifies the result of the deletion.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of deletion.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the deleted recording file.
	//
	// example:
	//
	// 13cbb83e-043c-4728-ac35-*****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) GetCode() *string {
	return s.Code
}

func (s *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) GetMessage() *string {
	return s.Message
}

func (s *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) GetRecordId() *string {
	return s.RecordId
}

func (s *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) SetCode(v string) *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList {
	s.Code = &v
	return s
}

func (s *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) SetMessage(v string) *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList {
	s.Message = &v
	return s
}

func (s *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) SetRecordId(v string) *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList {
	s.RecordId = &v
	return s
}

func (s *DeleteLiveRecordFilesResponseBodyDeleteFileInfoList) Validate() error {
	return dara.Validate(s)
}
