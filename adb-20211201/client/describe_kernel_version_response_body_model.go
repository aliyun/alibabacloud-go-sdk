// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKernelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableKernelVersions(v []*DescribeKernelVersionResponseBodyAvailableKernelVersions) *DescribeKernelVersionResponseBody
	GetAvailableKernelVersions() []*DescribeKernelVersionResponseBodyAvailableKernelVersions
	SetExpireDate(v string) *DescribeKernelVersionResponseBody
	GetExpireDate() *string
	SetKernelVersion(v string) *DescribeKernelVersionResponseBody
	GetKernelVersion() *string
	SetRequestId(v string) *DescribeKernelVersionResponseBody
	GetRequestId() *string
}

type DescribeKernelVersionResponseBody struct {
	// The list of target versions to which the current version can be upgraded.
	AvailableKernelVersions []*DescribeKernelVersionResponseBodyAvailableKernelVersions `json:"AvailableKernelVersions,omitempty" xml:"AvailableKernelVersions,omitempty" type:"Repeated"`
	// The end-of-support date of the version. The time is in the yyyy-MM-DD HH:mm:ss format. After this date, the system no longer maintains this version. If any issues occur, upgrade to a later version.
	//
	// example:
	//
	// 2025-01-12T16:00:00Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The kernel version number of the cluster. Example: **3.1.8**.
	//
	// example:
	//
	// 3.1.8
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKernelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKernelVersionResponseBody) GetAvailableKernelVersions() []*DescribeKernelVersionResponseBodyAvailableKernelVersions {
	return s.AvailableKernelVersions
}

func (s *DescribeKernelVersionResponseBody) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *DescribeKernelVersionResponseBody) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *DescribeKernelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKernelVersionResponseBody) SetAvailableKernelVersions(v []*DescribeKernelVersionResponseBodyAvailableKernelVersions) *DescribeKernelVersionResponseBody {
	s.AvailableKernelVersions = v
	return s
}

func (s *DescribeKernelVersionResponseBody) SetExpireDate(v string) *DescribeKernelVersionResponseBody {
	s.ExpireDate = &v
	return s
}

func (s *DescribeKernelVersionResponseBody) SetKernelVersion(v string) *DescribeKernelVersionResponseBody {
	s.KernelVersion = &v
	return s
}

func (s *DescribeKernelVersionResponseBody) SetRequestId(v string) *DescribeKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKernelVersionResponseBody) Validate() error {
	if s.AvailableKernelVersions != nil {
		for _, item := range s.AvailableKernelVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKernelVersionResponseBodyAvailableKernelVersions struct {
	// The end-of-support date of the version. The time is in the yyyy-MM-DD HH:mm:ss format. After this date, the system no longer maintains this version. If any issues occur, upgrade to a later version.
	//
	// example:
	//
	// 2025-07-17T16:00:00Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The kernel version number of the cluster. Example: **3.1.9**.
	//
	// example:
	//
	// 3.1.9
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The release time of the kernel version. The time format is yyyy-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-07-17T16:00:00Z
	ReleaseDate *string `json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
}

func (s DescribeKernelVersionResponseBodyAvailableKernelVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelVersionResponseBodyAvailableKernelVersions) GoString() string {
	return s.String()
}

func (s *DescribeKernelVersionResponseBodyAvailableKernelVersions) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *DescribeKernelVersionResponseBodyAvailableKernelVersions) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *DescribeKernelVersionResponseBodyAvailableKernelVersions) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *DescribeKernelVersionResponseBodyAvailableKernelVersions) SetExpireDate(v string) *DescribeKernelVersionResponseBodyAvailableKernelVersions {
	s.ExpireDate = &v
	return s
}

func (s *DescribeKernelVersionResponseBodyAvailableKernelVersions) SetKernelVersion(v string) *DescribeKernelVersionResponseBodyAvailableKernelVersions {
	s.KernelVersion = &v
	return s
}

func (s *DescribeKernelVersionResponseBodyAvailableKernelVersions) SetReleaseDate(v string) *DescribeKernelVersionResponseBodyAvailableKernelVersions {
	s.ReleaseDate = &v
	return s
}

func (s *DescribeKernelVersionResponseBodyAvailableKernelVersions) Validate() error {
	return dara.Validate(s)
}
