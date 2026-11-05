// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDetectUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHashKeyContextList(v []*CreateFileDetectUploadUrlRequestHashKeyContextList) *CreateFileDetectUploadUrlRequest
	GetHashKeyContextList() []*CreateFileDetectUploadUrlRequestHashKeyContextList
	SetHashKeyList(v []*string) *CreateFileDetectUploadUrlRequest
	GetHashKeyList() []*string
	SetType(v int32) *CreateFileDetectUploadUrlRequest
	GetType() *int32
}

type CreateFileDetectUploadUrlRequest struct {
	// The file identifier. Only MD5 is supported. You can specify up to 100 elements in the list.
	//
	// > **HashKeyList*	- and **HashKeyContextList*	- cannot be empty at the same time.
	HashKeyContextList []*CreateFileDetectUploadUrlRequestHashKeyContextList `json:"HashKeyContextList,omitempty" xml:"HashKeyContextList,omitempty" type:"Repeated"`
	// The unique identifier of the file. You can specify up to 100 elements in the list.
	//
	// > **HashKeyList*	- and **HashKeyContextList*	- cannot be empty at the same time.
	//
	// example:
	//
	// CreateFileDetectUploadUrl
	HashKeyList []*string `json:"HashKeyList,omitempty" xml:"HashKeyList,omitempty" type:"Repeated"`
	// The file type. Valid values:
	//
	// - **0**: Uncertain file type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFileDetectUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateFileDetectUploadUrlRequest) GetHashKeyContextList() []*CreateFileDetectUploadUrlRequestHashKeyContextList {
	return s.HashKeyContextList
}

func (s *CreateFileDetectUploadUrlRequest) GetHashKeyList() []*string {
	return s.HashKeyList
}

func (s *CreateFileDetectUploadUrlRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateFileDetectUploadUrlRequest) SetHashKeyContextList(v []*CreateFileDetectUploadUrlRequestHashKeyContextList) *CreateFileDetectUploadUrlRequest {
	s.HashKeyContextList = v
	return s
}

func (s *CreateFileDetectUploadUrlRequest) SetHashKeyList(v []*string) *CreateFileDetectUploadUrlRequest {
	s.HashKeyList = v
	return s
}

func (s *CreateFileDetectUploadUrlRequest) SetType(v int32) *CreateFileDetectUploadUrlRequest {
	s.Type = &v
	return s
}

func (s *CreateFileDetectUploadUrlRequest) Validate() error {
	if s.HashKeyContextList != nil {
		for _, item := range s.HashKeyContextList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFileDetectUploadUrlRequestHashKeyContextList struct {
	// The file size. If the file size exceeds the limit, the upload parameters are not returned. Unit: bytes.
	//
	// example:
	//
	// 2698557
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The file hash value.
	//
	// example:
	//
	// 30319dd5cee8f894766e479cac170da0
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
}

func (s CreateFileDetectUploadUrlRequestHashKeyContextList) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectUploadUrlRequestHashKeyContextList) GoString() string {
	return s.String()
}

func (s *CreateFileDetectUploadUrlRequestHashKeyContextList) GetFileSize() *int32 {
	return s.FileSize
}

func (s *CreateFileDetectUploadUrlRequestHashKeyContextList) GetHashKey() *string {
	return s.HashKey
}

func (s *CreateFileDetectUploadUrlRequestHashKeyContextList) SetFileSize(v int32) *CreateFileDetectUploadUrlRequestHashKeyContextList {
	s.FileSize = &v
	return s
}

func (s *CreateFileDetectUploadUrlRequestHashKeyContextList) SetHashKey(v string) *CreateFileDetectUploadUrlRequestHashKeyContextList {
	s.HashKey = &v
	return s
}

func (s *CreateFileDetectUploadUrlRequestHashKeyContextList) Validate() error {
	return dara.Validate(s)
}
