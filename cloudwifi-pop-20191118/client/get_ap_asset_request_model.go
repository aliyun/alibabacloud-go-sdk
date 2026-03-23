// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *GetApAssetRequest
	GetApMac() *string
	SetAppCode(v string) *GetApAssetRequest
	GetAppCode() *string
	SetAppName(v string) *GetApAssetRequest
	GetAppName() *string
}

type GetApAssetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 14:15:14:15:14:15
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XXXIIII
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ISV
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetApAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApAssetRequest) GoString() string {
	return s.String()
}

func (s *GetApAssetRequest) GetApMac() *string {
	return s.ApMac
}

func (s *GetApAssetRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApAssetRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApAssetRequest) SetApMac(v string) *GetApAssetRequest {
	s.ApMac = &v
	return s
}

func (s *GetApAssetRequest) SetAppCode(v string) *GetApAssetRequest {
	s.AppCode = &v
	return s
}

func (s *GetApAssetRequest) SetAppName(v string) *GetApAssetRequest {
	s.AppName = &v
	return s
}

func (s *GetApAssetRequest) Validate() error {
	return dara.Validate(s)
}
