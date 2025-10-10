// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppSecurityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *QueryAppSecurityInfoRequest
	GetAppKey() *string
}

type QueryAppSecurityInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryAppSecurityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAppSecurityInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *QueryAppSecurityInfoRequest) SetAppKey(v string) *QueryAppSecurityInfoRequest {
	s.AppKey = &v
	return s
}

func (s *QueryAppSecurityInfoRequest) Validate() error {
	return dara.Validate(s)
}
