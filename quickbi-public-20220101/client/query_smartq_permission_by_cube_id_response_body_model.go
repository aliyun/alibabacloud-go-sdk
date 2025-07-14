// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmartqPermissionByCubeIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QuerySmartqPermissionByCubeIdResponseBody
	GetRequestId() *string
	SetResult(v *QuerySmartqPermissionByCubeIdResponseBodyResult) *QuerySmartqPermissionByCubeIdResponseBody
	GetResult() *QuerySmartqPermissionByCubeIdResponseBodyResult
	SetSuccess(v bool) *QuerySmartqPermissionByCubeIdResponseBody
	GetSuccess() *bool
}

type QuerySmartqPermissionByCubeIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 617277******************ABA47E31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Basic information of the dataset.
	Result *QuerySmartqPermissionByCubeIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QuerySmartqPermissionByCubeIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmartqPermissionByCubeIdResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmartqPermissionByCubeIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmartqPermissionByCubeIdResponseBody) GetResult() *QuerySmartqPermissionByCubeIdResponseBodyResult {
	return s.Result
}

func (s *QuerySmartqPermissionByCubeIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySmartqPermissionByCubeIdResponseBody) SetRequestId(v string) *QuerySmartqPermissionByCubeIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponseBody) SetResult(v *QuerySmartqPermissionByCubeIdResponseBodyResult) *QuerySmartqPermissionByCubeIdResponseBody {
	s.Result = v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponseBody) SetSuccess(v bool) *QuerySmartqPermissionByCubeIdResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySmartqPermissionByCubeIdResponseBodyResult struct {
	// Dataset ID.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// Dataset name.
	//
	// example:
	//
	// test
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// Whether the current user has permission for the smart question. Note: \\"HasPerssion\\" seems to be a typo, it should probably be \\"HasPermission\\".
	HasPerssion *bool `json:"HasPerssion,omitempty" xml:"HasPerssion,omitempty"`
}

func (s QuerySmartqPermissionByCubeIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QuerySmartqPermissionByCubeIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySmartqPermissionByCubeIdResponseBodyResult) GetCubeId() *string {
	return s.CubeId
}

func (s *QuerySmartqPermissionByCubeIdResponseBodyResult) GetCubeName() *string {
	return s.CubeName
}

func (s *QuerySmartqPermissionByCubeIdResponseBodyResult) GetHasPerssion() *bool {
	return s.HasPerssion
}

func (s *QuerySmartqPermissionByCubeIdResponseBodyResult) SetCubeId(v string) *QuerySmartqPermissionByCubeIdResponseBodyResult {
	s.CubeId = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponseBodyResult) SetCubeName(v string) *QuerySmartqPermissionByCubeIdResponseBodyResult {
	s.CubeName = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponseBodyResult) SetHasPerssion(v bool) *QuerySmartqPermissionByCubeIdResponseBodyResult {
	s.HasPerssion = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
