// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *UploadFileResponseBody
	GetFileId() *string
	SetRequestId(v string) *UploadFileResponseBody
	GetRequestId() *string
}

type UploadFileResponseBody struct {
	// example:
	//
	// f-1671330gr7934d4771813f7141d28db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *UploadFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadFileResponseBody) SetFileId(v string) *UploadFileResponseBody {
	s.FileId = &v
	return s
}

func (s *UploadFileResponseBody) SetRequestId(v string) *UploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadFileResponseBody) Validate() error {
	return dara.Validate(s)
}
