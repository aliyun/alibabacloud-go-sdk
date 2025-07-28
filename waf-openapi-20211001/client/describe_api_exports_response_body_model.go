// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiExportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiExports(v []*DescribeApiExportsResponseBodyApiExports) *DescribeApiExportsResponseBody
	GetApiExports() []*DescribeApiExportsResponseBodyApiExports
	SetRequestId(v string) *DescribeApiExportsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeApiExportsResponseBody
	GetTotal() *int64
}

type DescribeApiExportsResponseBody struct {
	// The returned data export tasks.
	ApiExports []*DescribeApiExportsResponseBodyApiExports `json:"ApiExports,omitempty" xml:"ApiExports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D9532525-E885-54E7-A178-D5554D563AFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the data export task. Valid values:
	//
	// 	- **expired**: The file is expired.
	//
	// 	- **exporting**: Data is being exported.
	//
	// 	- **completed**: Data is exported.
	//
	// example:
	//
	// 7
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeApiExportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiExportsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiExportsResponseBody) GetApiExports() []*DescribeApiExportsResponseBodyApiExports {
	return s.ApiExports
}

func (s *DescribeApiExportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiExportsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeApiExportsResponseBody) SetApiExports(v []*DescribeApiExportsResponseBodyApiExports) *DescribeApiExportsResponseBody {
	s.ApiExports = v
	return s
}

func (s *DescribeApiExportsResponseBody) SetRequestId(v string) *DescribeApiExportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiExportsResponseBody) SetTotal(v int64) *DescribeApiExportsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeApiExportsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiExportsResponseBodyApiExports struct {
	// The time when the data export task was created. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1725604852
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// file_16109541456445334c0f01d9a7444e0e908***.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The download URL of the exported file.
	//
	// example:
	//
	// https://waf-api-sec-cn.***.aliyuncs.com/file_1610954145***.csv
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The format of the exported file.
	//
	// example:
	//
	// CSV
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The status of the data export task. Valid values:
	//
	// 	- **expired**: The file is expired.
	//
	// 	- **exporting**: Data is being exported.
	//
	// 	- **completed**: Data is exported.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the data export task. Valid values:
	//
	// 	- **apisec_api**: API tasks
	//
	// 	- **apisec_abnormal**: API risk tasks
	//
	// 	- **apisec_event**: API security event tasks
	//
	// example:
	//
	// apisec_api
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApiExportsResponseBodyApiExports) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiExportsResponseBodyApiExports) GoString() string {
	return s.String()
}

func (s *DescribeApiExportsResponseBodyApiExports) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeApiExportsResponseBodyApiExports) GetFileName() *string {
	return s.FileName
}

func (s *DescribeApiExportsResponseBodyApiExports) GetFileUrl() *string {
	return s.FileUrl
}

func (s *DescribeApiExportsResponseBodyApiExports) GetFormat() *string {
	return s.Format
}

func (s *DescribeApiExportsResponseBodyApiExports) GetStatus() *string {
	return s.Status
}

func (s *DescribeApiExportsResponseBodyApiExports) GetType() *string {
	return s.Type
}

func (s *DescribeApiExportsResponseBodyApiExports) SetCreateTime(v int64) *DescribeApiExportsResponseBodyApiExports {
	s.CreateTime = &v
	return s
}

func (s *DescribeApiExportsResponseBodyApiExports) SetFileName(v string) *DescribeApiExportsResponseBodyApiExports {
	s.FileName = &v
	return s
}

func (s *DescribeApiExportsResponseBodyApiExports) SetFileUrl(v string) *DescribeApiExportsResponseBodyApiExports {
	s.FileUrl = &v
	return s
}

func (s *DescribeApiExportsResponseBodyApiExports) SetFormat(v string) *DescribeApiExportsResponseBodyApiExports {
	s.Format = &v
	return s
}

func (s *DescribeApiExportsResponseBodyApiExports) SetStatus(v string) *DescribeApiExportsResponseBodyApiExports {
	s.Status = &v
	return s
}

func (s *DescribeApiExportsResponseBodyApiExports) SetType(v string) *DescribeApiExportsResponseBodyApiExports {
	s.Type = &v
	return s
}

func (s *DescribeApiExportsResponseBodyApiExports) Validate() error {
	return dara.Validate(s)
}
