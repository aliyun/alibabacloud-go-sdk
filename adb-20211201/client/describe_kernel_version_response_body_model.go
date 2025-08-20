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
	// The minor versions to which you can update the current minor version of the cluster.
	AvailableKernelVersions []*DescribeKernelVersionResponseBodyAvailableKernelVersions `json:"AvailableKernelVersions,omitempty" xml:"AvailableKernelVersions,omitempty" type:"Repeated"`
	// The maintenance expiration time of the version. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. After the time arrives, the system no longer maintains the version. If any issues occur, update the minor version of the cluster to a later version.
	//
	// example:
	//
	// 2025-01-12T16:00:00Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The minor version of the cluster. Example: **3.1.8**.
	//
	// example:
	//
	// 3.1.8
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The request ID.
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
	return dara.Validate(s)
}

type DescribeKernelVersionResponseBodyAvailableKernelVersions struct {
	// The maintenance expiration time of the version. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. After the time arrives, the system no longer maintains the version. If any issues occur, update the minor version of the cluster to a later version.
	//
	// example:
	//
	// 2025-07-17T16:00:00Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The minor version. Example: **3.1.9**.
	//
	// example:
	//
	// 3.1.9
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The time when the minor version was released. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
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
