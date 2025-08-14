// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSimulationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *StartSimulationTaskRequest
	GetLang() *string
	SetId(v string) *StartSimulationTaskRequest
	GetId() *string
	SetRegId(v string) *StartSimulationTaskRequest
	GetRegId() *string
}

type StartSimulationTaskRequest struct {
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
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s StartSimulationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSimulationTaskRequest) GoString() string {
	return s.String()
}

func (s *StartSimulationTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *StartSimulationTaskRequest) GetId() *string {
	return s.Id
}

func (s *StartSimulationTaskRequest) GetRegId() *string {
	return s.RegId
}

func (s *StartSimulationTaskRequest) SetLang(v string) *StartSimulationTaskRequest {
	s.Lang = &v
	return s
}

func (s *StartSimulationTaskRequest) SetId(v string) *StartSimulationTaskRequest {
	s.Id = &v
	return s
}

func (s *StartSimulationTaskRequest) SetRegId(v string) *StartSimulationTaskRequest {
	s.RegId = &v
	return s
}

func (s *StartSimulationTaskRequest) Validate() error {
	return dara.Validate(s)
}
