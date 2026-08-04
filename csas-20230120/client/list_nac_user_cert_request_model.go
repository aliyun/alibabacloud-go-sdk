// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacUserCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *ListNacUserCertRequest
	GetCurrentPage() *string
	SetDepartment(v string) *ListNacUserCertRequest
	GetDepartment() *string
	SetDeviceType(v string) *ListNacUserCertRequest
	GetDeviceType() *string
	SetEndTime(v int64) *ListNacUserCertRequest
	GetEndTime() *int64
	SetPageSize(v string) *ListNacUserCertRequest
	GetPageSize() *string
	SetStartTime(v int64) *ListNacUserCertRequest
	GetStartTime() *int64
	SetStatus(v string) *ListNacUserCertRequest
	GetStatus() *string
	SetUsername(v string) *ListNacUserCertRequest
	GetUsername() *string
}

type ListNacUserCertRequest struct {
	// Page number of the current page in a paged query. Valid values: 1 to 10000.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Department that the user belongs to. The value must be 1 to 128 characters in length. It can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), commas (,), semicolons (;), hyphens (-), underscores (_), forward slashes (/), at signs (@), and spaces.
	//
	// example:
	//
	// 测试部
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// Operating system type of the endpoint device. Valid values:
	//
	// - **windows**: Windows.
	//
	// - **macos**: macOS.
	//
	// - **linux**: Linux.
	//
	// - **android**: Android.
	//
	// - **ios**: iOS.
	//
	// example:
	//
	// windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// End time.
	//
	// example:
	//
	// 1702770400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Number of entries per page for a paged query. Valid values: 1 to 1000.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start time.
	//
	// example:
	//
	// 1702260834
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Certificate status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Username.
	//
	// example:
	//
	// zhang**
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListNacUserCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNacUserCertRequest) GoString() string {
	return s.String()
}

func (s *ListNacUserCertRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *ListNacUserCertRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListNacUserCertRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListNacUserCertRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListNacUserCertRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListNacUserCertRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListNacUserCertRequest) GetStatus() *string {
	return s.Status
}

func (s *ListNacUserCertRequest) GetUsername() *string {
	return s.Username
}

func (s *ListNacUserCertRequest) SetCurrentPage(v string) *ListNacUserCertRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListNacUserCertRequest) SetDepartment(v string) *ListNacUserCertRequest {
	s.Department = &v
	return s
}

func (s *ListNacUserCertRequest) SetDeviceType(v string) *ListNacUserCertRequest {
	s.DeviceType = &v
	return s
}

func (s *ListNacUserCertRequest) SetEndTime(v int64) *ListNacUserCertRequest {
	s.EndTime = &v
	return s
}

func (s *ListNacUserCertRequest) SetPageSize(v string) *ListNacUserCertRequest {
	s.PageSize = &v
	return s
}

func (s *ListNacUserCertRequest) SetStartTime(v int64) *ListNacUserCertRequest {
	s.StartTime = &v
	return s
}

func (s *ListNacUserCertRequest) SetStatus(v string) *ListNacUserCertRequest {
	s.Status = &v
	return s
}

func (s *ListNacUserCertRequest) SetUsername(v string) *ListNacUserCertRequest {
	s.Username = &v
	return s
}

func (s *ListNacUserCertRequest) Validate() error {
	return dara.Validate(s)
}
