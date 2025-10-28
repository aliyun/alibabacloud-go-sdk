// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetApplicationRequest
	GetAppId() *string
}

type GetApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 29f0******************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationRequest) SetAppId(v string) *GetApplicationRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
