// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheReserveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCacheReserveInstanceId(v string) *GetCacheReserveResponseBody
	GetCacheReserveInstanceId() *string
	SetEnable(v string) *GetCacheReserveResponseBody
	GetEnable() *string
	SetRequestId(v string) *GetCacheReserveResponseBody
	GetRequestId() *string
}

type GetCacheReserveResponseBody struct {
	// The cache reserve instance ID.
	//
	// example:
	//
	// cr_hk_123456789
	CacheReserveInstanceId *string `json:"CacheReserveInstanceId,omitempty" xml:"CacheReserveInstanceId,omitempty"`
	// The switch status. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCacheReserveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCacheReserveResponseBody) GoString() string {
	return s.String()
}

func (s *GetCacheReserveResponseBody) GetCacheReserveInstanceId() *string {
	return s.CacheReserveInstanceId
}

func (s *GetCacheReserveResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetCacheReserveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCacheReserveResponseBody) SetCacheReserveInstanceId(v string) *GetCacheReserveResponseBody {
	s.CacheReserveInstanceId = &v
	return s
}

func (s *GetCacheReserveResponseBody) SetEnable(v string) *GetCacheReserveResponseBody {
	s.Enable = &v
	return s
}

func (s *GetCacheReserveResponseBody) SetRequestId(v string) *GetCacheReserveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCacheReserveResponseBody) Validate() error {
	return dara.Validate(s)
}
