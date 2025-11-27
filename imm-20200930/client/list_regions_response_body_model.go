// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*RegionType) *ListRegionsResponseBody
	GetRegions() []*RegionType
	SetRequestId(v string) *ListRegionsResponseBody
	GetRequestId() *string
}

type ListRegionsResponseBody struct {
	// List of Regions.
	Regions []*RegionType `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Request ID for the interface.
	//
	// example:
	//
	// 7F7D235C-76FF-4B65-800C-8238AE3F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) GetRegions() []*RegionType {
	return s.Regions
}

func (s *ListRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionsResponseBody) SetRegions(v []*RegionType) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
