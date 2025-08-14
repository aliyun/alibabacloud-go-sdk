// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchQueryVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SwitchQueryVariableRequest
	GetLang() *string
	SetId(v int64) *SwitchQueryVariableRequest
	GetId() *int64
	SetRegId(v string) *SwitchQueryVariableRequest
	GetRegId() *string
	SetStatus(v string) *SwitchQueryVariableRequest
	GetStatus() *string
}

type SwitchQueryVariableRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SwitchQueryVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchQueryVariableRequest) GoString() string {
	return s.String()
}

func (s *SwitchQueryVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *SwitchQueryVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *SwitchQueryVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *SwitchQueryVariableRequest) GetStatus() *string {
	return s.Status
}

func (s *SwitchQueryVariableRequest) SetLang(v string) *SwitchQueryVariableRequest {
	s.Lang = &v
	return s
}

func (s *SwitchQueryVariableRequest) SetId(v int64) *SwitchQueryVariableRequest {
	s.Id = &v
	return s
}

func (s *SwitchQueryVariableRequest) SetRegId(v string) *SwitchQueryVariableRequest {
	s.RegId = &v
	return s
}

func (s *SwitchQueryVariableRequest) SetStatus(v string) *SwitchQueryVariableRequest {
	s.Status = &v
	return s
}

func (s *SwitchQueryVariableRequest) Validate() error {
	return dara.Validate(s)
}
