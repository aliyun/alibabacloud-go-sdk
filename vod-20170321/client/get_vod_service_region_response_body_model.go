// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodServiceRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetVodServiceRegionResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetVodServiceRegionResponseBody
	GetRequestId() *string
}

type GetVodServiceRegionResponseBody struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVodServiceRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVodServiceRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetVodServiceRegionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVodServiceRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVodServiceRegionResponseBody) SetRegionId(v string) *GetVodServiceRegionResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVodServiceRegionResponseBody) SetRequestId(v string) *GetVodServiceRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVodServiceRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
