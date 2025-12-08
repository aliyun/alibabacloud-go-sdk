// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DeleteFaceRequest
	GetDbName() *string
	SetFaceId(v string) *DeleteFaceRequest
	GetFaceId() *string
}

type DeleteFaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 001
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s DeleteFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteFaceRequest) GetFaceId() *string {
	return s.FaceId
}

func (s *DeleteFaceRequest) SetDbName(v string) *DeleteFaceRequest {
	s.DbName = &v
	return s
}

func (s *DeleteFaceRequest) SetFaceId(v string) *DeleteFaceRequest {
	s.FaceId = &v
	return s
}

func (s *DeleteFaceRequest) Validate() error {
	return dara.Validate(s)
}
