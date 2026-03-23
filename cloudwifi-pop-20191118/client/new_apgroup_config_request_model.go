// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewApgroupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *NewApgroupConfigRequest
	GetAppCode() *string
	SetAppName(v string) *NewApgroupConfigRequest
	GetAppName() *string
	SetName(v string) *NewApgroupConfigRequest
	GetName() *string
	SetParentApgroupId(v string) *NewApgroupConfigRequest
	GetParentApgroupId() *string
}

type NewApgroupConfigRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ParentApgroupId *string `json:"ParentApgroupId,omitempty" xml:"ParentApgroupId,omitempty"`
}

func (s NewApgroupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s NewApgroupConfigRequest) GoString() string {
	return s.String()
}

func (s *NewApgroupConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *NewApgroupConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *NewApgroupConfigRequest) GetName() *string {
	return s.Name
}

func (s *NewApgroupConfigRequest) GetParentApgroupId() *string {
	return s.ParentApgroupId
}

func (s *NewApgroupConfigRequest) SetAppCode(v string) *NewApgroupConfigRequest {
	s.AppCode = &v
	return s
}

func (s *NewApgroupConfigRequest) SetAppName(v string) *NewApgroupConfigRequest {
	s.AppName = &v
	return s
}

func (s *NewApgroupConfigRequest) SetName(v string) *NewApgroupConfigRequest {
	s.Name = &v
	return s
}

func (s *NewApgroupConfigRequest) SetParentApgroupId(v string) *NewApgroupConfigRequest {
	s.ParentApgroupId = &v
	return s
}

func (s *NewApgroupConfigRequest) Validate() error {
	return dara.Validate(s)
}
