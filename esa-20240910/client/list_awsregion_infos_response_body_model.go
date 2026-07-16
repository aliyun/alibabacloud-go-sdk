// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAWSRegionInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*ListAWSRegionInfosResponseBodyRegions) *ListAWSRegionInfosResponseBody
	GetRegions() []*ListAWSRegionInfosResponseBodyRegions
	SetRequestId(v string) *ListAWSRegionInfosResponseBody
	GetRequestId() *string
}

type ListAWSRegionInfosResponseBody struct {
	// The region information.
	Regions []*ListAWSRegionInfosResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAWSRegionInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAWSRegionInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListAWSRegionInfosResponseBody) GetRegions() []*ListAWSRegionInfosResponseBodyRegions {
	return s.Regions
}

func (s *ListAWSRegionInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAWSRegionInfosResponseBody) SetRegions(v []*ListAWSRegionInfosResponseBodyRegions) *ListAWSRegionInfosResponseBody {
	s.Regions = v
	return s
}

func (s *ListAWSRegionInfosResponseBody) SetRequestId(v string) *ListAWSRegionInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAWSRegionInfosResponseBody) Validate() error {
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

type ListAWSRegionInfosResponseBodyRegions struct {
	// The region code.
	//
	// example:
	//
	// us-east-2
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The region name.
	//
	// example:
	//
	// US East (Ohio)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAWSRegionInfosResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListAWSRegionInfosResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListAWSRegionInfosResponseBodyRegions) GetCode() *string {
	return s.Code
}

func (s *ListAWSRegionInfosResponseBodyRegions) GetName() *string {
	return s.Name
}

func (s *ListAWSRegionInfosResponseBodyRegions) SetCode(v string) *ListAWSRegionInfosResponseBodyRegions {
	s.Code = &v
	return s
}

func (s *ListAWSRegionInfosResponseBodyRegions) SetName(v string) *ListAWSRegionInfosResponseBodyRegions {
	s.Name = &v
	return s
}

func (s *ListAWSRegionInfosResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
