// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeHybridProxyPolicyResponseBody
	GetCount() *int32
	SetPolicyList(v []*DescribeHybridProxyPolicyResponseBodyPolicyList) *DescribeHybridProxyPolicyResponseBody
	GetPolicyList() []*DescribeHybridProxyPolicyResponseBodyPolicyList
	SetRequestId(v string) *DescribeHybridProxyPolicyResponseBody
	GetRequestId() *string
}

type DescribeHybridProxyPolicyResponseBody struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The data collection configurations of the proxy cluster.
	PolicyList []*DescribeHybridProxyPolicyResponseBodyPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F7A1B40A-7EED-55A0-BCBC-2F83A486F0AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridProxyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyPolicyResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeHybridProxyPolicyResponseBody) GetPolicyList() []*DescribeHybridProxyPolicyResponseBodyPolicyList {
	return s.PolicyList
}

func (s *DescribeHybridProxyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridProxyPolicyResponseBody) SetCount(v int32) *DescribeHybridProxyPolicyResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBody) SetPolicyList(v []*DescribeHybridProxyPolicyResponseBodyPolicyList) *DescribeHybridProxyPolicyResponseBody {
	s.PolicyList = v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBody) SetRequestId(v string) *DescribeHybridProxyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBody) Validate() error {
	if s.PolicyList != nil {
		for _, item := range s.PolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridProxyPolicyResponseBodyPolicyList struct {
	// The information about the policy.
	Info []*DescribeHybridProxyPolicyResponseBodyPolicyListInfo `json:"Info,omitempty" xml:"Info,omitempty" type:"Repeated"`
	// The type of the policy. Valid values:
	//
	// 	- **limitFrequency**
	//
	// 	- **limitBandWidth**
	//
	// example:
	//
	// limitBandWidth
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeHybridProxyPolicyResponseBodyPolicyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyPolicyResponseBodyPolicyList) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyList) GetInfo() []*DescribeHybridProxyPolicyResponseBodyPolicyListInfo {
	return s.Info
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyList) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyList) SetInfo(v []*DescribeHybridProxyPolicyResponseBodyPolicyListInfo) *DescribeHybridProxyPolicyResponseBodyPolicyList {
	s.Info = v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyList) SetPolicyType(v string) *DescribeHybridProxyPolicyResponseBodyPolicyList {
	s.PolicyType = &v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyList) Validate() error {
	if s.Info != nil {
		for _, item := range s.Info {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridProxyPolicyResponseBodyPolicyListInfo struct {
	// The value of the policy configurations.
	//
	// example:
	//
	// 10
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the file. After you configure a blocking policy, the blocked data is written to the file.
	//
	// example:
	//
	// test
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the policy that you configured. Valid values:
	//
	// 	- **file**
	//
	// 	- **net**
	//
	// 	- **process**
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHybridProxyPolicyResponseBodyPolicyListInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyPolicyResponseBodyPolicyListInfo) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyListInfo) GetConfig() *string {
	return s.Config
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyListInfo) GetFileName() *string {
	return s.FileName
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyListInfo) GetType() *string {
	return s.Type
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyListInfo) SetConfig(v string) *DescribeHybridProxyPolicyResponseBodyPolicyListInfo {
	s.Config = &v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyListInfo) SetFileName(v string) *DescribeHybridProxyPolicyResponseBodyPolicyListInfo {
	s.FileName = &v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyListInfo) SetType(v string) *DescribeHybridProxyPolicyResponseBodyPolicyListInfo {
	s.Type = &v
	return s
}

func (s *DescribeHybridProxyPolicyResponseBodyPolicyListInfo) Validate() error {
	return dara.Validate(s)
}
