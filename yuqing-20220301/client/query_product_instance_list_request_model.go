// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *QueryProductInstanceListRequest
	GetAppCode() *string
	SetFromTime(v int64) *QueryProductInstanceListRequest
	GetFromTime() *int64
	SetRequestId(v string) *QueryProductInstanceListRequest
	GetRequestId() *string
	SetTenantUid(v string) *QueryProductInstanceListRequest
	GetTenantUid() *string
	SetToTime(v int64) *QueryProductInstanceListRequest
	GetToTime() *int64
}

type QueryProductInstanceListRequest struct {
	AppCode   *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	FromTime  *int64  `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TenantUid *string `json:"tenantUid,omitempty" xml:"tenantUid,omitempty"`
	ToTime    *int64  `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s QueryProductInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryProductInstanceListRequest) GoString() string {
	return s.String()
}

func (s *QueryProductInstanceListRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryProductInstanceListRequest) GetFromTime() *int64 {
	return s.FromTime
}

func (s *QueryProductInstanceListRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProductInstanceListRequest) GetTenantUid() *string {
	return s.TenantUid
}

func (s *QueryProductInstanceListRequest) GetToTime() *int64 {
	return s.ToTime
}

func (s *QueryProductInstanceListRequest) SetAppCode(v string) *QueryProductInstanceListRequest {
	s.AppCode = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetFromTime(v int64) *QueryProductInstanceListRequest {
	s.FromTime = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetRequestId(v string) *QueryProductInstanceListRequest {
	s.RequestId = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetTenantUid(v string) *QueryProductInstanceListRequest {
	s.TenantUid = &v
	return s
}

func (s *QueryProductInstanceListRequest) SetToTime(v int64) *QueryProductInstanceListRequest {
	s.ToTime = &v
	return s
}

func (s *QueryProductInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
