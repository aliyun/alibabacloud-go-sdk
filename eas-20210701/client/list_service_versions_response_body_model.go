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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 166
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The historical versions of the service.
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
	return dara.Validate(s)
}

type ListServiceVersionsResponseBodyVersions struct {
	// The time when the service version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-29T22:02:14Z
	BuildTime *string `json:"BuildTime,omitempty" xml:"BuildTime,omitempty"`
	// Indicates whether the image is available. Valid values:
	//
	// 	- true: The image is available.
	//
	// 	- false: The image is unavailable.
	//
	// 	- unknown: The availability of the image is unknown.
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
	// The returned message.
	//
	// example:
	//
	// Service is Running
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The service deployment configurations. This parameter is returned only if the service is deployed by using a custom image.
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
	// Indicates whether Elastic Algorithm service (EAS) is activated. Valid values:
	//
	// 	- true: EAS is activated.
	//
	// 	- false: EAS is not activated.
	//
	// 	- unknown: The activation of EAS is unknown.
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
