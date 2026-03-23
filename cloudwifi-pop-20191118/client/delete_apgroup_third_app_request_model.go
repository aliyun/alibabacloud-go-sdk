// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupThirdAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *DeleteApgroupThirdAppRequest
	GetAppCode() *string
	SetAppName(v string) *DeleteApgroupThirdAppRequest
	GetAppName() *string
	SetId(v int64) *DeleteApgroupThirdAppRequest
	GetId() *int64
}

type DeleteApgroupThirdAppRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteApgroupThirdAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupThirdAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteApgroupThirdAppRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DeleteApgroupThirdAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteApgroupThirdAppRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteApgroupThirdAppRequest) SetAppCode(v string) *DeleteApgroupThirdAppRequest {
	s.AppCode = &v
	return s
}

func (s *DeleteApgroupThirdAppRequest) SetAppName(v string) *DeleteApgroupThirdAppRequest {
	s.AppName = &v
	return s
}

func (s *DeleteApgroupThirdAppRequest) SetId(v int64) *DeleteApgroupThirdAppRequest {
	s.Id = &v
	return s
}

func (s *DeleteApgroupThirdAppRequest) Validate() error {
	return dara.Validate(s)
}
