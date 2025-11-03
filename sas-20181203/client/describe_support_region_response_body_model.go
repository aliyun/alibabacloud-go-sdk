// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSupportRegionResponseBody
	GetRequestId() *string
	SetSupportRegion(v []*string) *DescribeSupportRegionResponseBody
	GetSupportRegion() []*string
}

type DescribeSupportRegionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2C0699D3-4107-5A46-A4C4-E129A5967788
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of supported regions
	SupportRegion []*string `json:"SupportRegion,omitempty" xml:"SupportRegion,omitempty" type:"Repeated"`
}

func (s DescribeSupportRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportRegionResponseBody) GetSupportRegion() []*string {
	return s.SupportRegion
}

func (s *DescribeSupportRegionResponseBody) SetRequestId(v string) *DescribeSupportRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportRegionResponseBody) SetSupportRegion(v []*string) *DescribeSupportRegionResponseBody {
	s.SupportRegion = v
	return s
}

func (s *DescribeSupportRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
