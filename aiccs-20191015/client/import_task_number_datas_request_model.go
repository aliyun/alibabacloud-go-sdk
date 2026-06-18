// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportTaskNumberDatasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *ImportTaskNumberDatasRequest
	GetDataType() *string
	SetEncryptionType(v int64) *ImportTaskNumberDatasRequest
	GetEncryptionType() *int64
	SetOssFileName(v string) *ImportTaskNumberDatasRequest
	GetOssFileName() *string
	SetOwnerId(v int64) *ImportTaskNumberDatasRequest
	GetOwnerId() *int64
	SetPhoneNumberList(v []map[string]interface{}) *ImportTaskNumberDatasRequest
	GetPhoneNumberList() []map[string]interface{}
	SetResourceOwnerAccount(v string) *ImportTaskNumberDatasRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportTaskNumberDatasRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *ImportTaskNumberDatasRequest
	GetTaskId() *string
}

type ImportTaskNumberDatasRequest struct {
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
	PhoneNumberList      []map[string]interface{} `json:"PhoneNumberList,omitempty" xml:"PhoneNumberList,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the call task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 119181071278******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportTaskNumberDatasRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportTaskNumberDatasRequest) GoString() string {
	return s.String()
}

func (s *ImportTaskNumberDatasRequest) GetDataType() *string {
	return s.DataType
}

func (s *ImportTaskNumberDatasRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *ImportTaskNumberDatasRequest) GetOssFileName() *string {
	return s.OssFileName
}

func (s *ImportTaskNumberDatasRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportTaskNumberDatasRequest) GetPhoneNumberList() []map[string]interface{} {
	return s.PhoneNumberList
}

func (s *ImportTaskNumberDatasRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportTaskNumberDatasRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportTaskNumberDatasRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ImportTaskNumberDatasRequest) SetDataType(v string) *ImportTaskNumberDatasRequest {
	s.DataType = &v
	return s
}

func (s *ImportTaskNumberDatasRequest) SetEncryptionType(v int64) *ImportTaskNumberDatasRequest {
	s.EncryptionType = &v
	return s
}

func (s *ImportTaskNumberDatasRequest) SetOssFileName(v string) *ImportTaskNumberDatasRequest {
	s.OssFileName = &v
	return s
}

func (s *ImportTaskNumberDatasRequest) SetOwnerId(v int64) *ImportTaskNumberDatasRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportTaskNumberDatasRequest) SetPhoneNumberList(v []map[string]interface{}) *ImportTaskNumberDatasRequest {
	s.PhoneNumberList = v
	return s
}

func (s *ImportTaskNumberDatasRequest) SetResourceOwnerAccount(v string) *ImportTaskNumberDatasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportTaskNumberDatasRequest) SetResourceOwnerId(v int64) *ImportTaskNumberDatasRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportTaskNumberDatasRequest) SetTaskId(v string) *ImportTaskNumberDatasRequest {
	s.TaskId = &v
	return s
}

func (s *ImportTaskNumberDatasRequest) Validate() error {
	return dara.Validate(s)
}
