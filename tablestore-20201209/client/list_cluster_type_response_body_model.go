// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterTypeDetailList(v []*ListClusterTypeResponseBodyClusterTypeDetailList) *ListClusterTypeResponseBody
	GetClusterTypeDetailList() []*ListClusterTypeResponseBodyClusterTypeDetailList
	SetClusterTypeList(v []*string) *ListClusterTypeResponseBody
	GetClusterTypeList() []*string
	SetRequestId(v string) *ListClusterTypeResponseBody
	GetRequestId() *string
}

type ListClusterTypeResponseBody struct {
	ClusterTypeDetailList []*ListClusterTypeResponseBodyClusterTypeDetailList `json:"ClusterTypeDetailList,omitempty" xml:"ClusterTypeDetailList,omitempty" type:"Repeated"`
	ClusterTypeList       []*string                                           `json:"ClusterTypeList,omitempty" xml:"ClusterTypeList,omitempty" type:"Repeated"`
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

func (s *ListClusterTypeResponseBody) GetClusterTypeDetailList() []*ListClusterTypeResponseBodyClusterTypeDetailList {
	return s.ClusterTypeDetailList
}

func (s *ListClusterTypeResponseBody) GetClusterTypeList() []*string {
	return s.ClusterTypeList
}

func (s *ListClusterTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterTypeResponseBody) SetClusterTypeDetailList(v []*ListClusterTypeResponseBodyClusterTypeDetailList) *ListClusterTypeResponseBody {
	s.ClusterTypeDetailList = v
	return s
}

func (s *ListClusterTypeResponseBody) SetClusterTypeList(v []*string) *ListClusterTypeResponseBody {
	s.ClusterTypeList = v
	return s
}

func (s *ListClusterTypeResponseBody) SetRequestId(v string) *ListClusterTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTypeResponseBody) Validate() error {
	if s.ClusterTypeDetailList != nil {
		for _, item := range s.ClusterTypeDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterTypeResponseBodyClusterTypeDetailList struct {
	// example:
	//
	// SSD
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// true
	IsMultiAZ *bool `json:"IsMultiAZ,omitempty" xml:"IsMultiAZ,omitempty"`
}

func (s ListClusterTypeResponseBodyClusterTypeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypeResponseBodyClusterTypeDetailList) GoString() string {
	return s.String()
}

func (s *ListClusterTypeResponseBodyClusterTypeDetailList) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClusterTypeResponseBodyClusterTypeDetailList) GetIsMultiAZ() *bool {
	return s.IsMultiAZ
}

func (s *ListClusterTypeResponseBodyClusterTypeDetailList) SetClusterType(v string) *ListClusterTypeResponseBodyClusterTypeDetailList {
	s.ClusterType = &v
	return s
}

func (s *ListClusterTypeResponseBodyClusterTypeDetailList) SetIsMultiAZ(v bool) *ListClusterTypeResponseBodyClusterTypeDetailList {
	s.IsMultiAZ = &v
	return s
}

func (s *ListClusterTypeResponseBodyClusterTypeDetailList) Validate() error {
	return dara.Validate(s)
}
