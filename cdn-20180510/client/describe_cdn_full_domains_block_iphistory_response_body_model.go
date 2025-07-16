// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnFullDomainsBlockIPHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCdnFullDomainsBlockIPHistoryResponseBody
	GetCode() *int32
	SetDescription(v string) *DescribeCdnFullDomainsBlockIPHistoryResponseBody
	GetDescription() *string
	SetIPBlockInfo(v []*DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) *DescribeCdnFullDomainsBlockIPHistoryResponseBody
	GetIPBlockInfo() []*DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo
	SetRequestId(v string) *DescribeCdnFullDomainsBlockIPHistoryResponseBody
	GetRequestId() *string
}

type DescribeCdnFullDomainsBlockIPHistoryResponseBody struct {
	// The response code.
	//
	// The value of Code is not 0 in the following scenarios:
	//
	// 	- The format of the IP address is invalid.
	//
	// 	- The format of the time is invalid.
	//
	// 	- Other abnormal scenarios.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the status returned.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The result of the operation.
	IPBlockInfo []*DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo `json:"IPBlockInfo,omitempty" xml:"IPBlockInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BCD7D917-76F1-442F-BB75-C810DE34C761
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnFullDomainsBlockIPHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnFullDomainsBlockIPHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) GetIPBlockInfo() []*DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	return s.IPBlockInfo
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) SetCode(v int32) *DescribeCdnFullDomainsBlockIPHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) SetDescription(v string) *DescribeCdnFullDomainsBlockIPHistoryResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) SetIPBlockInfo(v []*DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) *DescribeCdnFullDomainsBlockIPHistoryResponseBody {
	s.IPBlockInfo = v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) SetRequestId(v string) *DescribeCdnFullDomainsBlockIPHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo struct {
	// The blocked IP address or CIDR block.
	//
	// example:
	//
	// 1.XXX.XXX.0~1.XXX.XXX.255
	BlockIP *string `json:"BlockIP,omitempty" xml:"BlockIP,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 2023-04-24 18:49:37
	DeliverTime *string `json:"DeliverTime,omitempty" xml:"DeliverTime,omitempty"`
	// The delivery status.
	//
	// 	- Success
	//
	// 	- Failed
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetBlockIP() *string {
	return s.BlockIP
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetDeliverTime() *string {
	return s.DeliverTime
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetBlockIP(v string) *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.BlockIP = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetDeliverTime(v string) *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.DeliverTime = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetStatus(v string) *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.Status = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) Validate() error {
	return dara.Validate(s)
}
