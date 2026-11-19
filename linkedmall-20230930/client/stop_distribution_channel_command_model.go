// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDistributionChannelCommand interface {
	dara.Model
	String() string
	GoString() string
	SetChannelCodes(v []*string) *StopDistributionChannelCommand
	GetChannelCodes() []*string
	SetLmShopId(v string) *StopDistributionChannelCommand
	GetLmShopId() *string
	SetProductId(v string) *StopDistributionChannelCommand
	GetProductId() *string
	SetRequestId(v string) *StopDistributionChannelCommand
	GetRequestId() *string
	SetRequestTime(v string) *StopDistributionChannelCommand
	GetRequestTime() *string
	SetRequestUser(v string) *StopDistributionChannelCommand
	GetRequestUser() *string
}

type StopDistributionChannelCommand struct {
	// The set of channel IDs.
	ChannelCodes []*string `json:"channelCodes,omitempty" xml:"channelCodes,omitempty" type:"Repeated"`
	// The store ID.
	//
	// example:
	//
	// 1100001
	LmShopId *string `json:"lmShopId,omitempty" xml:"lmShopId,omitempty"`
	// The commodity ID.
	//
	// example:
	//
	// 660******080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// The requesting user.
	//
	// example:
	//
	// 12000000001
	RequestUser *string `json:"requestUser,omitempty" xml:"requestUser,omitempty"`
}

func (s StopDistributionChannelCommand) String() string {
	return dara.Prettify(s)
}

func (s StopDistributionChannelCommand) GoString() string {
	return s.String()
}

func (s *StopDistributionChannelCommand) GetChannelCodes() []*string {
	return s.ChannelCodes
}

func (s *StopDistributionChannelCommand) GetLmShopId() *string {
	return s.LmShopId
}

func (s *StopDistributionChannelCommand) GetProductId() *string {
	return s.ProductId
}

func (s *StopDistributionChannelCommand) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDistributionChannelCommand) GetRequestTime() *string {
	return s.RequestTime
}

func (s *StopDistributionChannelCommand) GetRequestUser() *string {
	return s.RequestUser
}

func (s *StopDistributionChannelCommand) SetChannelCodes(v []*string) *StopDistributionChannelCommand {
	s.ChannelCodes = v
	return s
}

func (s *StopDistributionChannelCommand) SetLmShopId(v string) *StopDistributionChannelCommand {
	s.LmShopId = &v
	return s
}

func (s *StopDistributionChannelCommand) SetProductId(v string) *StopDistributionChannelCommand {
	s.ProductId = &v
	return s
}

func (s *StopDistributionChannelCommand) SetRequestId(v string) *StopDistributionChannelCommand {
	s.RequestId = &v
	return s
}

func (s *StopDistributionChannelCommand) SetRequestTime(v string) *StopDistributionChannelCommand {
	s.RequestTime = &v
	return s
}

func (s *StopDistributionChannelCommand) SetRequestUser(v string) *StopDistributionChannelCommand {
	s.RequestUser = &v
	return s
}

func (s *StopDistributionChannelCommand) Validate() error {
	return dara.Validate(s)
}
