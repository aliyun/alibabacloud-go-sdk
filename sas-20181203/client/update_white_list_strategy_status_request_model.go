// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteListStrategyStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateWhiteListStrategyStatusRequest
	GetLang() *string
	SetSourceIp(v string) *UpdateWhiteListStrategyStatusRequest
	GetSourceIp() *string
	SetStatus(v int32) *UpdateWhiteListStrategyStatusRequest
	GetStatus() *int32
	SetStrategyIds(v string) *UpdateWhiteListStrategyStatusRequest
	GetStrategyIds() *string
}

type UpdateWhiteListStrategyStatusRequest struct {
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
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 219.143.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **0**: deleted
	//
	// 	- **1**: learning
	//
	// 	- **2**: paused
	//
	// 	- **3**: learning completed
	//
	// 	- **4**: enabled
	//
	// >
	//
	// 	- You can change the status to **paused*	- only if the policy status is **learning**.
	//
	// 	- You can change the status to **learning*	- only if the policy status is **paused**.
	//
	// 	- You can change the status to **enabled*	- only if the policy status is **learning completed**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the policy.
	//
	// >  You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8516
	StrategyIds *string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty"`
}

func (s UpdateWhiteListStrategyStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteListStrategyStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhiteListStrategyStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateWhiteListStrategyStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *UpdateWhiteListStrategyStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateWhiteListStrategyStatusRequest) GetStrategyIds() *string {
	return s.StrategyIds
}

func (s *UpdateWhiteListStrategyStatusRequest) SetLang(v string) *UpdateWhiteListStrategyStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateWhiteListStrategyStatusRequest) SetSourceIp(v string) *UpdateWhiteListStrategyStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateWhiteListStrategyStatusRequest) SetStatus(v int32) *UpdateWhiteListStrategyStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateWhiteListStrategyStatusRequest) SetStrategyIds(v string) *UpdateWhiteListStrategyStatusRequest {
	s.StrategyIds = &v
	return s
}

func (s *UpdateWhiteListStrategyStatusRequest) Validate() error {
	return dara.Validate(s)
}
