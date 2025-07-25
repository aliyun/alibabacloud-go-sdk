// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIspFlushCacheInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateIspFlushCacheInstanceConfigRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateIspFlushCacheInstanceConfigRequest
	GetInstanceName() *string
	SetLang(v string) *UpdateIspFlushCacheInstanceConfigRequest
	GetLang() *string
}

type UpdateIspFlushCacheInstanceConfigRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s UpdateIspFlushCacheInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIspFlushCacheInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateIspFlushCacheInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateIspFlushCacheInstanceConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateIspFlushCacheInstanceConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateIspFlushCacheInstanceConfigRequest) SetInstanceId(v string) *UpdateIspFlushCacheInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateIspFlushCacheInstanceConfigRequest) SetInstanceName(v string) *UpdateIspFlushCacheInstanceConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateIspFlushCacheInstanceConfigRequest) SetLang(v string) *UpdateIspFlushCacheInstanceConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateIspFlushCacheInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
