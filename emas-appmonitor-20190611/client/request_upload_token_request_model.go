// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestUploadTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *RequestUploadTokenRequest
	GetAppKey() *int64
	SetOs(v string) *RequestUploadTokenRequest
	GetOs() *string
}

type RequestUploadTokenRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
}

func (s RequestUploadTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RequestUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *RequestUploadTokenRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *RequestUploadTokenRequest) GetOs() *string {
	return s.Os
}

func (s *RequestUploadTokenRequest) SetAppKey(v int64) *RequestUploadTokenRequest {
	s.AppKey = &v
	return s
}

func (s *RequestUploadTokenRequest) SetOs(v string) *RequestUploadTokenRequest {
	s.Os = &v
	return s
}

func (s *RequestUploadTokenRequest) Validate() error {
	return dara.Validate(s)
}
