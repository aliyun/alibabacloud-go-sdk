// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMessageAppRequest
	GetAppId() *string
}

type DeleteMessageAppRequest struct {
	// The ID of the interactive messaging application that you want to delete. You can specify only one ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMessageAppRequest) SetAppId(v string) *DeleteMessageAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
