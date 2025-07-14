// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmartqPermissionByCubeIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *QuerySmartqPermissionByCubeIdRequest
	GetCubeId() *string
	SetUserId(v string) *QuerySmartqPermissionByCubeIdRequest
	GetUserId() *string
}

type QuerySmartqPermissionByCubeIdRequest struct {
	// Dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// User ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95c4d**************3852e202
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySmartqPermissionByCubeIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmartqPermissionByCubeIdRequest) GoString() string {
	return s.String()
}

func (s *QuerySmartqPermissionByCubeIdRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *QuerySmartqPermissionByCubeIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *QuerySmartqPermissionByCubeIdRequest) SetCubeId(v string) *QuerySmartqPermissionByCubeIdRequest {
	s.CubeId = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdRequest) SetUserId(v string) *QuerySmartqPermissionByCubeIdRequest {
	s.UserId = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdRequest) Validate() error {
	return dara.Validate(s)
}
