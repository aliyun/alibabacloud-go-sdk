// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApGroupUUId(v string) *DeleteApgroupConfigRequest
	GetApGroupUUId() *string
	SetAppCode(v string) *DeleteApgroupConfigRequest
	GetAppCode() *string
	SetAppName(v string) *DeleteApgroupConfigRequest
	GetAppName() *string
}

type DeleteApgroupConfigRequest struct {
	// This parameter is required.
	ApGroupUUId *string `json:"ApGroupUUId,omitempty" xml:"ApGroupUUId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s DeleteApgroupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteApgroupConfigRequest) GetApGroupUUId() *string {
	return s.ApGroupUUId
}

func (s *DeleteApgroupConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DeleteApgroupConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteApgroupConfigRequest) SetApGroupUUId(v string) *DeleteApgroupConfigRequest {
	s.ApGroupUUId = &v
	return s
}

func (s *DeleteApgroupConfigRequest) SetAppCode(v string) *DeleteApgroupConfigRequest {
	s.AppCode = &v
	return s
}

func (s *DeleteApgroupConfigRequest) SetAppName(v string) *DeleteApgroupConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteApgroupConfigRequest) Validate() error {
	return dara.Validate(s)
}
