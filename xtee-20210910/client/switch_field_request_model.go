// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SwitchFieldRequest
	GetLang() *string
	SetId(v int64) *SwitchFieldRequest
	GetId() *int64
	SetName(v string) *SwitchFieldRequest
	GetName() *string
	SetRegId(v string) *SwitchFieldRequest
	GetRegId() *string
	SetSource(v string) *SwitchFieldRequest
	GetSource() *string
	SetStatus(v string) *SwitchFieldRequest
	GetStatus() *string
}

type SwitchFieldRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary Key ID
	//
	// example:
	//
	// 250002
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Parameter Name.
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region Code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Field Source
	//
	// example:
	//
	// DEFINE
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SwitchFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchFieldRequest) GoString() string {
	return s.String()
}

func (s *SwitchFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *SwitchFieldRequest) GetId() *int64 {
	return s.Id
}

func (s *SwitchFieldRequest) GetName() *string {
	return s.Name
}

func (s *SwitchFieldRequest) GetRegId() *string {
	return s.RegId
}

func (s *SwitchFieldRequest) GetSource() *string {
	return s.Source
}

func (s *SwitchFieldRequest) GetStatus() *string {
	return s.Status
}

func (s *SwitchFieldRequest) SetLang(v string) *SwitchFieldRequest {
	s.Lang = &v
	return s
}

func (s *SwitchFieldRequest) SetId(v int64) *SwitchFieldRequest {
	s.Id = &v
	return s
}

func (s *SwitchFieldRequest) SetName(v string) *SwitchFieldRequest {
	s.Name = &v
	return s
}

func (s *SwitchFieldRequest) SetRegId(v string) *SwitchFieldRequest {
	s.RegId = &v
	return s
}

func (s *SwitchFieldRequest) SetSource(v string) *SwitchFieldRequest {
	s.Source = &v
	return s
}

func (s *SwitchFieldRequest) SetStatus(v string) *SwitchFieldRequest {
	s.Status = &v
	return s
}

func (s *SwitchFieldRequest) Validate() error {
	return dara.Validate(s)
}
