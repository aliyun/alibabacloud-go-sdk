// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMd5List(v []*string) *CreateBatchUploadUrlRequest
	GetMd5List() []*string
	SetType(v int32) *CreateBatchUploadUrlRequest
	GetType() *int32
}

type CreateBatchUploadUrlRequest struct {
	// The identifiers of files. Only MD5 hash values are supported.
	//
	// This parameter is required.
	Md5List []*string `json:"Md5List,omitempty" xml:"Md5List,omitempty" type:"Repeated"`
	// The type of the file. Valid values:
	//
	// 	- **0**: unknown file
	//
	// 	- **1**: binary file
	//
	// 	- **2**: webshell file
	//
	// 	- **4**: script file
	//
	// > If you do not know the type of the file, set this parameter to **0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBatchUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchUploadUrlRequest) GetMd5List() []*string {
	return s.Md5List
}

func (s *CreateBatchUploadUrlRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateBatchUploadUrlRequest) SetMd5List(v []*string) *CreateBatchUploadUrlRequest {
	s.Md5List = v
	return s
}

func (s *CreateBatchUploadUrlRequest) SetType(v int32) *CreateBatchUploadUrlRequest {
	s.Type = &v
	return s
}

func (s *CreateBatchUploadUrlRequest) Validate() error {
	return dara.Validate(s)
}
