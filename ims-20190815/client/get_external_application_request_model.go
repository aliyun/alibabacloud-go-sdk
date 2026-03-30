// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetExternalApplicationRequest
	GetAppId() *string
}

type GetExternalApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetExternalApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExternalApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetExternalApplicationRequest) SetAppId(v string) *GetExternalApplicationRequest {
	s.AppId = &v
	return s
}

func (s *GetExternalApplicationRequest) Validate() error {
	return dara.Validate(s)
}
