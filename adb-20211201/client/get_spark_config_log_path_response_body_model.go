// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkConfigLogPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkConfigLogPathResponseBodyData) *GetSparkConfigLogPathResponseBody
	GetData() *GetSparkConfigLogPathResponseBodyData
	SetRequestId(v string) *GetSparkConfigLogPathResponseBody
	GetRequestId() *string
}

type GetSparkConfigLogPathResponseBody struct {
	// The queried Spark log configuration.
	Data *GetSparkConfigLogPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1919-xxx-ssdfsdff
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkConfigLogPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkConfigLogPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathResponseBody) GetData() *GetSparkConfigLogPathResponseBodyData {
	return s.Data
}

func (s *GetSparkConfigLogPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkConfigLogPathResponseBody) SetData(v *GetSparkConfigLogPathResponseBodyData) *GetSparkConfigLogPathResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkConfigLogPathResponseBody) SetRequestId(v string) *GetSparkConfigLogPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSparkConfigLogPathResponseBodyData struct {
	// The default log path.
	//
	// example:
	//
	// oss://aliyun-oa-adb-spark-1111-oss-cn-hanghzou/spark-logs
	DefaultLogPath *string `json:"DefaultLogPath,omitempty" xml:"DefaultLogPath,omitempty"`
	// Indicates whether a log path exists.
	//
	// example:
	//
	// true
	IsLogPathExists *bool `json:"IsLogPathExists,omitempty" xml:"IsLogPathExists,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1675654361000
	ModifiedTimestamp *string `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	// The account ID of the modifier.
	//
	// example:
	//
	// 10130223128xxx
	ModifiedUid *string `json:"ModifiedUid,omitempty" xml:"ModifiedUid,omitempty"`
	// The recorded log path.
	//
	// example:
	//
	// oss://test/spark-logs/
	RecordedLogPath *string `json:"RecordedLogPath,omitempty" xml:"RecordedLogPath,omitempty"`
}

func (s GetSparkConfigLogPathResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkConfigLogPathResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathResponseBodyData) GetDefaultLogPath() *string {
	return s.DefaultLogPath
}

func (s *GetSparkConfigLogPathResponseBodyData) GetIsLogPathExists() *bool {
	return s.IsLogPathExists
}

func (s *GetSparkConfigLogPathResponseBodyData) GetModifiedTimestamp() *string {
	return s.ModifiedTimestamp
}

func (s *GetSparkConfigLogPathResponseBodyData) GetModifiedUid() *string {
	return s.ModifiedUid
}

func (s *GetSparkConfigLogPathResponseBodyData) GetRecordedLogPath() *string {
	return s.RecordedLogPath
}

func (s *GetSparkConfigLogPathResponseBodyData) SetDefaultLogPath(v string) *GetSparkConfigLogPathResponseBodyData {
	s.DefaultLogPath = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetIsLogPathExists(v bool) *GetSparkConfigLogPathResponseBodyData {
	s.IsLogPathExists = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetModifiedTimestamp(v string) *GetSparkConfigLogPathResponseBodyData {
	s.ModifiedTimestamp = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetModifiedUid(v string) *GetSparkConfigLogPathResponseBodyData {
	s.ModifiedUid = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetRecordedLogPath(v string) *GetSparkConfigLogPathResponseBodyData {
	s.RecordedLogPath = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) Validate() error {
	return dara.Validate(s)
}
