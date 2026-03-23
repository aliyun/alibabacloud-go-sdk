// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApInfoFromPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *GetApInfoFromPoolRequest
	GetApMac() *string
	SetAppCode(v string) *GetApInfoFromPoolRequest
	GetAppCode() *string
	SetAppName(v string) *GetApInfoFromPoolRequest
	GetAppName() *string
}

type GetApInfoFromPoolRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetApInfoFromPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApInfoFromPoolRequest) GoString() string {
	return s.String()
}

func (s *GetApInfoFromPoolRequest) GetApMac() *string {
	return s.ApMac
}

func (s *GetApInfoFromPoolRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApInfoFromPoolRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApInfoFromPoolRequest) SetApMac(v string) *GetApInfoFromPoolRequest {
	s.ApMac = &v
	return s
}

func (s *GetApInfoFromPoolRequest) SetAppCode(v string) *GetApInfoFromPoolRequest {
	s.AppCode = &v
	return s
}

func (s *GetApInfoFromPoolRequest) SetAppName(v string) *GetApInfoFromPoolRequest {
	s.AppName = &v
	return s
}

func (s *GetApInfoFromPoolRequest) Validate() error {
	return dara.Validate(s)
}
