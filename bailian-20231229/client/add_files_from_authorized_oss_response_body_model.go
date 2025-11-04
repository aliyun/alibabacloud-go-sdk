// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilesFromAuthorizedOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddFilesFromAuthorizedOssResponseBody
	GetCode() *string
	SetData(v *AddFilesFromAuthorizedOssResponseBodyData) *AddFilesFromAuthorizedOssResponseBody
	GetData() *AddFilesFromAuthorizedOssResponseBodyData
	SetMessage(v string) *AddFilesFromAuthorizedOssResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddFilesFromAuthorizedOssResponseBody
	GetRequestId() *string
	SetStatus(v string) *AddFilesFromAuthorizedOssResponseBody
	GetStatus() *string
	SetSuccess(v string) *AddFilesFromAuthorizedOssResponseBody
	GetSuccess() *string
}

type AddFilesFromAuthorizedOssResponseBody struct {
	// example:
	//
	// success
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddFilesFromAuthorizedOssResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Cant find out category for category_id param.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFilesFromAuthorizedOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssResponseBody) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddFilesFromAuthorizedOssResponseBody) GetData() *AddFilesFromAuthorizedOssResponseBodyData {
	return s.Data
}

func (s *AddFilesFromAuthorizedOssResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddFilesFromAuthorizedOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFilesFromAuthorizedOssResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AddFilesFromAuthorizedOssResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AddFilesFromAuthorizedOssResponseBody) SetCode(v string) *AddFilesFromAuthorizedOssResponseBody {
	s.Code = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBody) SetData(v *AddFilesFromAuthorizedOssResponseBodyData) *AddFilesFromAuthorizedOssResponseBody {
	s.Data = v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBody) SetMessage(v string) *AddFilesFromAuthorizedOssResponseBody {
	s.Message = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBody) SetRequestId(v string) *AddFilesFromAuthorizedOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBody) SetStatus(v string) *AddFilesFromAuthorizedOssResponseBody {
	s.Status = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBody) SetSuccess(v string) *AddFilesFromAuthorizedOssResponseBody {
	s.Success = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFilesFromAuthorizedOssResponseBodyData struct {
	AddFileResultList []*AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList `json:"AddFileResultList,omitempty" xml:"AddFileResultList,omitempty" type:"Repeated"`
}

func (s AddFilesFromAuthorizedOssResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssResponseBodyData) GetAddFileResultList() []*AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList {
	return s.AddFileResultList
}

func (s *AddFilesFromAuthorizedOssResponseBodyData) SetAddFileResultList(v []*AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) *AddFilesFromAuthorizedOssResponseBodyData {
	s.AddFileResultList = v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBodyData) Validate() error {
	if s.AddFileResultList != nil {
		for _, item := range s.AddFileResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList struct {
	// example:
	//
	// file_809f469a59ac449586ec692576xxxxx_102248XXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// size too large
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// root/path/this_is_temp_xxxx.pdf
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) GetFileId() *string {
	return s.FileId
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) GetMsg() *string {
	return s.Msg
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) GetOssKey() *string {
	return s.OssKey
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) GetStatus() *string {
	return s.Status
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) SetFileId(v string) *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList {
	s.FileId = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) SetMsg(v string) *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList {
	s.Msg = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) SetOssKey(v string) *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList {
	s.OssKey = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) SetStatus(v string) *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList {
	s.Status = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponseBodyDataAddFileResultList) Validate() error {
	return dara.Validate(s)
}
