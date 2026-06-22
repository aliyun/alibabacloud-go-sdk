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
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. The system automatically obtains this value.
	//
	// example:
	//
	// 219.143.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The policy status. Valid values:
	//
	// - **0**: Deleted.
	//
	// - **1**: Learning.
	//
	// - **2**: Paused.
	//
	// - **3**: Learning complete.
	//
	// - **4**: Active.
	//
	// > - Only a policy in the **Learning*	- state can be changed to the **Paused*	- state.
	//
	// > - Only a policy in the **Paused*	- state can be changed to the **Learning*	- state.
	//
	// > - Only a policy in the **Learning complete*	- state can be changed to the **Active*	- state.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The policy ID.
	//
	// >Call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to obtain this parameter.
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
