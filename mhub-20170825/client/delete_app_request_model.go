// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *DeleteAppRequest
	GetAppKey() *string
}

type DeleteAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *DeleteAppRequest) SetAppKey(v string) *DeleteAppRequest {
	s.AppKey = &v
	return s
}

func (s *DeleteAppRequest) Validate() error {
	return dara.Validate(s)
}
