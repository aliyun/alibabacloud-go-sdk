// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicIpAddressPoolCidrBlocksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListPublicIpAddressPoolCidrBlocksResponseBody
	GetNextToken() *string
	SetPublicIpPoolCidrBlockList(v []*ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) *ListPublicIpAddressPoolCidrBlocksResponseBody
	GetPublicIpPoolCidrBlockList() []*ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList
	SetRequestId(v string) *ListPublicIpAddressPoolCidrBlocksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPublicIpAddressPoolCidrBlocksResponseBody
	GetTotalCount() *int32
}

type ListPublicIpAddressPoolCidrBlocksResponseBody struct {
	// The token that is used for the next query. Valid values:
	//
	// 	- If **NextToken*	- was not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- is returned, the value is the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	PublicIpPoolCidrBlockList []*ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList `json:"PublicIpPoolCidrBlockList,omitempty" xml:"PublicIpPoolCidrBlockList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum number of entries returned. Valid values: **10*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicIpAddressPoolCidrBlocksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolCidrBlocksResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) GetPublicIpPoolCidrBlockList() []*ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList {
	return s.PublicIpPoolCidrBlockList
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) SetNextToken(v string) *ListPublicIpAddressPoolCidrBlocksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) SetPublicIpPoolCidrBlockList(v []*ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) *ListPublicIpAddressPoolCidrBlocksResponseBody {
	s.PublicIpPoolCidrBlockList = v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) SetRequestId(v string) *ListPublicIpAddressPoolCidrBlocksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) SetTotalCount(v int32) *ListPublicIpAddressPoolCidrBlocksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList struct {
	// The ID of the IP address pool.
	//
	// example:
	//
	// 47.0.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The CIDR blocks.
	//
	// example:
	//
	// 2022-05-10T01:37:38Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The information about the CIDR blocks.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The time when the CIDR block was created. The time is displayed in `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of available IP addresses in the CIDR block.
	//
	// example:
	//
	// 20
	TotalIpNum *int32 `json:"TotalIpNum,omitempty" xml:"TotalIpNum,omitempty"`
	// The status of the CIDR block in the IP address pool. Valid values:
	//
	// 	- **Created**: available
	//
	// 	- **Deleting**: being deleted
	//
	// 	- **Modifying**: being modified
	//
	// example:
	//
	// 20
	UsedIpNum *int32 `json:"UsedIpNum,omitempty" xml:"UsedIpNum,omitempty"`
}

func (s ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) GetStatus() *string {
	return s.Status
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) GetTotalIpNum() *int32 {
	return s.TotalIpNum
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) GetUsedIpNum() *int32 {
	return s.UsedIpNum
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) SetCidrBlock(v string) *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList {
	s.CidrBlock = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) SetCreationTime(v string) *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList {
	s.CreationTime = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) SetPublicIpAddressPoolId(v string) *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) SetStatus(v string) *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList {
	s.Status = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) SetTotalIpNum(v int32) *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList {
	s.TotalIpNum = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) SetUsedIpNum(v int32) *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList {
	s.UsedIpNum = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponseBodyPublicIpPoolCidrBlockList) Validate() error {
	return dara.Validate(s)
}
