// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteCodeBundleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetByteSize(v int64) *CompleteCodeBundleRequest
	GetByteSize() *int64
	SetContentType(v string) *CompleteCodeBundleRequest
	GetContentType() *string
}

type CompleteCodeBundleRequest struct {
	// The declared size of the uploaded object. This value must match the OSS Content-Length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ByteSize *int64 `json:"byteSize,omitempty" xml:"byteSize,omitempty"`
	// The MIME type of the stored code bundle. This is typically application/octet-stream for pre-signed PUT operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// application/octet-stream
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
}

func (s CompleteCodeBundleRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteCodeBundleRequest) GoString() string {
	return s.String()
}

func (s *CompleteCodeBundleRequest) GetByteSize() *int64 {
	return s.ByteSize
}

func (s *CompleteCodeBundleRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CompleteCodeBundleRequest) SetByteSize(v int64) *CompleteCodeBundleRequest {
	s.ByteSize = &v
	return s
}

func (s *CompleteCodeBundleRequest) SetContentType(v string) *CompleteCodeBundleRequest {
	s.ContentType = &v
	return s
}

func (s *CompleteCodeBundleRequest) Validate() error {
	return dara.Validate(s)
}
