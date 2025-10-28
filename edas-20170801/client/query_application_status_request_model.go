// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApplicationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryApplicationStatusRequest
	GetAppId() *string
}

type QueryApplicationStatusRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-441**************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s QueryApplicationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryApplicationStatusRequest) SetAppId(v string) *QueryApplicationStatusRequest {
	s.AppId = &v
	return s
}

func (s *QueryApplicationStatusRequest) Validate() error {
	return dara.Validate(s)
}
