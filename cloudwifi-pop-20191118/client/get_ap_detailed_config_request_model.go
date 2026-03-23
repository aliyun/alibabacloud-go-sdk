// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApDetailedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *GetApDetailedConfigRequest
	GetApMac() *string
	SetAppCode(v string) *GetApDetailedConfigRequest
	GetAppCode() *string
	SetAppName(v string) *GetApDetailedConfigRequest
	GetAppName() *string
}

type GetApDetailedConfigRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetApDetailedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApDetailedConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApDetailedConfigRequest) GetApMac() *string {
	return s.ApMac
}

func (s *GetApDetailedConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApDetailedConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApDetailedConfigRequest) SetApMac(v string) *GetApDetailedConfigRequest {
	s.ApMac = &v
	return s
}

func (s *GetApDetailedConfigRequest) SetAppCode(v string) *GetApDetailedConfigRequest {
	s.AppCode = &v
	return s
}

func (s *GetApDetailedConfigRequest) SetAppName(v string) *GetApDetailedConfigRequest {
	s.AppName = &v
	return s
}

func (s *GetApDetailedConfigRequest) Validate() error {
	return dara.Validate(s)
}
