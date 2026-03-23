// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupSsidConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupId(v int64) *DeleteApgroupSsidConfigRequest
	GetApgroupId() *int64
	SetAppCode(v string) *DeleteApgroupSsidConfigRequest
	GetAppCode() *string
	SetAppName(v string) *DeleteApgroupSsidConfigRequest
	GetAppName() *string
	SetId(v int64) *DeleteApgroupSsidConfigRequest
	GetId() *int64
}

type DeleteApgroupSsidConfigRequest struct {
	// This parameter is required.
	ApgroupId *int64 `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteApgroupSsidConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupSsidConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteApgroupSsidConfigRequest) GetApgroupId() *int64 {
	return s.ApgroupId
}

func (s *DeleteApgroupSsidConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DeleteApgroupSsidConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteApgroupSsidConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteApgroupSsidConfigRequest) SetApgroupId(v int64) *DeleteApgroupSsidConfigRequest {
	s.ApgroupId = &v
	return s
}

func (s *DeleteApgroupSsidConfigRequest) SetAppCode(v string) *DeleteApgroupSsidConfigRequest {
	s.AppCode = &v
	return s
}

func (s *DeleteApgroupSsidConfigRequest) SetAppName(v string) *DeleteApgroupSsidConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteApgroupSsidConfigRequest) SetId(v int64) *DeleteApgroupSsidConfigRequest {
	s.Id = &v
	return s
}

func (s *DeleteApgroupSsidConfigRequest) Validate() error {
	return dara.Validate(s)
}
