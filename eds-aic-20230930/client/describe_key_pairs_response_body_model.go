// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyPairsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeKeyPairsResponseBodyData) *DescribeKeyPairsResponseBody
	GetData() []*DescribeKeyPairsResponseBodyData
	SetNextToken(v string) *DescribeKeyPairsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeKeyPairsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeKeyPairsResponseBody
	GetTotalCount() *int32
}

type DescribeKeyPairsResponseBody struct {
	// The objects that are returned.
	Data []*DescribeKeyPairsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// FFbc8N4E1iOlcSxC+8boa0HHH2LKWbggYUinyrZWvtS1oTrMYCg1HuMLGuftj0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 565FB06A-AE04-5AD0-8A32-5BA92CA5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeyPairsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBody) GetData() []*DescribeKeyPairsResponseBodyData {
	return s.Data
}

func (s *DescribeKeyPairsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeKeyPairsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKeyPairsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeKeyPairsResponseBody) SetData(v []*DescribeKeyPairsResponseBodyData) *DescribeKeyPairsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetNextToken(v string) *DescribeKeyPairsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetRequestId(v string) *DescribeKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetTotalCount(v int32) *DescribeKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeKeyPairsResponseBodyData struct {
	// The time when the ADB key pair was created.
	//
	// example:
	//
	// 2022-10-11T08:53:32Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The ID of the ADB key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the ADB key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s DescribeKeyPairsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeKeyPairsResponseBodyData) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DescribeKeyPairsResponseBodyData) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeKeyPairsResponseBodyData) SetGmtCreated(v string) *DescribeKeyPairsResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyData) SetKeyPairId(v string) *DescribeKeyPairsResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyData) SetKeyPairName(v string) *DescribeKeyPairsResponseBodyData {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
