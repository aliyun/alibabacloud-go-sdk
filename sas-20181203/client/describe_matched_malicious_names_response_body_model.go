// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMatchedMaliciousNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeMatchedMaliciousNamesResponseBody
	GetCount() *int32
	SetData(v []*DescribeMatchedMaliciousNamesResponseBodyData) *DescribeMatchedMaliciousNamesResponseBody
	GetData() []*DescribeMatchedMaliciousNamesResponseBodyData
	SetRequestId(v string) *DescribeMatchedMaliciousNamesResponseBody
	GetRequestId() *string
}

type DescribeMatchedMaliciousNamesResponseBody struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The response parameters.
	Data []*DescribeMatchedMaliciousNamesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C7CD1BE6-97A2-5524-A529-B55C63E55D59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMatchedMaliciousNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMatchedMaliciousNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMatchedMaliciousNamesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeMatchedMaliciousNamesResponseBody) GetData() []*DescribeMatchedMaliciousNamesResponseBodyData {
	return s.Data
}

func (s *DescribeMatchedMaliciousNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMatchedMaliciousNamesResponseBody) SetCount(v int32) *DescribeMatchedMaliciousNamesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponseBody) SetData(v []*DescribeMatchedMaliciousNamesResponseBodyData) *DescribeMatchedMaliciousNamesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponseBody) SetRequestId(v string) *DescribeMatchedMaliciousNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMatchedMaliciousNamesResponseBodyData struct {
	// The display name of the malicious image sample type.
	//
	// example:
	//
	// displayname
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The key of the malicious image sample type.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeMatchedMaliciousNamesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMatchedMaliciousNamesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMatchedMaliciousNamesResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeMatchedMaliciousNamesResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *DescribeMatchedMaliciousNamesResponseBodyData) SetDisplayName(v string) *DescribeMatchedMaliciousNamesResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponseBodyData) SetKey(v string) *DescribeMatchedMaliciousNamesResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeMatchedMaliciousNamesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
