// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSoarStrategyTaskDetailRequest
	GetLang() *string
	SetStrategyTaskId(v int64) *DescribeSoarStrategyTaskDetailRequest
	GetStrategyTaskId() *int64
}

type DescribeSoarStrategyTaskDetailRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the policy task.
	//
	// >  You can call the [DescribeSoarStrategyTasks](~~DescribeSoarStrategyTasks~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10585
	StrategyTaskId *int64 `json:"StrategyTaskId,omitempty" xml:"StrategyTaskId,omitempty"`
}

func (s DescribeSoarStrategyTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSoarStrategyTaskDetailRequest) GetStrategyTaskId() *int64 {
	return s.StrategyTaskId
}

func (s *DescribeSoarStrategyTaskDetailRequest) SetLang(v string) *DescribeSoarStrategyTaskDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailRequest) SetStrategyTaskId(v int64) *DescribeSoarStrategyTaskDetailRequest {
	s.StrategyTaskId = &v
	return s
}

func (s *DescribeSoarStrategyTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
