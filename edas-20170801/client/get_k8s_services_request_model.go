// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetK8sServicesRequest
	GetAppId() *string
}

type GetK8sServicesRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****-a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetK8sServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetK8sServicesRequest) GoString() string {
	return s.String()
}

func (s *GetK8sServicesRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetK8sServicesRequest) SetAppId(v string) *GetK8sServicesRequest {
	s.AppId = &v
	return s
}

func (s *GetK8sServicesRequest) Validate() error {
	return dara.Validate(s)
}
