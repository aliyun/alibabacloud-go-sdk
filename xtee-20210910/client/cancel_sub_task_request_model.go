// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSubTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CancelSubTaskRequest
	GetLang() *string
	SetRegId(v string) *CancelSubTaskRequest
	GetRegId() *string
	SetSubTaskId(v int32) *CancelSubTaskRequest
	GetSubTaskId() *int32
	SetTab(v string) *CancelSubTaskRequest
	GetTab() *string
}

type CancelSubTaskRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// example:
	//
	// 2
	SubTaskId *int32 `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s CancelSubTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelSubTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelSubTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CancelSubTaskRequest) GetRegId() *string {
	return s.RegId
}

func (s *CancelSubTaskRequest) GetSubTaskId() *int32 {
	return s.SubTaskId
}

func (s *CancelSubTaskRequest) GetTab() *string {
	return s.Tab
}

func (s *CancelSubTaskRequest) SetLang(v string) *CancelSubTaskRequest {
	s.Lang = &v
	return s
}

func (s *CancelSubTaskRequest) SetRegId(v string) *CancelSubTaskRequest {
	s.RegId = &v
	return s
}

func (s *CancelSubTaskRequest) SetSubTaskId(v int32) *CancelSubTaskRequest {
	s.SubTaskId = &v
	return s
}

func (s *CancelSubTaskRequest) SetTab(v string) *CancelSubTaskRequest {
	s.Tab = &v
	return s
}

func (s *CancelSubTaskRequest) Validate() error {
	return dara.Validate(s)
}
