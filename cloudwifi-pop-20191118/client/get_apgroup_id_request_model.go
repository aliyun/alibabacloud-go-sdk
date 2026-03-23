// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *GetApgroupIdRequest
	GetApMac() *string
	SetAppCode(v string) *GetApgroupIdRequest
	GetAppCode() *string
	SetAppName(v string) *GetApgroupIdRequest
	GetAppName() *string
}

type GetApgroupIdRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetApgroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupIdRequest) GoString() string {
	return s.String()
}

func (s *GetApgroupIdRequest) GetApMac() *string {
	return s.ApMac
}

func (s *GetApgroupIdRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApgroupIdRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApgroupIdRequest) SetApMac(v string) *GetApgroupIdRequest {
	s.ApMac = &v
	return s
}

func (s *GetApgroupIdRequest) SetAppCode(v string) *GetApgroupIdRequest {
	s.AppCode = &v
	return s
}

func (s *GetApgroupIdRequest) SetAppName(v string) *GetApgroupIdRequest {
	s.AppName = &v
	return s
}

func (s *GetApgroupIdRequest) Validate() error {
	return dara.Validate(s)
}
