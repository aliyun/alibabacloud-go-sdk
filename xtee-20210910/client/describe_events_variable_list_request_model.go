// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsVariableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventsVariableListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeEventsVariableListRequest
	GetCreateType() *string
	SetEventCodes(v string) *DescribeEventsVariableListRequest
	GetEventCodes() *string
	SetFilterDTO(v string) *DescribeEventsVariableListRequest
	GetFilterDTO() *string
	SetRegId(v string) *DescribeEventsVariableListRequest
	GetRegId() *string
	SetScene(v string) *DescribeEventsVariableListRequest
	GetScene() *string
}

type DescribeEventsVariableListRequest struct {
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
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code.
	//
	// This parameter is required.
	//
	// example:
	//
	// de_afghcf6411
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Filter object
	//
	// example:
	//
	// {\\"type\\":\\"EXPRESSION\\",\\"name\\":\\"ex_NgR6nDVD821c\\"}
	FilterDTO *string `json:"filterDTO,omitempty" xml:"filterDTO,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Applicable scene code
	//
	// This parameter is required.
	//
	// example:
	//
	// VELOCITY
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s DescribeEventsVariableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsVariableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsVariableListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventsVariableListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeEventsVariableListRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeEventsVariableListRequest) GetFilterDTO() *string {
	return s.FilterDTO
}

func (s *DescribeEventsVariableListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventsVariableListRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeEventsVariableListRequest) SetLang(v string) *DescribeEventsVariableListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventsVariableListRequest) SetCreateType(v string) *DescribeEventsVariableListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeEventsVariableListRequest) SetEventCodes(v string) *DescribeEventsVariableListRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeEventsVariableListRequest) SetFilterDTO(v string) *DescribeEventsVariableListRequest {
	s.FilterDTO = &v
	return s
}

func (s *DescribeEventsVariableListRequest) SetRegId(v string) *DescribeEventsVariableListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventsVariableListRequest) SetScene(v string) *DescribeEventsVariableListRequest {
	s.Scene = &v
	return s
}

func (s *DescribeEventsVariableListRequest) Validate() error {
	return dara.Validate(s)
}
