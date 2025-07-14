// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppDescription(v string) *UpdateApplicationDescriptionRequest
	GetAppDescription() *string
	SetAppId(v string) *UpdateApplicationDescriptionRequest
	GetAppId() *string
}

type UpdateApplicationDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// newdesc
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s UpdateApplicationDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionRequest) GetAppDescription() *string {
	return s.AppDescription
}

func (s *UpdateApplicationDescriptionRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationDescriptionRequest) SetAppDescription(v string) *UpdateApplicationDescriptionRequest {
	s.AppDescription = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) SetAppId(v string) *UpdateApplicationDescriptionRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
