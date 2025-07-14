// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetSwitchInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDatasetSwitchInfoResponseBody
	GetRequestId() *string
	SetResult(v *QueryDatasetSwitchInfoResponseBodyResult) *QueryDatasetSwitchInfoResponseBody
	GetResult() *QueryDatasetSwitchInfoResponseBodyResult
	SetSuccess(v bool) *QueryDatasetSwitchInfoResponseBody
	GetSuccess() *bool
}

type QueryDatasetSwitchInfoResponseBody struct {
	// example:
	//
	// FAECEFA8-09BB-58AB-BC58-C8ACEFE4D232
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDatasetSwitchInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetSwitchInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetSwitchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDatasetSwitchInfoResponseBody) GetResult() *QueryDatasetSwitchInfoResponseBodyResult {
	return s.Result
}

func (s *QueryDatasetSwitchInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDatasetSwitchInfoResponseBody) SetRequestId(v string) *QueryDatasetSwitchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBody) SetResult(v *QueryDatasetSwitchInfoResponseBodyResult) *QueryDatasetSwitchInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBody) SetSuccess(v bool) *QueryDatasetSwitchInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDatasetSwitchInfoResponseBodyResult struct {
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// example:
	//
	// 1
	IsOpenColumnLevelPermission *int32 `json:"IsOpenColumnLevelPermission,omitempty" xml:"IsOpenColumnLevelPermission,omitempty"`
	// example:
	//
	// 1
	IsOpenRowLevelPermission *int32 `json:"IsOpenRowLevelPermission,omitempty" xml:"IsOpenRowLevelPermission,omitempty"`
}

func (s QueryDatasetSwitchInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetSwitchInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) GetIsOpenColumnLevelPermission() *int32 {
	return s.IsOpenColumnLevelPermission
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) GetIsOpenRowLevelPermission() *int32 {
	return s.IsOpenRowLevelPermission
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) SetCubeId(v string) *QueryDatasetSwitchInfoResponseBodyResult {
	s.CubeId = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) SetIsOpenColumnLevelPermission(v int32) *QueryDatasetSwitchInfoResponseBodyResult {
	s.IsOpenColumnLevelPermission = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) SetIsOpenRowLevelPermission(v int32) *QueryDatasetSwitchInfoResponseBodyResult {
	s.IsOpenRowLevelPermission = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
