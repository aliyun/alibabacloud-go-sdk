// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableArmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DisableArmsRequest
	GetAppId() *string
}

type DisableArmsRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69a26c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DisableArmsRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableArmsRequest) GoString() string {
	return s.String()
}

func (s *DisableArmsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DisableArmsRequest) SetAppId(v string) *DisableArmsRequest {
	s.AppId = &v
	return s
}

func (s *DisableArmsRequest) Validate() error {
	return dara.Validate(s)
}
