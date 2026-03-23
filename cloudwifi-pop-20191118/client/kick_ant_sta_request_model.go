// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickAntStaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *KickAntStaRequest
	GetAppCode() *string
	SetAppName(v string) *KickAntStaRequest
	GetAppName() *string
	SetStaMac(v string) *KickAntStaRequest
	GetStaMac() *string
}

type KickAntStaRequest struct {
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
	// BA:5F:40:45:63:80
	StaMac *string `json:"StaMac,omitempty" xml:"StaMac,omitempty"`
}

func (s KickAntStaRequest) String() string {
	return dara.Prettify(s)
}

func (s KickAntStaRequest) GoString() string {
	return s.String()
}

func (s *KickAntStaRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *KickAntStaRequest) GetAppName() *string {
	return s.AppName
}

func (s *KickAntStaRequest) GetStaMac() *string {
	return s.StaMac
}

func (s *KickAntStaRequest) SetAppCode(v string) *KickAntStaRequest {
	s.AppCode = &v
	return s
}

func (s *KickAntStaRequest) SetAppName(v string) *KickAntStaRequest {
	s.AppName = &v
	return s
}

func (s *KickAntStaRequest) SetStaMac(v string) *KickAntStaRequest {
	s.StaMac = &v
	return s
}

func (s *KickAntStaRequest) Validate() error {
	return dara.Validate(s)
}
