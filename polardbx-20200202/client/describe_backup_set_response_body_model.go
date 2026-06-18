// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeBackupSetResponseBodyData) *DescribeBackupSetResponseBody
	GetData() []*DescribeBackupSetResponseBodyData
	SetMessage(v string) *DescribeBackupSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBackupSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupSetResponseBody
	GetSuccess() *bool
}

type DescribeBackupSetResponseBody struct {
	// The data list.
	Data []*DescribeBackupSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// successs
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A6D328C-84B8-40DC-BF49-6C73984D7494
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponseBody) GetData() []*DescribeBackupSetResponseBodyData {
	return s.Data
}

func (s *DescribeBackupSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupSetResponseBody) SetData(v []*DescribeBackupSetResponseBodyData) *DescribeBackupSetResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupSetResponseBody) SetMessage(v string) *DescribeBackupSetResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupSetResponseBody) SetRequestId(v string) *DescribeBackupSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetResponseBody) SetSuccess(v bool) *DescribeBackupSetResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupSetResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupSetResponseBodyData struct {
	// The backup mode. Valid values:
	//
	// - **0**: Automatic backup.
	//
	// - **1**: Manual backup.
	//
	// example:
	//
	// 0
	BackupModel *int32 `json:"BackupModel,omitempty" xml:"BackupModel,omitempty"`
	// The backup set ID.
	//
	// example:
	//
	// 111
	BackupSetId *int64 `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The size of the backup set, in bytes.
	//
	// example:
	//
	// 88803195
	BackupSetSize *int64 `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	// The backup type. Valid values:
	//
	// - **0**: Fast backup.
	//
	// - **1**: Consistent backup.
	//
	// example:
	//
	// 1
	BackupType *int32 `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The time when the backup started.
	//
	// example:
	//
	// 1650250861754
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The time when the backup ended.
	//
	// example:
	//
	// 1650251308000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of OSS information for the backup set.
	OSSList []*DescribeBackupSetResponseBodyDataOSSList `json:"OSSList,omitempty" xml:"OSSList,omitempty" type:"Repeated"`
	// The status of the backup set. Valid values:
	//
	// - **0**: Backing up.
	//
	// - **1**: Backup succeeded.
	//
	// - **2**: Backup failed.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponseBodyData) GetBackupModel() *int32 {
	return s.BackupModel
}

func (s *DescribeBackupSetResponseBodyData) GetBackupSetId() *int64 {
	return s.BackupSetId
}

func (s *DescribeBackupSetResponseBodyData) GetBackupSetSize() *int64 {
	return s.BackupSetSize
}

func (s *DescribeBackupSetResponseBodyData) GetBackupType() *int32 {
	return s.BackupType
}

func (s *DescribeBackupSetResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeBackupSetResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeBackupSetResponseBodyData) GetOSSList() []*DescribeBackupSetResponseBodyDataOSSList {
	return s.OSSList
}

func (s *DescribeBackupSetResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeBackupSetResponseBodyData) SetBackupModel(v int32) *DescribeBackupSetResponseBodyData {
	s.BackupModel = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBackupSetId(v int64) *DescribeBackupSetResponseBodyData {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBackupSetSize(v int64) *DescribeBackupSetResponseBodyData {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBackupType(v int32) *DescribeBackupSetResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetBeginTime(v int64) *DescribeBackupSetResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetEndTime(v int64) *DescribeBackupSetResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetOSSList(v []*DescribeBackupSetResponseBodyDataOSSList) *DescribeBackupSetResponseBodyData {
	s.OSSList = v
	return s
}

func (s *DescribeBackupSetResponseBodyData) SetStatus(v int32) *DescribeBackupSetResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeBackupSetResponseBodyData) Validate() error {
	if s.OSSList != nil {
		for _, item := range s.OSSList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupSetResponseBodyDataOSSList struct {
	// The OSS file name.
	//
	// example:
	//
	// hins3084_data_20220418110623_qp.xb
	BackupSetFile *string `json:"BackupSetFile,omitempty" xml:"BackupSetFile,omitempty"`
	// The public download URL.
	//
	// example:
	//
	// https://pre-rdsbak-cn-xxx.oss-cn-beijing.aliyuncs.com/custins2255/hins3084_data_20220418110623_qp.xb?OSSAccessKeyId=LTAI****************&Expires=1650441697&dd=7KJzkUSbXf6dwy
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The internal network download URL for the OSS file.
	//
	// example:
	//
	// http://pre-rdsbak-cn-beijing.oss-cn-beijing-internal.aliyuncs.com/custins2255/hins3084_data_20220418110623_qp.xb?
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// The expiration time of the OSS file link.
	//
	// example:
	//
	// 2022-04-20T08:01:37Z
	LinkExpiredTime *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
}

func (s DescribeBackupSetResponseBodyDataOSSList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetResponseBodyDataOSSList) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponseBodyDataOSSList) GetBackupSetFile() *string {
	return s.BackupSetFile
}

func (s *DescribeBackupSetResponseBodyDataOSSList) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeBackupSetResponseBodyDataOSSList) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *DescribeBackupSetResponseBodyDataOSSList) GetLinkExpiredTime() *string {
	return s.LinkExpiredTime
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetBackupSetFile(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.BackupSetFile = &v
	return s
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetDownloadLink(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetIntranetDownloadLink(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBackupSetResponseBodyDataOSSList) SetLinkExpiredTime(v string) *DescribeBackupSetResponseBodyDataOSSList {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeBackupSetResponseBodyDataOSSList) Validate() error {
	return dara.Validate(s)
}
