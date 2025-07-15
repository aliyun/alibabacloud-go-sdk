// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *CreateAppCodeRequest
	GetAppCode() *string
	SetAppId(v string) *CreateAppCodeRequest
	GetAppId() *string
}

type CreateAppCodeRequest struct {
	// The application AppCode.
	//
	// example:
	//
	// 3aaf905a0a1f4f0eabc6d891dfa08afc
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111203109
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateAppCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppCodeRequest) GoString() string {
	return s.String()
}

func (s *CreateAppCodeRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *CreateAppCodeRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppCodeRequest) SetAppCode(v string) *CreateAppCodeRequest {
	s.AppCode = &v
	return s
}

func (s *CreateAppCodeRequest) SetAppId(v string) *CreateAppCodeRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppCodeRequest) Validate() error {
	return dara.Validate(s)
}
