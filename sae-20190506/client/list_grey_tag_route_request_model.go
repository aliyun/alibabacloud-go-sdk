// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGreyTagRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListGreyTagRouteRequest
	GetAppId() *string
}

type ListGreyTagRouteRequest struct {
	// 7171a6ca-d1cd-4928-8642-7d5cfe69\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListGreyTagRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListGreyTagRouteRequest) SetAppId(v string) *ListGreyTagRouteRequest {
	s.AppId = &v
	return s
}

func (s *ListGreyTagRouteRequest) Validate() error {
	return dara.Validate(s)
}
