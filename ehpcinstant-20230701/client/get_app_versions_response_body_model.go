// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersions(v []*GetAppVersionsResponseBodyAppVersions) *GetAppVersionsResponseBody
	GetAppVersions() []*GetAppVersionsResponseBodyAppVersions
	SetPageNumber(v int64) *GetAppVersionsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetAppVersionsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetAppVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppVersionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetAppVersionsResponseBody
	GetTotalCount() *int32
}

type GetAppVersionsResponseBody struct {
	AppVersions []*GetAppVersionsResponseBodyAppVersions `json:"AppVersions,omitempty" xml:"AppVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAppVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppVersionsResponseBody) GetAppVersions() []*GetAppVersionsResponseBodyAppVersions {
	return s.AppVersions
}

func (s *GetAppVersionsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetAppVersionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAppVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAppVersionsResponseBody) SetAppVersions(v []*GetAppVersionsResponseBodyAppVersions) *GetAppVersionsResponseBody {
	s.AppVersions = v
	return s
}

func (s *GetAppVersionsResponseBody) SetPageNumber(v int64) *GetAppVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetPageSize(v int64) *GetAppVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetRequestId(v string) *GetAppVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetSuccess(v bool) *GetAppVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetTotalCount(v int32) *GetAppVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetAppVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAppVersionsResponseBodyAppVersions struct {
	// This parameter is required.
	//
	// example:
	//
	// m-f8z0dfa96luomqly****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// V-Ray
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// v1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAppVersionsResponseBodyAppVersions) String() string {
	return dara.Prettify(s)
}

func (s GetAppVersionsResponseBodyAppVersions) GoString() string {
	return s.String()
}

func (s *GetAppVersionsResponseBodyAppVersions) GetImageId() *string {
	return s.ImageId
}

func (s *GetAppVersionsResponseBodyAppVersions) GetName() *string {
	return s.Name
}

func (s *GetAppVersionsResponseBodyAppVersions) GetVersion() *string {
	return s.Version
}

func (s *GetAppVersionsResponseBodyAppVersions) SetImageId(v string) *GetAppVersionsResponseBodyAppVersions {
	s.ImageId = &v
	return s
}

func (s *GetAppVersionsResponseBodyAppVersions) SetName(v string) *GetAppVersionsResponseBodyAppVersions {
	s.Name = &v
	return s
}

func (s *GetAppVersionsResponseBodyAppVersions) SetVersion(v string) *GetAppVersionsResponseBodyAppVersions {
	s.Version = &v
	return s
}

func (s *GetAppVersionsResponseBodyAppVersions) Validate() error {
	return dara.Validate(s)
}
