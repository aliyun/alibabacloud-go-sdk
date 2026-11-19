// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopilotEmbedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCopilotEmbedConfigResponseBody
	GetRequestId() *string
	SetResult(v []*QueryCopilotEmbedConfigResponseBodyResult) *QueryCopilotEmbedConfigResponseBody
	GetResult() []*QueryCopilotEmbedConfigResponseBodyResult
	SetSuccess(v bool) *QueryCopilotEmbedConfigResponseBody
	GetSuccess() *bool
}

type QueryCopilotEmbedConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1FC71085-D5FD-08E0-813A-4D4BD1031BC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of embedding configurations.
	Result []*QueryCopilotEmbedConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCopilotEmbedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCopilotEmbedConfigResponseBody) GetResult() []*QueryCopilotEmbedConfigResponseBodyResult {
	return s.Result
}

func (s *QueryCopilotEmbedConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCopilotEmbedConfigResponseBody) SetRequestId(v string) *QueryCopilotEmbedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBody) SetResult(v []*QueryCopilotEmbedConfigResponseBodyResult) *QueryCopilotEmbedConfigResponseBody {
	s.Result = v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBody) SetSuccess(v bool) *QueryCopilotEmbedConfigResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCopilotEmbedConfigResponseBodyResult struct {
	// The nickname of the chatbot.
	//
	// example:
	//
	// Xiao Q
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The embedding ID.
	//
	// example:
	//
	// 9c079710-ddbe-48b3-b495-7c83c8d57cc4
	CopilotId *string `json:"CopilotId,omitempty" xml:"CopilotId,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// qweqw12312423521
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The nickname of the creator.
	//
	// example:
	//
	// Zhuge
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// The data scope, including analysis topics and data query resources.
	DataRange *QueryCopilotEmbedConfigResponseBodyResultDataRange `json:"DataRange,omitempty" xml:"DataRange,omitempty" type:"Struct"`
	// The user ID of the modifier.
	//
	// example:
	//
	// asda1231231dfs
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The module name.
	//
	// example:
	//
	// Xiao Q
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The name of the embedding module.
	//
	// example:
	//
	// 0327
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s QueryCopilotEmbedConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetCopilotId() *string {
	return s.CopilotId
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetCreateUser() *string {
	return s.CreateUser
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetDataRange() *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	return s.DataRange
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetModuleName() *string {
	return s.ModuleName
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) GetShowName() *string {
	return s.ShowName
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetAgentName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.AgentName = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetCopilotId(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.CopilotId = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetCreateUser(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetCreateUserName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.CreateUserName = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetDataRange(v *QueryCopilotEmbedConfigResponseBodyResultDataRange) *QueryCopilotEmbedConfigResponseBodyResult {
	s.DataRange = v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetModifyUser(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetModuleName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.ModuleName = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetShowName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.ShowName = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) Validate() error {
	if s.DataRange != nil {
		if err := s.DataRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCopilotEmbedConfigResponseBodyResultDataRange struct {
	// Indicates whether all data query resources are selected.
	//
	// example:
	//
	// true
	AllCube *bool `json:"AllCube,omitempty" xml:"AllCube,omitempty"`
	// Indicates whether all analysis topics are selected.
	//
	// example:
	//
	// true
	AllTheme *bool `json:"AllTheme,omitempty" xml:"AllTheme,omitempty"`
	// The set of data query resource IDs.
	LlmCubes []*string `json:"LlmCubes,omitempty" xml:"LlmCubes,omitempty" type:"Repeated"`
	// The set of analysis topic IDs.
	Themes []*string `json:"Themes,omitempty" xml:"Themes,omitempty" type:"Repeated"`
}

func (s QueryCopilotEmbedConfigResponseBodyResultDataRange) String() string {
	return dara.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponseBodyResultDataRange) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) GetAllCube() *bool {
	return s.AllCube
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) GetAllTheme() *bool {
	return s.AllTheme
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) GetLlmCubes() []*string {
	return s.LlmCubes
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) GetThemes() []*string {
	return s.Themes
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetAllCube(v bool) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.AllCube = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetAllTheme(v bool) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.AllTheme = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetLlmCubes(v []*string) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.LlmCubes = v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetThemes(v []*string) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.Themes = v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) Validate() error {
	return dara.Validate(s)
}
