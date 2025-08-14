// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectName(v string) *ModelFileUploadRequest
	GetObjectName() *string
	SetRegId(v string) *ModelFileUploadRequest
	GetRegId() *string
}

type ModelFileUploadRequest struct {
	// File name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-07-10/商品信息批量导出-20230710132028
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s ModelFileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelFileUploadRequest) GoString() string {
	return s.String()
}

func (s *ModelFileUploadRequest) GetObjectName() *string {
	return s.ObjectName
}

func (s *ModelFileUploadRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModelFileUploadRequest) SetObjectName(v string) *ModelFileUploadRequest {
	s.ObjectName = &v
	return s
}

func (s *ModelFileUploadRequest) SetRegId(v string) *ModelFileUploadRequest {
	s.RegId = &v
	return s
}

func (s *ModelFileUploadRequest) Validate() error {
	return dara.Validate(s)
}
