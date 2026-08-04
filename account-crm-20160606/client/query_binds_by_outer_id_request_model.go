// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBindsByOuterIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *QueryBindsByOuterIdRequest
	GetAppName() *string
	SetMinorOuterId(v string) *QueryBindsByOuterIdRequest
	GetMinorOuterId() *string
	SetOuterId(v string) *QueryBindsByOuterIdRequest
	GetOuterId() *string
	SetTenantId(v string) *QueryBindsByOuterIdRequest
	GetTenantId() *string
}

type QueryBindsByOuterIdRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	MinorOuterId *string `json:"MinorOuterId,omitempty" xml:"MinorOuterId,omitempty"`
	// This parameter is required.
	OuterId *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryBindsByOuterIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByOuterIdRequest) GoString() string {
	return s.String()
}

func (s *QueryBindsByOuterIdRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryBindsByOuterIdRequest) GetMinorOuterId() *string {
	return s.MinorOuterId
}

func (s *QueryBindsByOuterIdRequest) GetOuterId() *string {
	return s.OuterId
}

func (s *QueryBindsByOuterIdRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryBindsByOuterIdRequest) SetAppName(v string) *QueryBindsByOuterIdRequest {
	s.AppName = &v
	return s
}

func (s *QueryBindsByOuterIdRequest) SetMinorOuterId(v string) *QueryBindsByOuterIdRequest {
	s.MinorOuterId = &v
	return s
}

func (s *QueryBindsByOuterIdRequest) SetOuterId(v string) *QueryBindsByOuterIdRequest {
	s.OuterId = &v
	return s
}

func (s *QueryBindsByOuterIdRequest) SetTenantId(v string) *QueryBindsByOuterIdRequest {
	s.TenantId = &v
	return s
}

func (s *QueryBindsByOuterIdRequest) Validate() error {
	return dara.Validate(s)
}
