// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHaVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipIds(v []*string) *CreateHaVipResponseBody
	GetHaVipIds() []*string
	SetRequestId(v string) *CreateHaVipResponseBody
	GetRequestId() *string
}

type CreateHaVipResponseBody struct {
	// The IDs of the HAVIPs.
	HaVipIds []*string `json:"HaVipIds,omitempty" xml:"HaVipIds,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHaVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHaVipResponseBody) GetHaVipIds() []*string {
	return s.HaVipIds
}

func (s *CreateHaVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHaVipResponseBody) SetHaVipIds(v []*string) *CreateHaVipResponseBody {
	s.HaVipIds = v
	return s
}

func (s *CreateHaVipResponseBody) SetRequestId(v string) *CreateHaVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHaVipResponseBody) Validate() error {
	return dara.Validate(s)
}
