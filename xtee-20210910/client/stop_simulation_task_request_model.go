// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSimulationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *StopSimulationTaskRequest
	GetLang() *string
	SetId(v string) *StopSimulationTaskRequest
	GetId() *string
	SetRegId(v string) *StopSimulationTaskRequest
	GetRegId() *string
}

type StopSimulationTaskRequest struct {
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
	// Task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 376773
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s StopSimulationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopSimulationTaskRequest) GoString() string {
	return s.String()
}

func (s *StopSimulationTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *StopSimulationTaskRequest) GetId() *string {
	return s.Id
}

func (s *StopSimulationTaskRequest) GetRegId() *string {
	return s.RegId
}

func (s *StopSimulationTaskRequest) SetLang(v string) *StopSimulationTaskRequest {
	s.Lang = &v
	return s
}

func (s *StopSimulationTaskRequest) SetId(v string) *StopSimulationTaskRequest {
	s.Id = &v
	return s
}

func (s *StopSimulationTaskRequest) SetRegId(v string) *StopSimulationTaskRequest {
	s.RegId = &v
	return s
}

func (s *StopSimulationTaskRequest) Validate() error {
	return dara.Validate(s)
}
