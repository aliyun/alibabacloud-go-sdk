// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHanaBackupSettingResponseBody
	GetCode() *string
	SetHanaBackupSetting(v *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) *DescribeHanaBackupSettingResponseBody
	GetHanaBackupSetting() *DescribeHanaBackupSettingResponseBodyHanaBackupSetting
	SetMessage(v string) *DescribeHanaBackupSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHanaBackupSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHanaBackupSettingResponseBody
	GetSuccess() *bool
}

type DescribeHanaBackupSettingResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The backup settings.
	HanaBackupSetting *DescribeHanaBackupSettingResponseBodyHanaBackupSetting `json:"HanaBackupSetting,omitempty" xml:"HanaBackupSetting,omitempty" type:"Struct"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9D0DB5BC-5071-5ADF-BCD1-14EBB0C17C54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHanaBackupSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaBackupSettingResponseBody) GetHanaBackupSetting() *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	return s.HanaBackupSetting
}

func (s *DescribeHanaBackupSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaBackupSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaBackupSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaBackupSettingResponseBody) SetCode(v string) *DescribeHanaBackupSettingResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetHanaBackupSetting(v *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) *DescribeHanaBackupSettingResponseBody {
	s.HanaBackupSetting = v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetMessage(v string) *DescribeHanaBackupSettingResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetRequestId(v string) *DescribeHanaBackupSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetSuccess(v bool) *DescribeHanaBackupSettingResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) Validate() error {
	if s.HanaBackupSetting != nil {
		if err := s.HanaBackupSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHanaBackupSettingResponseBodyHanaBackupSetting struct {
	// The configuration file for catalog backup.
	//
	// example:
	//
	// /usr/sap/SID/SYS/global/hdb/opt/hdbconfig/param
	CatalogBackupParameterFile *string `json:"CatalogBackupParameterFile,omitempty" xml:"CatalogBackupParameterFile,omitempty"`
	// Indicates whether Backint is used to back up catalogs. Valid values:
	//
	// 	- true: Backint is used to back up catalogs.
	//
	// 	- false: Backint is not used to back up catalogs.
	//
	// example:
	//
	// false
	CatalogBackupUsingBackint *bool `json:"CatalogBackupUsingBackint,omitempty" xml:"CatalogBackupUsingBackint,omitempty"`
	// The configuration file for data backup.
	//
	// example:
	//
	// /usr/sap/SID/SYS/global/hdb/opt/hdbconfig/param
	DataBackupParameterFile *string `json:"DataBackupParameterFile,omitempty" xml:"DataBackupParameterFile,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Indicates whether automatic log backup is enabled. Valid values:
	//
	// 	- **true**: Automatic log backup is enabled.
	//
	// 	- **false**: Automatic log backup is disabled.
	//
	// example:
	//
	// true
	EnableAutoLogBackup *bool `json:"EnableAutoLogBackup,omitempty" xml:"EnableAutoLogBackup,omitempty"`
	// The configuration file for log backup.
	//
	// example:
	//
	// /usr/sap/SID/SYS/global/hdb/opt/hdbconfig/param
	LogBackupParameterFile *string `json:"LogBackupParameterFile,omitempty" xml:"LogBackupParameterFile,omitempty"`
	// The interval at which logs are backed up. Unit: seconds.
	//
	// example:
	//
	// 900
	LogBackupTimeout *int64 `json:"LogBackupTimeout,omitempty" xml:"LogBackupTimeout,omitempty"`
	// Indicates whether Backint is used to back up logs. Valid values:
	//
	// 	- true: Backint is used to back up logs.
	//
	// 	- false: Backint is not used to back up logs.
	//
	// example:
	//
	// true
	LogBackupUsingBackint *bool `json:"LogBackupUsingBackint,omitempty" xml:"LogBackupUsingBackint,omitempty"`
}

func (s DescribeHanaBackupSettingResponseBodyHanaBackupSetting) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetCatalogBackupParameterFile() *string {
	return s.CatalogBackupParameterFile
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetCatalogBackupUsingBackint() *bool {
	return s.CatalogBackupUsingBackint
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetDataBackupParameterFile() *string {
	return s.DataBackupParameterFile
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetEnableAutoLogBackup() *bool {
	return s.EnableAutoLogBackup
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetLogBackupParameterFile() *string {
	return s.LogBackupParameterFile
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetLogBackupTimeout() *int64 {
	return s.LogBackupTimeout
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GetLogBackupUsingBackint() *bool {
	return s.LogBackupUsingBackint
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetCatalogBackupParameterFile(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.CatalogBackupParameterFile = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetCatalogBackupUsingBackint(v bool) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.CatalogBackupUsingBackint = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetDataBackupParameterFile(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.DataBackupParameterFile = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetDatabaseName(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetEnableAutoLogBackup(v bool) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.EnableAutoLogBackup = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetLogBackupParameterFile(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.LogBackupParameterFile = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetLogBackupTimeout(v int64) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.LogBackupTimeout = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetLogBackupUsingBackint(v bool) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.LogBackupUsingBackint = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) Validate() error {
	return dara.Validate(s)
}
