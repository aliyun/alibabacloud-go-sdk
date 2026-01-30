// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterTypeInfos(v []*ListClusterTypeResponseBodyClusterTypeInfos) *ListClusterTypeResponseBody
	GetClusterTypeInfos() []*ListClusterTypeResponseBodyClusterTypeInfos
	SetClusterTypes(v []*string) *ListClusterTypeResponseBody
	GetClusterTypes() []*string
	SetRequestId(v string) *ListClusterTypeResponseBody
	GetRequestId() *string
}

type ListClusterTypeResponseBody struct {
	ClusterTypeInfos []*ListClusterTypeResponseBodyClusterTypeInfos `json:"ClusterTypeInfos,omitempty" xml:"ClusterTypeInfos,omitempty" type:"Repeated"`
	ClusterTypes     []*string                                      `json:"ClusterTypes,omitempty" xml:"ClusterTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 88BA193C-4918-08C4-9870-E1FE6BBECF34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTypeResponseBody) GetClusterTypeInfos() []*ListClusterTypeResponseBodyClusterTypeInfos {
	return s.ClusterTypeInfos
}

func (s *ListClusterTypeResponseBody) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *ListClusterTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterTypeResponseBody) SetClusterTypeInfos(v []*ListClusterTypeResponseBodyClusterTypeInfos) *ListClusterTypeResponseBody {
	s.ClusterTypeInfos = v
	return s
}

func (s *ListClusterTypeResponseBody) SetClusterTypes(v []*string) *ListClusterTypeResponseBody {
	s.ClusterTypes = v
	return s
}

func (s *ListClusterTypeResponseBody) SetRequestId(v string) *ListClusterTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTypeResponseBody) Validate() error {
	if s.ClusterTypeInfos != nil {
		for _, item := range s.ClusterTypeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterTypeResponseBodyClusterTypeInfos struct {
	// example:
	//
	// SSD
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	IsMultiAZ   *bool   `json:"IsMultiAZ,omitempty" xml:"IsMultiAZ,omitempty"`
}

func (s ListClusterTypeResponseBodyClusterTypeInfos) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypeResponseBodyClusterTypeInfos) GoString() string {
	return s.String()
}

func (s *ListClusterTypeResponseBodyClusterTypeInfos) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClusterTypeResponseBodyClusterTypeInfos) GetIsMultiAZ() *bool {
	return s.IsMultiAZ
}

func (s *ListClusterTypeResponseBodyClusterTypeInfos) SetClusterType(v string) *ListClusterTypeResponseBodyClusterTypeInfos {
	s.ClusterType = &v
	return s
}

func (s *ListClusterTypeResponseBodyClusterTypeInfos) SetIsMultiAZ(v bool) *ListClusterTypeResponseBodyClusterTypeInfos {
	s.IsMultiAZ = &v
	return s
}

func (s *ListClusterTypeResponseBodyClusterTypeInfos) Validate() error {
	return dara.Validate(s)
}
