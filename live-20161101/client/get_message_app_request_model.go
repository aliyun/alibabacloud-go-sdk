// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMessageAppRequest
	GetAppId() *string
}

type GetMessageAppRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageAppRequest) GoString() string {
	return s.String()
}

func (s *GetMessageAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageAppRequest) SetAppId(v string) *GetMessageAppRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
