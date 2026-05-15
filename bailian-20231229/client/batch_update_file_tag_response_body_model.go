// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFileTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchUpdateFileTagResponseBody
	GetCode() *string
	SetData(v *BatchUpdateFileTagResponseBodyData) *BatchUpdateFileTagResponseBody
	GetData() *BatchUpdateFileTagResponseBodyData
	SetMessage(v string) *BatchUpdateFileTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchUpdateFileTagResponseBody
	GetRequestId() *string
	SetStatus(v string) *BatchUpdateFileTagResponseBody
	GetStatus() *string
	SetSuccess(v bool) *BatchUpdateFileTagResponseBody
	GetSuccess() *bool
}

type BatchUpdateFileTagResponseBody struct {
	// example:
	//
	// Success
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BatchUpdateFileTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(FileId) missing or invalid, please check the request parameters.
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUpdateFileTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileTagResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchUpdateFileTagResponseBody) GetData() *BatchUpdateFileTagResponseBodyData {
	return s.Data
}

func (s *BatchUpdateFileTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchUpdateFileTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateFileTagResponseBody) GetStatus() *string {
	return s.Status
}

func (s *BatchUpdateFileTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchUpdateFileTagResponseBody) SetCode(v string) *BatchUpdateFileTagResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUpdateFileTagResponseBody) SetData(v *BatchUpdateFileTagResponseBodyData) *BatchUpdateFileTagResponseBody {
	s.Data = v
	return s
}

func (s *BatchUpdateFileTagResponseBody) SetMessage(v string) *BatchUpdateFileTagResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUpdateFileTagResponseBody) SetRequestId(v string) *BatchUpdateFileTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateFileTagResponseBody) SetStatus(v string) *BatchUpdateFileTagResponseBody {
	s.Status = &v
	return s
}

func (s *BatchUpdateFileTagResponseBody) SetSuccess(v bool) *BatchUpdateFileTagResponseBody {
	s.Success = &v
	return s
}

func (s *BatchUpdateFileTagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchUpdateFileTagResponseBodyData struct {
	UpdateFileTagResultList []*BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList `json:"UpdateFileTagResultList,omitempty" xml:"UpdateFileTagResultList,omitempty" type:"Repeated"`
}

func (s BatchUpdateFileTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileTagResponseBodyData) GetUpdateFileTagResultList() []*BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList {
	return s.UpdateFileTagResultList
}

func (s *BatchUpdateFileTagResponseBodyData) SetUpdateFileTagResultList(v []*BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) *BatchUpdateFileTagResponseBodyData {
	s.UpdateFileTagResultList = v
	return s
}

func (s *BatchUpdateFileTagResponseBodyData) Validate() error {
	if s.UpdateFileTagResultList != nil {
		for _, item := range s.UpdateFileTagResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList struct {
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// FileId not exists.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// file_f40f2a32205d44b4a93b11617113da15_10045951
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) GetFileId() *string {
	return s.FileId
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) GetSuccess() *bool {
	return s.Success
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) SetErrorCode(v string) *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList {
	s.ErrorCode = &v
	return s
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) SetErrorMessage(v string) *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList {
	s.ErrorMessage = &v
	return s
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) SetFileId(v string) *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList {
	s.FileId = &v
	return s
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) SetSuccess(v bool) *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList {
	s.Success = &v
	return s
}

func (s *BatchUpdateFileTagResponseBodyDataUpdateFileTagResultList) Validate() error {
	return dara.Validate(s)
}
