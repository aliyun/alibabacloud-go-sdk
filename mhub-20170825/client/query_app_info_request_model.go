// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *QueryAppInfoRequest
	GetAppKey() *string
}

type QueryAppInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAppInfoRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *QueryAppInfoRequest) SetAppKey(v string) *QueryAppInfoRequest {
	s.AppKey = &v
	return s
}

func (s *QueryAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
