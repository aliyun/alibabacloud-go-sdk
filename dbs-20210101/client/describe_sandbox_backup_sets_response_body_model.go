// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxBackupSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSandboxBackupSetsResponseBody
	GetCode() *string
	SetData(v string) *DescribeSandboxBackupSetsResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeSandboxBackupSetsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSandboxBackupSetsResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeSandboxBackupSetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSandboxBackupSetsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSandboxBackupSetsResponseBody
	GetSuccess() *string
}

type DescribeSandboxBackupSetsResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data. The following parameters are contained:
	//
	// 	- **backupSetTime**: the point in time when the snapshot was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- **backupSetId**: the ID of the backup set.
	//
	// 	- **backupSetType**: the type of the snapshot. A value of **Full*	- indicates that the snapshot is a full backup snapshot. A value of **Inc*	- indicates that the snapshot is an incremental backup snapshot.
	//
	// 	- **backupPlanId**: the ID of the backup schedule.
	//
	// example:
	//
	// "Data": {     "number": 2,     "size": 2,     "content": [       {         "backupSetTime": "2021-08-28T23:12:31Z",         "backupSetId": "Inc_1hxxxx8xxxxxa_20210801064200_mysql-bin.000134",         "backupSetType": "Inc",         "backupPlanId": "1hxxxx8xxxxxa"       },       {         "backupSetTime": "2021-08-28T22:42:28Z",         "backupSetId": "1hxxxx8xxxxxa_20210829064228",         "backupSetType": "FULL",         "backupPlanId": "1hxxxx8xxxxxa"       }     ],     "totalElements": 2   },
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxBackupSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxBackupSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSandboxBackupSetsResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeSandboxBackupSetsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSandboxBackupSetsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSandboxBackupSetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSandboxBackupSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSandboxBackupSetsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSandboxBackupSetsResponseBody) SetCode(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetData(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetErrCode(v string) *DescribeSandboxBackupSetsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetErrMessage(v string) *DescribeSandboxBackupSetsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetMessage(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetRequestId(v string) *DescribeSandboxBackupSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetSuccess(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) Validate() error {
	return dara.Validate(s)
}
