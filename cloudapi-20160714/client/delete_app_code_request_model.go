// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *DeleteAppCodeRequest
	GetAppCode() *string
	SetAppId(v string) *DeleteAppCodeRequest
	GetAppId() *string
}

type DeleteAppCodeRequest struct {
	// The application AppCode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d13f021c5cd4997831a9717e75b0663
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111265074
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAppCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppCodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppCodeRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DeleteAppCodeRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppCodeRequest) SetAppCode(v string) *DeleteAppCodeRequest {
	s.AppCode = &v
	return s
}

func (s *DeleteAppCodeRequest) SetAppId(v string) *DeleteAppCodeRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppCodeRequest) Validate() error {
	return dara.Validate(s)
}
