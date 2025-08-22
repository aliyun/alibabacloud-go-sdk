// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnFullDomainsBlockIPHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody
	GetCode() *int32
	SetDescription(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody
	GetDescription() *string
	SetIPBlockInfo(v []*DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody
	GetIPBlockInfo() []*DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo
	SetRequestId(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody
	GetRequestId() *string
}

type DescribeDcdnFullDomainsBlockIPHistoryResponseBody struct {
	// The response code.
	//
	// The value of Code is not 0 in the following scenarios:
	//
	// 	- The format of the IP address is invalid.
	//
	// 	- The format of the time is invalid.
	//
	// 	- Other abnormal scenarios
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
	IPBlockInfo []*DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo `json:"IPBlockInfo,omitempty" xml:"IPBlockInfo,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 13A2B792-9212-1CC9-8525-59EBEF3FFE01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnFullDomainsBlockIPHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnFullDomainsBlockIPHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) GetIPBlockInfo() []*DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	return s.IPBlockInfo
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) SetCode(v int32) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) SetDescription(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) SetIPBlockInfo(v []*DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody {
	s.IPBlockInfo = v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) SetRequestId(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo struct {
	// The blocked IP address or CIDR block.
	//
	// example:
	//
	// 1.XXX.XXX.0~1.XXX.XXX.255
	BlockIP       *string `json:"BlockIP,omitempty" xml:"BlockIP,omitempty"`
	BlockInterval *string `json:"BlockInterval,omitempty" xml:"BlockInterval,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 2023-04-24 18:49:37
	DeliverTime   *string `json:"DeliverTime,omitempty" xml:"DeliverTime,omitempty"`
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The delivery status.
	//
	// 	- Success
	//
	// 	- Failed
	//
	// example:
	//
	// Success
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
}

func (s DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetBlockIP() *string {
	return s.BlockIP
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetBlockInterval() *string {
	return s.BlockInterval
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetDeliverTime() *string {
	return s.DeliverTime
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) GetUpdateType() *string {
	return s.UpdateType
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetBlockIP(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.BlockIP = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetBlockInterval(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.BlockInterval = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetDeliverTime(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.DeliverTime = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetOperationType(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.OperationType = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetStatus(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.Status = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) SetUpdateType(v string) *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo {
	s.UpdateType = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponseBodyIPBlockInfo) Validate() error {
	return dara.Validate(s)
}
