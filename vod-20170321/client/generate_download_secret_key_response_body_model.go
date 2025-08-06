// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDownloadSecretKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppEncryptKey(v string) *GenerateDownloadSecretKeyResponseBody
	GetAppEncryptKey() *string
	SetRequestId(v string) *GenerateDownloadSecretKeyResponseBody
	GetRequestId() *string
}

type GenerateDownloadSecretKeyResponseBody struct {
	// The key file for secure download.
	//
	// example:
	//
	// ***
	AppEncryptKey *string `json:"AppEncryptKey,omitempty" xml:"AppEncryptKey,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E99B9BAD-7F9D-552B-A689-B72E92EA040E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDownloadSecretKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDownloadSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDownloadSecretKeyResponseBody) GetAppEncryptKey() *string {
	return s.AppEncryptKey
}

func (s *GenerateDownloadSecretKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDownloadSecretKeyResponseBody) SetAppEncryptKey(v string) *GenerateDownloadSecretKeyResponseBody {
	s.AppEncryptKey = &v
	return s
}

func (s *GenerateDownloadSecretKeyResponseBody) SetRequestId(v string) *GenerateDownloadSecretKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDownloadSecretKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
