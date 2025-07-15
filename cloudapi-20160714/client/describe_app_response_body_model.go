// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *DescribeAppResponseBody
	GetAppId() *int64
	SetAppName(v string) *DescribeAppResponseBody
	GetAppName() *string
	SetCreatedTime(v string) *DescribeAppResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeAppResponseBody
	GetDescription() *string
	SetDisabled(v bool) *DescribeAppResponseBody
	GetDisabled() *bool
	SetExtend(v string) *DescribeAppResponseBody
	GetExtend() *string
	SetModifiedTime(v string) *DescribeAppResponseBody
	GetModifiedTime() *string
	SetRequestId(v string) *DescribeAppResponseBody
	GetRequestId() *string
}

type DescribeAppResponseBody struct {
	// The ID of the app.
	//
	// example:
	//
	// 110843374
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the app.
	//
	// example:
	//
	// CreateApptest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the app was created.
	//
	// example:
	//
	// 2019-01-29T09:33:01Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the app.
	//
	// example:
	//
	// Estimated on October 15, 2021 at 10:20:27
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// 扩展信息
	//
	// example:
	//
	// 110243810311
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The time when the app was modified.
	//
	// example:
	//
	// 2019-01-29T09:33:01Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DBDB3B0F-EC61-5F33-88AD-EC2446FA1DDB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBody) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeAppResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppResponseBody) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeAppResponseBody) GetExtend() *string {
	return s.Extend
}

func (s *DescribeAppResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppResponseBody) SetAppId(v int64) *DescribeAppResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeAppResponseBody) SetAppName(v string) *DescribeAppResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeAppResponseBody) SetCreatedTime(v string) *DescribeAppResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppResponseBody) SetDescription(v string) *DescribeAppResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBody) SetDisabled(v bool) *DescribeAppResponseBody {
	s.Disabled = &v
	return s
}

func (s *DescribeAppResponseBody) SetExtend(v string) *DescribeAppResponseBody {
	s.Extend = &v
	return s
}

func (s *DescribeAppResponseBody) SetModifiedTime(v string) *DescribeAppResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppResponseBody) SetRequestId(v string) *DescribeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppResponseBody) Validate() error {
	return dara.Validate(s)
}
