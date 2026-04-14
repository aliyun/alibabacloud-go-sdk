// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileUrl(v string) *GetValidateFileResponseBody
	GetFileUrl() *string
	SetRequestId(v string) *GetValidateFileResponseBody
	GetRequestId() *string
}

type GetValidateFileResponseBody struct {
	// example:
	//
	// https://xxxxxx/yyy
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// yyyy-yyyy-yyyy-yyyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetValidateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetValidateFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetValidateFileResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetValidateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetValidateFileResponseBody) SetFileUrl(v string) *GetValidateFileResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetValidateFileResponseBody) SetRequestId(v string) *GetValidateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetValidateFileResponseBody) Validate() error {
	return dara.Validate(s)
}
