// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageRegionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionList(v *GetStorageRegionListResponseBodyRegionList) *GetStorageRegionListResponseBody
	GetRegionList() *GetStorageRegionListResponseBodyRegionList
	SetRequestId(v string) *GetStorageRegionListResponseBody
	GetRequestId() *string
}

type GetStorageRegionListResponseBody struct {
	RegionList *GetStorageRegionListResponseBodyRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Struct"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetStorageRegionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageRegionListResponseBody) GetRegionList() *GetStorageRegionListResponseBodyRegionList {
	return s.RegionList
}

func (s *GetStorageRegionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageRegionListResponseBody) SetRegionList(v *GetStorageRegionListResponseBodyRegionList) *GetStorageRegionListResponseBody {
	s.RegionList = v
	return s
}

func (s *GetStorageRegionListResponseBody) SetRequestId(v string) *GetStorageRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageRegionListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStorageRegionListResponseBodyRegionList struct {
	Region []*GetStorageRegionListResponseBodyRegionListRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s GetStorageRegionListResponseBodyRegionList) String() string {
	return dara.Prettify(s)
}

func (s GetStorageRegionListResponseBodyRegionList) GoString() string {
	return s.String()
}

func (s *GetStorageRegionListResponseBodyRegionList) GetRegion() []*GetStorageRegionListResponseBodyRegionListRegion {
	return s.Region
}

func (s *GetStorageRegionListResponseBodyRegionList) SetRegion(v []*GetStorageRegionListResponseBodyRegionListRegion) *GetStorageRegionListResponseBodyRegionList {
	s.Region = v
	return s
}

func (s *GetStorageRegionListResponseBodyRegionList) Validate() error {
	return dara.Validate(s)
}

type GetStorageRegionListResponseBodyRegionListRegion struct {
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
}

func (s GetStorageRegionListResponseBodyRegionListRegion) String() string {
	return dara.Prettify(s)
}

func (s GetStorageRegionListResponseBodyRegionListRegion) GoString() string {
	return s.String()
}

func (s *GetStorageRegionListResponseBodyRegionListRegion) GetRegionID() *string {
	return s.RegionID
}

func (s *GetStorageRegionListResponseBodyRegionListRegion) SetRegionID(v string) *GetStorageRegionListResponseBodyRegionListRegion {
	s.RegionID = &v
	return s
}

func (s *GetStorageRegionListResponseBodyRegionListRegion) Validate() error {
	return dara.Validate(s)
}
