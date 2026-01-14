// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTempStorageLeaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *ApplyTempStorageLeaseRequest
	GetFileName() *string
	SetSizeInBytes(v int64) *ApplyTempStorageLeaseRequest
	GetSizeInBytes() *int64
}

type ApplyTempStorageLeaseRequest struct {
	// The file name, including the file extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the file, in bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
}

func (s ApplyTempStorageLeaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyTempStorageLeaseRequest) GoString() string {
	return s.String()
}

func (s *ApplyTempStorageLeaseRequest) GetFileName() *string {
	return s.FileName
}

func (s *ApplyTempStorageLeaseRequest) GetSizeInBytes() *int64 {
	return s.SizeInBytes
}

func (s *ApplyTempStorageLeaseRequest) SetFileName(v string) *ApplyTempStorageLeaseRequest {
	s.FileName = &v
	return s
}

func (s *ApplyTempStorageLeaseRequest) SetSizeInBytes(v int64) *ApplyTempStorageLeaseRequest {
	s.SizeInBytes = &v
	return s
}

func (s *ApplyTempStorageLeaseRequest) Validate() error {
	return dara.Validate(s)
}
