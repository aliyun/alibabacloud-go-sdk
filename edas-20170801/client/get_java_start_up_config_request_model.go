// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJavaStartUpConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetJavaStartUpConfigRequest
	GetAppId() *string
}

type GetJavaStartUpConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5fdf50e8-*
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetJavaStartUpConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJavaStartUpConfigRequest) GoString() string {
	return s.String()
}

func (s *GetJavaStartUpConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetJavaStartUpConfigRequest) SetAppId(v string) *GetJavaStartUpConfigRequest {
	s.AppId = &v
	return s
}

func (s *GetJavaStartUpConfigRequest) Validate() error {
	return dara.Validate(s)
}
