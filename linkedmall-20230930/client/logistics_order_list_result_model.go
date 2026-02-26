// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogisticsOrderListResult interface {
	dara.Model
	String() string
	GoString() string
	SetLogisticsOrderList(v []*LogisticsOrderResult) *LogisticsOrderListResult
	GetLogisticsOrderList() []*LogisticsOrderResult
	SetRequestId(v string) *LogisticsOrderListResult
	GetRequestId() *string
}

type LogisticsOrderListResult struct {
	LogisticsOrderList []*LogisticsOrderResult `json:"logisticsOrderList,omitempty" xml:"logisticsOrderList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LogisticsOrderListResult) String() string {
	return dara.Prettify(s)
}

func (s LogisticsOrderListResult) GoString() string {
	return s.String()
}

func (s *LogisticsOrderListResult) GetLogisticsOrderList() []*LogisticsOrderResult {
	return s.LogisticsOrderList
}

func (s *LogisticsOrderListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *LogisticsOrderListResult) SetLogisticsOrderList(v []*LogisticsOrderResult) *LogisticsOrderListResult {
	s.LogisticsOrderList = v
	return s
}

func (s *LogisticsOrderListResult) SetRequestId(v string) *LogisticsOrderListResult {
	s.RequestId = &v
	return s
}

func (s *LogisticsOrderListResult) Validate() error {
	if s.LogisticsOrderList != nil {
		for _, item := range s.LogisticsOrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
