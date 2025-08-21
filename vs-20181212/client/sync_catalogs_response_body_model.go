// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SyncCatalogsResponseBody
	GetId() *string
	SetRequestId(v string) *SyncCatalogsResponseBody
	GetRequestId() *string
}

type SyncCatalogsResponseBody struct {
	// example:
	//
	// 3238****739092996
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *SyncCatalogsResponseBody) GetId() *string {
	return s.Id
}

func (s *SyncCatalogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncCatalogsResponseBody) SetId(v string) *SyncCatalogsResponseBody {
	s.Id = &v
	return s
}

func (s *SyncCatalogsResponseBody) SetRequestId(v string) *SyncCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncCatalogsResponseBody) Validate() error {
	return dara.Validate(s)
}
