// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendValidateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *SendValidateFileResponseBody
	GetFileId() *string
	SetRequestId(v string) *SendValidateFileResponseBody
	GetRequestId() *string
}

type SendValidateFileResponseBody struct {
	// example:
	//
	// yyyy-yyyy-yyyy-yyyy
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendValidateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendValidateFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendValidateFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *SendValidateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendValidateFileResponseBody) SetFileId(v string) *SendValidateFileResponseBody {
	s.FileId = &v
	return s
}

func (s *SendValidateFileResponseBody) SetRequestId(v string) *SendValidateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendValidateFileResponseBody) Validate() error {
	return dara.Validate(s)
}
