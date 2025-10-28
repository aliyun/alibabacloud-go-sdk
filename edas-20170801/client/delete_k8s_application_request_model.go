// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteK8sApplicationRequest
	GetAppId() *string
	SetForce(v bool) *DeleteK8sApplicationRequest
	GetForce() *bool
}

type DeleteK8sApplicationRequest struct {
	// The ID of the application that you want to delete. You can call the ListApplication operation to query the application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbac7e3c-****-49bc-b6de-ffc550018b45
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to forcibly delete the application and disable application deletion protection.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteK8sApplicationRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteK8sApplicationRequest) SetAppId(v string) *DeleteK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeleteK8sApplicationRequest) SetForce(v bool) *DeleteK8sApplicationRequest {
	s.Force = &v
	return s
}

func (s *DeleteK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
