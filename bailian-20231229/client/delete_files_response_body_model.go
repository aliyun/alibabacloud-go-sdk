// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFilesResponseBody
	GetCode() *string
	SetData(v *DeleteFilesResponseBodyData) *DeleteFilesResponseBody
	GetData() *DeleteFilesResponseBodyData
	SetMessage(v string) *DeleteFilesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFilesResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteFilesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteFilesResponseBody
	GetSuccess() *bool
}

type DeleteFilesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-7734-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the API.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFilesResponseBody) GetData() *DeleteFilesResponseBodyData {
	return s.Data
}

func (s *DeleteFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFilesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFilesResponseBody) SetCode(v string) *DeleteFilesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFilesResponseBody) SetData(v *DeleteFilesResponseBodyData) *DeleteFilesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteFilesResponseBody) SetMessage(v string) *DeleteFilesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFilesResponseBody) SetRequestId(v string) *DeleteFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFilesResponseBody) SetStatus(v string) *DeleteFilesResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteFilesResponseBody) SetSuccess(v bool) *DeleteFilesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFilesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteFilesResponseBodyData struct {
	// The deletion results.
	DeleteFileResultList []*DeleteFilesResponseBodyDataDeleteFileResultList `json:"DeleteFileResultList,omitempty" xml:"DeleteFileResultList,omitempty" type:"Repeated"`
}

func (s DeleteFilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteFilesResponseBodyData) GetDeleteFileResultList() []*DeleteFilesResponseBodyDataDeleteFileResultList {
	return s.DeleteFileResultList
}

func (s *DeleteFilesResponseBodyData) SetDeleteFileResultList(v []*DeleteFilesResponseBodyDataDeleteFileResultList) *DeleteFilesResponseBodyData {
	s.DeleteFileResultList = v
	return s
}

func (s *DeleteFilesResponseBodyData) Validate() error {
	if s.DeleteFileResultList != nil {
		for _, item := range s.DeleteFileResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteFilesResponseBodyDataDeleteFileResultList struct {
	// The file ID.
	//
	// example:
	//
	// file_6b193b9b4b1546ef9eaa7340e69adfca_10052857
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file deletion status. Valid values:
	//
	// - DELETED: The file is deleted.
	//
	// - FAILED: The file fails to be deleted.
	//
	// - NOT_FOUND: The file is not found.
	//
	// example:
	//
	// DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteFilesResponseBodyDataDeleteFileResultList) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesResponseBodyDataDeleteFileResultList) GoString() string {
	return s.String()
}

func (s *DeleteFilesResponseBodyDataDeleteFileResultList) GetFileId() *string {
	return s.FileId
}

func (s *DeleteFilesResponseBodyDataDeleteFileResultList) GetStatus() *string {
	return s.Status
}

func (s *DeleteFilesResponseBodyDataDeleteFileResultList) SetFileId(v string) *DeleteFilesResponseBodyDataDeleteFileResultList {
	s.FileId = &v
	return s
}

func (s *DeleteFilesResponseBodyDataDeleteFileResultList) SetStatus(v string) *DeleteFilesResponseBodyDataDeleteFileResultList {
	s.Status = &v
	return s
}

func (s *DeleteFilesResponseBodyDataDeleteFileResultList) Validate() error {
	return dara.Validate(s)
}
