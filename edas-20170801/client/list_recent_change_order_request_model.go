// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentChangeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListRecentChangeOrderRequest
	GetAppId() *string
}

type ListRecentChangeOrderRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListRecentChangeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecentChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListRecentChangeOrderRequest) SetAppId(v string) *ListRecentChangeOrderRequest {
	s.AppId = &v
	return s
}

func (s *ListRecentChangeOrderRequest) Validate() error {
	return dara.Validate(s)
}
