// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDataRangeResponseBody
	GetRequestId() *string
	SetResult(v *QueryDataRangeResponseBodyResult) *QueryDataRangeResponseBody
	GetResult() *QueryDataRangeResponseBodyResult
	SetSuccess(v bool) *QueryDataRangeResponseBody
	GetSuccess() *bool
}

type QueryDataRangeResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-****-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Data range object.
	Result *QueryDataRangeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRangeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDataRangeResponseBody) GetResult() *QueryDataRangeResponseBodyResult {
	return s.Result
}

func (s *QueryDataRangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDataRangeResponseBody) SetRequestId(v string) *QueryDataRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataRangeResponseBody) SetResult(v *QueryDataRangeResponseBodyResult) *QueryDataRangeResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataRangeResponseBody) SetSuccess(v bool) *QueryDataRangeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDataRangeResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDataRangeResponseBodyResult struct {
	// Array of LlmCube resources.
	ApiCopilotLlmCubeModels []*QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels `json:"ApiCopilotLlmCubeModels,omitempty" xml:"ApiCopilotLlmCubeModels,omitempty" type:"Repeated"`
	// Array of analysis themes.
	ApiCopilotThemeModels []*QueryDataRangeResponseBodyResultApiCopilotThemeModels `json:"ApiCopilotThemeModels,omitempty" xml:"ApiCopilotThemeModels,omitempty" type:"Repeated"`
}

func (s QueryDataRangeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRangeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResult) GetApiCopilotLlmCubeModels() []*QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels {
	return s.ApiCopilotLlmCubeModels
}

func (s *QueryDataRangeResponseBodyResult) GetApiCopilotThemeModels() []*QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	return s.ApiCopilotThemeModels
}

func (s *QueryDataRangeResponseBodyResult) SetApiCopilotLlmCubeModels(v []*QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) *QueryDataRangeResponseBodyResult {
	s.ApiCopilotLlmCubeModels = v
	return s
}

func (s *QueryDataRangeResponseBodyResult) SetApiCopilotThemeModels(v []*QueryDataRangeResponseBodyResultApiCopilotThemeModels) *QueryDataRangeResponseBodyResult {
	s.ApiCopilotThemeModels = v
	return s
}

func (s *QueryDataRangeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels struct {
	// Alias of the LlmCube resource.
	//
	// example:
	//
	// test
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// Nickname of the creator.
	//
	// example:
	//
	// zhuge
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// LlmCube resource ID.
	//
	// example:
	//
	// sdasdafas23342342342
	LlmCubeId *string `json:"LlmCubeId,omitempty" xml:"LlmCubeId,omitempty"`
}

func (s QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) GetAlias() *string {
	return s.Alias
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) GetCreateUser() *string {
	return s.CreateUser
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) GetLlmCubeId() *string {
	return s.LlmCubeId
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) SetAlias(v string) *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels {
	s.Alias = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) SetCreateUser(v string) *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels {
	s.CreateUser = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) SetLlmCubeId(v string) *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels {
	s.LlmCubeId = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) Validate() error {
	return dara.Validate(s)
}

type QueryDataRangeResponseBodyResultApiCopilotThemeModels struct {
	// Array of LlmCube resources.
	ApiCopilotLlmCubeModels []*QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels `json:"ApiCopilotLlmCubeModels,omitempty" xml:"ApiCopilotLlmCubeModels,omitempty" type:"Repeated"`
	// Nickname of the creator.
	//
	// example:
	//
	// zhuge
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Analysis theme ID.
	//
	// example:
	//
	// 36631232342312312
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// Nickname of the analysis theme.
	//
	// example:
	//
	// test theme
	ThemeName *string `json:"ThemeName,omitempty" xml:"ThemeName,omitempty"`
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModels) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModels) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) GetApiCopilotLlmCubeModels() []*QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels {
	return s.ApiCopilotLlmCubeModels
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) GetCreateUser() *string {
	return s.CreateUser
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) GetThemeId() *string {
	return s.ThemeId
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) GetThemeName() *string {
	return s.ThemeName
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetApiCopilotLlmCubeModels(v []*QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.ApiCopilotLlmCubeModels = v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetCreateUser(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.CreateUser = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetThemeId(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.ThemeId = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetThemeName(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.ThemeName = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) Validate() error {
	return dara.Validate(s)
}

type QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels struct {
	// Alias of the LLM cube resource.
	//
	// example:
	//
	// test
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// Nickname of the creator.
	//
	// example:
	//
	// zhuge
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// LlmCube resource ID.
	//
	// example:
	//
	// 1231242231asdasda
	LlmCubeId *string `json:"LlmCubeId,omitempty" xml:"LlmCubeId,omitempty"`
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) GetAlias() *string {
	return s.Alias
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) GetCreateUser() *string {
	return s.CreateUser
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) GetLlmCubeId() *string {
	return s.LlmCubeId
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) SetAlias(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels {
	s.Alias = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) SetCreateUser(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels {
	s.CreateUser = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) SetLlmCubeId(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels {
	s.LlmCubeId = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) Validate() error {
	return dara.Validate(s)
}
