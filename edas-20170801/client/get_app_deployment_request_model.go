// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAppDeploymentRequest
	GetAppId() *string
}

type GetAppDeploymentRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93fdd228-*****-ed2ae98de18d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetAppDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppDeploymentRequest) GoString() string {
	return s.String()
}

func (s *GetAppDeploymentRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAppDeploymentRequest) SetAppId(v string) *GetAppDeploymentRequest {
	s.AppId = &v
	return s
}

func (s *GetAppDeploymentRequest) Validate() error {
	return dara.Validate(s)
}
