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
	// b2a8a925-477a-4ed7-b825-d5e22500****
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
