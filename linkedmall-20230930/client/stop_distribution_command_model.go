// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDistributionCommand interface {
	dara.Model
	String() string
	GoString() string
	SetLmShopId(v string) *StopDistributionCommand
	GetLmShopId() *string
	SetProductId(v string) *StopDistributionCommand
	GetProductId() *string
	SetRequestId(v string) *StopDistributionCommand
	GetRequestId() *string
	SetRequestTime(v string) *StopDistributionCommand
	GetRequestTime() *string
	SetRequestUser(v string) *StopDistributionCommand
	GetRequestUser() *string
}

type StopDistributionCommand struct {
	LmShopId  *string `json:"lmShopId,omitempty" xml:"lmShopId,omitempty"`
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-12-01 10:01:00
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	RequestUser *string `json:"requestUser,omitempty" xml:"requestUser,omitempty"`
}

func (s StopDistributionCommand) String() string {
	return dara.Prettify(s)
}

func (s StopDistributionCommand) GoString() string {
	return s.String()
}

func (s *StopDistributionCommand) GetLmShopId() *string {
	return s.LmShopId
}

func (s *StopDistributionCommand) GetProductId() *string {
	return s.ProductId
}

func (s *StopDistributionCommand) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDistributionCommand) GetRequestTime() *string {
	return s.RequestTime
}

func (s *StopDistributionCommand) GetRequestUser() *string {
	return s.RequestUser
}

func (s *StopDistributionCommand) SetLmShopId(v string) *StopDistributionCommand {
	s.LmShopId = &v
	return s
}

func (s *StopDistributionCommand) SetProductId(v string) *StopDistributionCommand {
	s.ProductId = &v
	return s
}

func (s *StopDistributionCommand) SetRequestId(v string) *StopDistributionCommand {
	s.RequestId = &v
	return s
}

func (s *StopDistributionCommand) SetRequestTime(v string) *StopDistributionCommand {
	s.RequestTime = &v
	return s
}

func (s *StopDistributionCommand) SetRequestUser(v string) *StopDistributionCommand {
	s.RequestUser = &v
	return s
}

func (s *StopDistributionCommand) Validate() error {
	return dara.Validate(s)
}
