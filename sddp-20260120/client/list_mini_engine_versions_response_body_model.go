// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMiniEngineVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKernelVersions(v []*ListMiniEngineVersionsResponseBodyKernelVersions) *ListMiniEngineVersionsResponseBody
	GetKernelVersions() []*ListMiniEngineVersionsResponseBodyKernelVersions
	SetRequestId(v string) *ListMiniEngineVersionsResponseBody
	GetRequestId() *string
}

type ListMiniEngineVersionsResponseBody struct {
	KernelVersions []*ListMiniEngineVersionsResponseBodyKernelVersions `json:"KernelVersions,omitempty" xml:"KernelVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMiniEngineVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMiniEngineVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMiniEngineVersionsResponseBody) GetKernelVersions() []*ListMiniEngineVersionsResponseBodyKernelVersions {
	return s.KernelVersions
}

func (s *ListMiniEngineVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMiniEngineVersionsResponseBody) SetKernelVersions(v []*ListMiniEngineVersionsResponseBodyKernelVersions) *ListMiniEngineVersionsResponseBody {
	s.KernelVersions = v
	return s
}

func (s *ListMiniEngineVersionsResponseBody) SetRequestId(v string) *ListMiniEngineVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMiniEngineVersionsResponseBody) Validate() error {
	if s.KernelVersions != nil {
		for _, item := range s.KernelVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMiniEngineVersionsResponseBodyKernelVersions struct {
	// example:
	//
	// LTS
	KernelReleaseType *string `json:"KernelReleaseType,omitempty" xml:"KernelReleaseType,omitempty"`
	// example:
	//
	// rds_20220731
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// example:
	//
	// MySQL 8.0
	KernelVersionName *string `json:"KernelVersionName,omitempty" xml:"KernelVersionName,omitempty"`
}

func (s ListMiniEngineVersionsResponseBodyKernelVersions) String() string {
	return dara.Prettify(s)
}

func (s ListMiniEngineVersionsResponseBodyKernelVersions) GoString() string {
	return s.String()
}

func (s *ListMiniEngineVersionsResponseBodyKernelVersions) GetKernelReleaseType() *string {
	return s.KernelReleaseType
}

func (s *ListMiniEngineVersionsResponseBodyKernelVersions) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *ListMiniEngineVersionsResponseBodyKernelVersions) GetKernelVersionName() *string {
	return s.KernelVersionName
}

func (s *ListMiniEngineVersionsResponseBodyKernelVersions) SetKernelReleaseType(v string) *ListMiniEngineVersionsResponseBodyKernelVersions {
	s.KernelReleaseType = &v
	return s
}

func (s *ListMiniEngineVersionsResponseBodyKernelVersions) SetKernelVersion(v string) *ListMiniEngineVersionsResponseBodyKernelVersions {
	s.KernelVersion = &v
	return s
}

func (s *ListMiniEngineVersionsResponseBodyKernelVersions) SetKernelVersionName(v string) *ListMiniEngineVersionsResponseBodyKernelVersions {
	s.KernelVersionName = &v
	return s
}

func (s *ListMiniEngineVersionsResponseBodyKernelVersions) Validate() error {
	return dara.Validate(s)
}
