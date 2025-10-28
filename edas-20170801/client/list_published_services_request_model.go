// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListPublishedServicesRequest
	GetAppId() *string
}

type ListPublishedServicesRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1fbd8b72-****-****-bdfe-478dbc914121
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListPublishedServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedServicesRequest) SetAppId(v string) *ListPublishedServicesRequest {
	s.AppId = &v
	return s
}

func (s *ListPublishedServicesRequest) Validate() error {
	return dara.Validate(s)
}
