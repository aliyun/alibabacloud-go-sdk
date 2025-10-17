// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAppDiagnosisInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSparkAppDiagnosisInfoResponseBody
	GetAccessDeniedDetail() *string
	SetAppId(v string) *DescribeSparkAppDiagnosisInfoResponseBody
	GetAppId() *string
	SetCpuUtilization(v float64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetCpuUtilization() *float64
	SetDiagnosisInfoList(v []*Adb4MysqlSparkDiagnosisInfo) *DescribeSparkAppDiagnosisInfoResponseBody
	GetDiagnosisInfoList() []*Adb4MysqlSparkDiagnosisInfo
	SetDurationInMillis(v int64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetDurationInMillis() *int64
	SetJVMGcCostInMillis(v int64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetJVMGcCostInMillis() *int64
	SetPeakMemoryInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetPeakMemoryInByte() *int64
	SetRequestId(v string) *DescribeSparkAppDiagnosisInfoResponseBody
	GetRequestId() *string
	SetShuffleReadInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetShuffleReadInByte() *int64
	SetShuffleWriteInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetShuffleWriteInByte() *int64
	SetSpillInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetSpillInByte() *int64
	SetStartedTime(v int64) *DescribeSparkAppDiagnosisInfoResponseBody
	GetStartedTime() *int64
	SetState(v string) *DescribeSparkAppDiagnosisInfoResponseBody
	GetState() *string
}

type DescribeSparkAppDiagnosisInfoResponseBody struct {
	// The information about the request denial.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query all application IDs.
	//
	// example:
	//
	// s202404141952sz6a1391200****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The CPU utilization. Unit: %.
	//
	// example:
	//
	// 100
	CpuUtilization *float64 `json:"CpuUtilization,omitempty" xml:"CpuUtilization,omitempty"`
	// The queried diagnostic information.
	DiagnosisInfoList []*Adb4MysqlSparkDiagnosisInfo `json:"DiagnosisInfoList,omitempty" xml:"DiagnosisInfoList,omitempty" type:"Repeated"`
	// The execution duration of the application. Unit: milliseconds.
	//
	// example:
	//
	// 281063
	DurationInMillis *int64 `json:"DurationInMillis,omitempty" xml:"DurationInMillis,omitempty"`
	// The amount of time consumed by the Java virtual machine to perform garbage collection operations. Unit: milliseconds.
	//
	// example:
	//
	// 81055
	JVMGcCostInMillis *int64 `json:"JVMGcCostInMillis,omitempty" xml:"JVMGcCostInMillis,omitempty"`
	// The peak memory usage. Unit: bytes.
	//
	// example:
	//
	// 4096000
	PeakMemoryInByte *int64 `json:"PeakMemoryInByte,omitempty" xml:"PeakMemoryInByte,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FB5AC339-91F6-5000-8E5A-F47065B01B87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The amount of data used for shuffle reads. Unit: bytes.
	//
	// example:
	//
	// 4096000
	ShuffleReadInByte *int64 `json:"ShuffleReadInByte,omitempty" xml:"ShuffleReadInByte,omitempty"`
	// The amount of data used for shuffle writes. Unit: bytes.
	//
	// example:
	//
	// 4096000
	ShuffleWriteInByte *int64 `json:"ShuffleWriteInByte,omitempty" xml:"ShuffleWriteInByte,omitempty"`
	// The amount of data spilled to disks when the memory is insufficient. Unit: bytes.
	//
	// example:
	//
	// 0
	SpillInByte *int64 `json:"SpillInByte,omitempty" xml:"SpillInByte,omitempty"`
	// The time when the application started to be executed.
	//
	// example:
	//
	// 1718329831000
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the asynchronous import or export job. Valid values:
	//
	// 	- **RUNNING**
	//
	// 	- **FINISHED**
	//
	// 	- **FAILED**
	//
	// example:
	//
	// FINISHED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSparkAppDiagnosisInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAppDiagnosisInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetCpuUtilization() *float64 {
	return s.CpuUtilization
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetDiagnosisInfoList() []*Adb4MysqlSparkDiagnosisInfo {
	return s.DiagnosisInfoList
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetDurationInMillis() *int64 {
	return s.DurationInMillis
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetJVMGcCostInMillis() *int64 {
	return s.JVMGcCostInMillis
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetPeakMemoryInByte() *int64 {
	return s.PeakMemoryInByte
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetShuffleReadInByte() *int64 {
	return s.ShuffleReadInByte
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetShuffleWriteInByte() *int64 {
	return s.ShuffleWriteInByte
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetSpillInByte() *int64 {
	return s.SpillInByte
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetAccessDeniedDetail(v string) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetAppId(v string) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetCpuUtilization(v float64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.CpuUtilization = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetDiagnosisInfoList(v []*Adb4MysqlSparkDiagnosisInfo) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.DiagnosisInfoList = v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetDurationInMillis(v int64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.DurationInMillis = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetJVMGcCostInMillis(v int64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.JVMGcCostInMillis = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetPeakMemoryInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.PeakMemoryInByte = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetRequestId(v string) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetShuffleReadInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.ShuffleReadInByte = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetShuffleWriteInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.ShuffleWriteInByte = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetSpillInByte(v int64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.SpillInByte = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetStartedTime(v int64) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.StartedTime = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) SetState(v string) *DescribeSparkAppDiagnosisInfoResponseBody {
	s.State = &v
	return s
}

func (s *DescribeSparkAppDiagnosisInfoResponseBody) Validate() error {
	if s.DiagnosisInfoList != nil {
		for _, item := range s.DiagnosisInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
