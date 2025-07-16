// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenBackupSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DescribeOpenBackupSetResponseBody
	GetData() interface{}
	SetRequestId(v string) *DescribeOpenBackupSetResponseBody
	GetRequestId() *string
}

type DescribeOpenBackupSetResponseBody struct {
	// example:
	//
	// {"gmsBackupSet": {"pubFullDownloadUrl": "https://xxxxx","dnName": "pxc-xdb-m-xxxxxx","hostInstanceId": 0001,"binlogs": [],"backupEndTime": "2024-10-21T10:11:56Z","backupLinkExpiredTime": "2024-10-23T06:13:54Z","dnBackupSetId": "00088","notCompletedBinlogs": [],"commitIndex": "15249275","innerFullDownloadUrl": "http://xxxxx","backupStartTime": "2024-10-21T10:09:20Z","backupSetSize": 526118912},"dnBackupSets": [],"insName": "pxc-xxxxx","backupSetId": "cb-xxxxx","canBinlogRecoverToTime": 1729567925000,"backupEndTime": "2024-10-21T10:12:16Z","canBinlogRecoverToTimeUTC": "2024-10-22T03:32:05Z","canBackupMinRecoverToTimeUTC": "2024-10-21T10:11:56Z","pitrInvalid": false,"backupStartTime": "2024-10-21T10:09:16Z","canBackupMinRecoverToTime": 1729505516000}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenBackupSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenBackupSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenBackupSetResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DescribeOpenBackupSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenBackupSetResponseBody) SetData(v interface{}) *DescribeOpenBackupSetResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenBackupSetResponseBody) SetRequestId(v string) *DescribeOpenBackupSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenBackupSetResponseBody) Validate() error {
	return dara.Validate(s)
}
