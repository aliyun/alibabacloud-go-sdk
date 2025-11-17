// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLlmCubeWithThemeListByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryLlmCubeWithThemeListByUserIdResponseBody
	GetRequestId() *string
	SetResult(v *QueryLlmCubeWithThemeListByUserIdResponseBodyResult) *QueryLlmCubeWithThemeListByUserIdResponseBody
	GetResult() *QueryLlmCubeWithThemeListByUserIdResponseBodyResult
	SetSuccess(v bool) *QueryLlmCubeWithThemeListByUserIdResponseBody
	GetSuccess() *bool
}

type QueryLlmCubeWithThemeListByUserIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 2EE822B***************F-F5B42DDADC12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of datasets and analysis themes.
	Result *QueryLlmCubeWithThemeListByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLlmCubeWithThemeListByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLlmCubeWithThemeListByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBody) GetResult() *QueryLlmCubeWithThemeListByUserIdResponseBodyResult {
	return s.Result
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBody) SetRequestId(v string) *QueryLlmCubeWithThemeListByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBody) SetResult(v *QueryLlmCubeWithThemeListByUserIdResponseBodyResult) *QueryLlmCubeWithThemeListByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBody) SetSuccess(v bool) *QueryLlmCubeWithThemeListByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryLlmCubeWithThemeListByUserIdResponseBodyResult struct {
	// Dataset map.
	//
	// - key - Dataset ID
	//
	// - value - Dataset name
	CubeIds map[string]*string `json:"CubeIds,omitempty" xml:"CubeIds,omitempty"`
	// Analysis theme map.
	//
	// - key - Analysis theme ID
	//
	// - value - Analysis theme name
	ThemeIds map[string]*string `json:"ThemeIds,omitempty" xml:"ThemeIds,omitempty"`
}

func (s QueryLlmCubeWithThemeListByUserIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryLlmCubeWithThemeListByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBodyResult) GetCubeIds() map[string]*string {
	return s.CubeIds
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBodyResult) GetThemeIds() map[string]*string {
	return s.ThemeIds
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBodyResult) SetCubeIds(v map[string]*string) *QueryLlmCubeWithThemeListByUserIdResponseBodyResult {
	s.CubeIds = v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBodyResult) SetThemeIds(v map[string]*string) *QueryLlmCubeWithThemeListByUserIdResponseBodyResult {
	s.ThemeIds = v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
