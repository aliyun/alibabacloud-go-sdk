// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartApplicationRequest
	GetAppId() *string
}

type StartApplicationRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s StartApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartApplicationRequest) GoString() string {
	return s.String()
}

func (s *StartApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartApplicationRequest) SetAppId(v string) *StartApplicationRequest {
	s.AppId = &v
	return s
}

func (s *StartApplicationRequest) Validate() error {
	return dara.Validate(s)
}
