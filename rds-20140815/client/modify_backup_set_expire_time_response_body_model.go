// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSetExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyBackupSetExpireTimeResponseBody
	GetCode() *string
	SetData(v string) *ModifyBackupSetExpireTimeResponseBody
	GetData() *string
	SetMessage(v string) *ModifyBackupSetExpireTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBackupSetExpireTimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBackupSetExpireTimeResponseBody
	GetSuccess() *bool
}

type ModifyBackupSetExpireTimeResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the returned parameters.
	//
	// example:
	//
	// {
	//
	//       "SupportOnlineResizeDisk": true,
	//
	//       "DBInstanceName": "rm-bp****",
	//
	//       "maxSupportDiskSizeGB": 6144
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The information about the status code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 76364A52-E0AB-5CC8-9818-CF1DC482C092
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBackupSetExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSetExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetExpireTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBackupSetExpireTimeResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyBackupSetExpireTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBackupSetExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupSetExpireTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBackupSetExpireTimeResponseBody) SetCode(v string) *ModifyBackupSetExpireTimeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBackupSetExpireTimeResponseBody) SetData(v string) *ModifyBackupSetExpireTimeResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyBackupSetExpireTimeResponseBody) SetMessage(v string) *ModifyBackupSetExpireTimeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBackupSetExpireTimeResponseBody) SetRequestId(v string) *ModifyBackupSetExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupSetExpireTimeResponseBody) SetSuccess(v bool) *ModifyBackupSetExpireTimeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupSetExpireTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
