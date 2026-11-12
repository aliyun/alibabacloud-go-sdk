// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListServiceVersionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServiceVersionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListServiceVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListServiceVersionsResponseBody
	GetTotalCount() *int64
	SetVersions(v []*ListServiceVersionsResponseBodyVersions) *ListServiceVersionsResponseBody
	GetVersions() []*ListServiceVersionsResponseBodyVersions
}

type ListServiceVersionsResponseBody struct {
	// The current page number of the returned historical versions.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of historical versions returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E089D584-B6F4-50C4-9902-DA2295B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of versions.
	//
	// example:
	//
	// 166
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of versions.
	Versions []*ListServiceVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListServiceVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServiceVersionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListServiceVersionsResponseBody) GetVersions() []*ListServiceVersionsResponseBodyVersions {
	return s.Versions
}

func (s *ListServiceVersionsResponseBody) SetPageNumber(v int32) *ListServiceVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetPageSize(v int32) *ListServiceVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetRequestId(v string) *ListServiceVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetTotalCount(v int64) *ListServiceVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetVersions(v []*ListServiceVersionsResponseBodyVersions) *ListServiceVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListServiceVersionsResponseBody) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceVersionsResponseBodyVersions struct {
	// The UTC time when the version was created.
	//
	// example:
	//
	// 2022-08-29T22:02:14Z
	BuildTime *string `json:"BuildTime,omitempty" xml:"BuildTime,omitempty"`
	// The availability of the image. Valid values:
	//
	// - true: Available.
	//
	// - false: Unavailable.
	//
	// - unknown: Unknown.
	//
	// example:
	//
	// true
	ImageAvailable *string `json:"ImageAvailable,omitempty" xml:"ImageAvailable,omitempty"`
	// The image ID.
	//
	// example:
	//
	// 4
	ImageId *int32 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The summary of the service.
	//
	// example:
	//
	// Service is Running
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The remarks.
	//
	// example:
	//
	// new version.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The deployment configuration of the service. This field is returned only for custom image deployments.
	//
	// example:
	//
	// {
	//
	//     "metadata": {
	//
	//         "cpu": 1,
	//
	//         "instance": 1,
	//
	//         "memory": 1024
	//
	//     },
	//
	//     "name": "echo"
	//
	// }
	ServiceConfig *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// Indicates whether the EAS service is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// - unknown: Unknown.
	//
	// example:
	//
	// true
	ServiceRunnable *string `json:"ServiceRunnable,omitempty" xml:"ServiceRunnable,omitempty"`
}

func (s ListServiceVersionsResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListServiceVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponseBodyVersions) GetBuildTime() *string {
	return s.BuildTime
}

func (s *ListServiceVersionsResponseBodyVersions) GetImageAvailable() *string {
	return s.ImageAvailable
}

func (s *ListServiceVersionsResponseBodyVersions) GetImageId() *int32 {
	return s.ImageId
}

func (s *ListServiceVersionsResponseBodyVersions) GetMessage() *string {
	return s.Message
}

func (s *ListServiceVersionsResponseBodyVersions) GetRemark() *string {
	return s.Remark
}

func (s *ListServiceVersionsResponseBodyVersions) GetServiceConfig() *string {
	return s.ServiceConfig
}

func (s *ListServiceVersionsResponseBodyVersions) GetServiceRunnable() *string {
	return s.ServiceRunnable
}

func (s *ListServiceVersionsResponseBodyVersions) SetBuildTime(v string) *ListServiceVersionsResponseBodyVersions {
	s.BuildTime = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetImageAvailable(v string) *ListServiceVersionsResponseBodyVersions {
	s.ImageAvailable = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetImageId(v int32) *ListServiceVersionsResponseBodyVersions {
	s.ImageId = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetMessage(v string) *ListServiceVersionsResponseBodyVersions {
	s.Message = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetRemark(v string) *ListServiceVersionsResponseBodyVersions {
	s.Remark = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetServiceConfig(v string) *ListServiceVersionsResponseBodyVersions {
	s.ServiceConfig = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetServiceRunnable(v string) *ListServiceVersionsResponseBodyVersions {
	s.ServiceRunnable = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
