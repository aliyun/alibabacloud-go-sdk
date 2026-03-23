// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAntStaStatusByMacRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetAntStaStatusByMacRequest
	GetAppCode() *string
	SetAppName(v string) *GetAntStaStatusByMacRequest
	GetAppName() *string
	SetStaMac(v string) *GetAntStaStatusByMacRequest
	GetStaMac() *string
}

type GetAntStaStatusByMacRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4fcb82c0-ee83-11ea-88b6-02420b0f63f4
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ISV_ALIYUN_IOT
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BA:5F:40:45:63:89
	StaMac *string `json:"StaMac,omitempty" xml:"StaMac,omitempty"`
}

func (s GetAntStaStatusByMacRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAntStaStatusByMacRequest) GoString() string {
	return s.String()
}

func (s *GetAntStaStatusByMacRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetAntStaStatusByMacRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetAntStaStatusByMacRequest) GetStaMac() *string {
	return s.StaMac
}

func (s *GetAntStaStatusByMacRequest) SetAppCode(v string) *GetAntStaStatusByMacRequest {
	s.AppCode = &v
	return s
}

func (s *GetAntStaStatusByMacRequest) SetAppName(v string) *GetAntStaStatusByMacRequest {
	s.AppName = &v
	return s
}

func (s *GetAntStaStatusByMacRequest) SetStaMac(v string) *GetAntStaStatusByMacRequest {
	s.StaMac = &v
	return s
}

func (s *GetAntStaStatusByMacRequest) Validate() error {
	return dara.Validate(s)
}
