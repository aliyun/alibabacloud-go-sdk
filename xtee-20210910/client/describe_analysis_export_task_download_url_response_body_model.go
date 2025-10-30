// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisExportTaskDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAnalysisExportTaskDownloadUrlResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) *DescribeAnalysisExportTaskDownloadUrlResponseBody
	GetResultObject() *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject
}

type DescribeAnalysisExportTaskDownloadUrlResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeAnalysisExportTaskDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisExportTaskDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBody) GetResultObject() *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBody) SetRequestId(v string) *DescribeAnalysisExportTaskDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBody) SetResultObject(v *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) *DescribeAnalysisExportTaskDownloadUrlResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject struct {
	// Download URL.
	//
	// example:
	//
	// https://xxxxx-oss-xxxxx.xxxxxx.aliyuncs.com/xx/xx/xxx/xxxxxx.csv?Expires=1753433384&OSSAccessKeyId=xxxxxxxxx&Signature=%2F%xxxxxxxxxxxx%3D
	DownloadFileUrl *string `json:"downloadFileUrl,omitempty" xml:"downloadFileUrl,omitempty"`
	// Download execution time
	//
	// example:
	//
	// 1753891199000
	ExecuteTime *int64 `json:"executeTime,omitempty" xml:"executeTime,omitempty"`
	// Task status.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) GetDownloadFileUrl() *string {
	return s.DownloadFileUrl
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) SetDownloadFileUrl(v string) *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject {
	s.DownloadFileUrl = &v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) SetExecuteTime(v int64) *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) SetStatus(v string) *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
