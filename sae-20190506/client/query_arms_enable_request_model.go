// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryArmsEnableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryArmsEnableRequest
	GetAppId() *string
}

type QueryArmsEnableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s QueryArmsEnableRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryArmsEnableRequest) GoString() string {
	return s.String()
}

func (s *QueryArmsEnableRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryArmsEnableRequest) SetAppId(v string) *QueryArmsEnableRequest {
	s.AppId = &v
	return s
}

func (s *QueryArmsEnableRequest) Validate() error {
	return dara.Validate(s)
}
