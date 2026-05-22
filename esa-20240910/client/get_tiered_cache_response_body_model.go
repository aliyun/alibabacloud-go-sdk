// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTieredCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCacheArchitectureMode(v string) *GetTieredCacheResponseBody
	GetCacheArchitectureMode() *string
	SetRequestId(v string) *GetTieredCacheResponseBody
	GetRequestId() *string
}

type GetTieredCacheResponseBody struct {
	CacheArchitectureMode *string `json:"CacheArchitectureMode,omitempty" xml:"CacheArchitectureMode,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTieredCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTieredCacheResponseBody) GoString() string {
	return s.String()
}

func (s *GetTieredCacheResponseBody) GetCacheArchitectureMode() *string {
	return s.CacheArchitectureMode
}

func (s *GetTieredCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTieredCacheResponseBody) SetCacheArchitectureMode(v string) *GetTieredCacheResponseBody {
	s.CacheArchitectureMode = &v
	return s
}

func (s *GetTieredCacheResponseBody) SetRequestId(v string) *GetTieredCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTieredCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
