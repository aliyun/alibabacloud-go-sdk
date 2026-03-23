// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupSsidConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApGroupUUId(v string) *GetApgroupSsidConfigRequest
	GetApGroupUUId() *string
	SetAppCode(v string) *GetApgroupSsidConfigRequest
	GetAppCode() *string
	SetAppName(v string) *GetApgroupSsidConfigRequest
	GetAppName() *string
}

type GetApgroupSsidConfigRequest struct {
	// This parameter is required.
	ApGroupUUId *string `json:"ApGroupUUId,omitempty" xml:"ApGroupUUId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetApgroupSsidConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupSsidConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApgroupSsidConfigRequest) GetApGroupUUId() *string {
	return s.ApGroupUUId
}

func (s *GetApgroupSsidConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApgroupSsidConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApgroupSsidConfigRequest) SetApGroupUUId(v string) *GetApgroupSsidConfigRequest {
	s.ApGroupUUId = &v
	return s
}

func (s *GetApgroupSsidConfigRequest) SetAppCode(v string) *GetApgroupSsidConfigRequest {
	s.AppCode = &v
	return s
}

func (s *GetApgroupSsidConfigRequest) SetAppName(v string) *GetApgroupSsidConfigRequest {
	s.AppName = &v
	return s
}

func (s *GetApgroupSsidConfigRequest) Validate() error {
	return dara.Validate(s)
}
