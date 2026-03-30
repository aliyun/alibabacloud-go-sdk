// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppSecretIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppSecretIdsRequest
	GetAppId() *string
}

type ListAppSecretIdsRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListAppSecretIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppSecretIdsRequest) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppSecretIdsRequest) SetAppId(v string) *ListAppSecretIdsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppSecretIdsRequest) Validate() error {
	return dara.Validate(s)
}
