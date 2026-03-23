// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStaDetailedStatusByMacRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetStaDetailedStatusByMacRequest
	GetAppCode() *string
	SetAppName(v string) *GetStaDetailedStatusByMacRequest
	GetAppName() *string
	SetStaMac(v string) *GetStaDetailedStatusByMacRequest
	GetStaMac() *string
}

type GetStaDetailedStatusByMacRequest struct {
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StaMac  *string `json:"StaMac,omitempty" xml:"StaMac,omitempty"`
}

func (s GetStaDetailedStatusByMacRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStaDetailedStatusByMacRequest) GoString() string {
	return s.String()
}

func (s *GetStaDetailedStatusByMacRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetStaDetailedStatusByMacRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetStaDetailedStatusByMacRequest) GetStaMac() *string {
	return s.StaMac
}

func (s *GetStaDetailedStatusByMacRequest) SetAppCode(v string) *GetStaDetailedStatusByMacRequest {
	s.AppCode = &v
	return s
}

func (s *GetStaDetailedStatusByMacRequest) SetAppName(v string) *GetStaDetailedStatusByMacRequest {
	s.AppName = &v
	return s
}

func (s *GetStaDetailedStatusByMacRequest) SetStaMac(v string) *GetStaDetailedStatusByMacRequest {
	s.StaMac = &v
	return s
}

func (s *GetStaDetailedStatusByMacRequest) Validate() error {
	return dara.Validate(s)
}
