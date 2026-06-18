// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportTaskNumberDatasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *ImportTaskNumberDatasShrinkRequest
	GetDataType() *string
	SetEncryptionType(v int64) *ImportTaskNumberDatasShrinkRequest
	GetEncryptionType() *int64
	SetOssFileName(v string) *ImportTaskNumberDatasShrinkRequest
	GetOssFileName() *string
	SetOwnerId(v int64) *ImportTaskNumberDatasShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumberListShrink(v string) *ImportTaskNumberDatasShrinkRequest
	GetPhoneNumberListShrink() *string
	SetResourceOwnerAccount(v string) *ImportTaskNumberDatasShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportTaskNumberDatasShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *ImportTaskNumberDatasShrinkRequest
	GetTaskId() *string
}

type ImportTaskNumberDatasShrinkRequest struct {
	// The data type. Valid values:
	//
	// - EXCEL
	//
	// - JSON
	//
	//
	//   	Notice:
	//
	//   API calls currently support only the JSON data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// JSON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 1
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The OSS file path. This parameter is optional.
	//
	// > Importing data by specifying an OSS file path is not available because API calls currently support only the JSON data type.
	//
	// example:
	//
	// 123dsdfsdfsdf.xlsx
	OssFileName *string `json:"OssFileName,omitempty" xml:"OssFileName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// If `DataType` is set to `JSON`, you must use this parameter to upload the data. You can import up to 1,000 records per request.
	PhoneNumberListShrink *string `json:"PhoneNumberList,omitempty" xml:"PhoneNumberList,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the call task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 119181071278******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportTaskNumberDatasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportTaskNumberDatasShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportTaskNumberDatasShrinkRequest) GetDataType() *string {
	return s.DataType
}

func (s *ImportTaskNumberDatasShrinkRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *ImportTaskNumberDatasShrinkRequest) GetOssFileName() *string {
	return s.OssFileName
}

func (s *ImportTaskNumberDatasShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportTaskNumberDatasShrinkRequest) GetPhoneNumberListShrink() *string {
	return s.PhoneNumberListShrink
}

func (s *ImportTaskNumberDatasShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportTaskNumberDatasShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportTaskNumberDatasShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ImportTaskNumberDatasShrinkRequest) SetDataType(v string) *ImportTaskNumberDatasShrinkRequest {
	s.DataType = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) SetEncryptionType(v int64) *ImportTaskNumberDatasShrinkRequest {
	s.EncryptionType = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) SetOssFileName(v string) *ImportTaskNumberDatasShrinkRequest {
	s.OssFileName = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) SetOwnerId(v int64) *ImportTaskNumberDatasShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) SetPhoneNumberListShrink(v string) *ImportTaskNumberDatasShrinkRequest {
	s.PhoneNumberListShrink = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) SetResourceOwnerAccount(v string) *ImportTaskNumberDatasShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) SetResourceOwnerId(v int64) *ImportTaskNumberDatasShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) SetTaskId(v string) *ImportTaskNumberDatasShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ImportTaskNumberDatasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
