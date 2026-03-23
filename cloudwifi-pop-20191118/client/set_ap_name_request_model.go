// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *SetApNameRequest
	GetApMac() *string
	SetAppCode(v string) *SetApNameRequest
	GetAppCode() *string
	SetAppName(v string) *SetApNameRequest
	GetAppName() *string
	SetName(v string) *SetApNameRequest
	GetName() *string
}

type SetApNameRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetApNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApNameRequest) GoString() string {
	return s.String()
}

func (s *SetApNameRequest) GetApMac() *string {
	return s.ApMac
}

func (s *SetApNameRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SetApNameRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetApNameRequest) GetName() *string {
	return s.Name
}

func (s *SetApNameRequest) SetApMac(v string) *SetApNameRequest {
	s.ApMac = &v
	return s
}

func (s *SetApNameRequest) SetAppCode(v string) *SetApNameRequest {
	s.AppCode = &v
	return s
}

func (s *SetApNameRequest) SetAppName(v string) *SetApNameRequest {
	s.AppName = &v
	return s
}

func (s *SetApNameRequest) SetName(v string) *SetApNameRequest {
	s.Name = &v
	return s
}

func (s *SetApNameRequest) Validate() error {
	return dara.Validate(s)
}
