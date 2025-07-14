// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumedServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListConsumedServicesRequest
	GetAppId() *string
}

type ListConsumedServicesRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListConsumedServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesRequest) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListConsumedServicesRequest) SetAppId(v string) *ListConsumedServicesRequest {
	s.AppId = &v
	return s
}

func (s *ListConsumedServicesRequest) Validate() error {
	return dara.Validate(s)
}
