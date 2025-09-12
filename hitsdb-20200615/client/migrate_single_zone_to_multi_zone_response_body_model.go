// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateSingleZoneToMultiZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *MigrateSingleZoneToMultiZoneResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *MigrateSingleZoneToMultiZoneResponseBody
	GetRequestId() *string
}

type MigrateSingleZoneToMultiZoneResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateSingleZoneToMultiZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateSingleZoneToMultiZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateSingleZoneToMultiZoneResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *MigrateSingleZoneToMultiZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateSingleZoneToMultiZoneResponseBody) SetAccessDeniedDetail(v string) *MigrateSingleZoneToMultiZoneResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneResponseBody) SetRequestId(v string) *MigrateSingleZoneToMultiZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
