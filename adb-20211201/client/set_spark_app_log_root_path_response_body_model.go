// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSparkAppLogRootPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SetSparkAppLogRootPathResponseBodyData) *SetSparkAppLogRootPathResponseBody
	GetData() *SetSparkAppLogRootPathResponseBodyData
	SetRequestId(v string) *SetSparkAppLogRootPathResponseBody
	GetRequestId() *string
}

type SetSparkAppLogRootPathResponseBody struct {
	// The returned data.
	Data *SetSparkAppLogRootPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSparkAppLogRootPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSparkAppLogRootPathResponseBody) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathResponseBody) GetData() *SetSparkAppLogRootPathResponseBodyData {
	return s.Data
}

func (s *SetSparkAppLogRootPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSparkAppLogRootPathResponseBody) SetData(v *SetSparkAppLogRootPathResponseBodyData) *SetSparkAppLogRootPathResponseBody {
	s.Data = v
	return s
}

func (s *SetSparkAppLogRootPathResponseBody) SetRequestId(v string) *SetSparkAppLogRootPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSparkAppLogRootPathResponseBodyData struct {
	// The default log path.
	//
	// example:
	//
	// oss://path/to/log
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
	// 1675236908
	ModifiedTimestamp *string `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	// The modifier ID.
	//
	// example:
	//
	// 1111111
	ModifiedUid *string `json:"ModifiedUid,omitempty" xml:"ModifiedUid,omitempty"`
	// The recorded log path.
	//
	// example:
	//
	// oss://path/to/log
	RecordedLogPath *string `json:"RecordedLogPath,omitempty" xml:"RecordedLogPath,omitempty"`
}

func (s SetSparkAppLogRootPathResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetSparkAppLogRootPathResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathResponseBodyData) GetDefaultLogPath() *string {
	return s.DefaultLogPath
}

func (s *SetSparkAppLogRootPathResponseBodyData) GetIsLogPathExists() *bool {
	return s.IsLogPathExists
}

func (s *SetSparkAppLogRootPathResponseBodyData) GetModifiedTimestamp() *string {
	return s.ModifiedTimestamp
}

func (s *SetSparkAppLogRootPathResponseBodyData) GetModifiedUid() *string {
	return s.ModifiedUid
}

func (s *SetSparkAppLogRootPathResponseBodyData) GetRecordedLogPath() *string {
	return s.RecordedLogPath
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetDefaultLogPath(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.DefaultLogPath = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetIsLogPathExists(v bool) *SetSparkAppLogRootPathResponseBodyData {
	s.IsLogPathExists = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetModifiedTimestamp(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.ModifiedTimestamp = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetModifiedUid(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.ModifiedUid = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetRecordedLogPath(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.RecordedLogPath = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) Validate() error {
	return dara.Validate(s)
}
