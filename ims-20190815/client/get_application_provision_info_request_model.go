// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetApplicationProvisionInfoRequest
	GetAppId() *string
}

type GetApplicationProvisionInfoRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationProvisionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationProvisionInfoRequest) SetAppId(v string) *GetApplicationProvisionInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationProvisionInfoRequest) Validate() error {
	return dara.Validate(s)
}
