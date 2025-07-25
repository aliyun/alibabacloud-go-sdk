// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAddressPoolAvailableConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeInfos(v *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) *DescribeDnsGtmAddressPoolAvailableConfigResponseBody
	GetAttributeInfos() *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos
	SetRequestId(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBody
	GetRequestId() *string
}

type DescribeDnsGtmAddressPoolAvailableConfigResponseBody struct {
	// The supported source regions.
	AttributeInfos *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos `json:"AttributeInfos,omitempty" xml:"AttributeInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 199C3699-9A7B-41A1-BB5A-F1E862D3CB38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) GetAttributeInfos() *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos {
	return s.AttributeInfos
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) SetAttributeInfos(v *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) *DescribeDnsGtmAddressPoolAvailableConfigResponseBody {
	s.AttributeInfos = v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos struct {
	AttributeInfo []*DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) GetAttributeInfo() []*DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo {
	return s.AttributeInfo
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) SetAttributeInfo(v []*DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos {
	s.AttributeInfo = v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo struct {
	// The parent line code of the source region.
	//
	// example:
	//
	// telecom
	FatherCode *string `json:"FatherCode,omitempty" xml:"FatherCode,omitempty"`
	// The code of the source region group.
	//
	// example:
	//
	// default
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the request source group.
	//
	// example:
	//
	// global
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The line code of the source region.
	//
	// example:
	//
	// default
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The line name of the source region.
	//
	// example:
	//
	// global
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) GetFatherCode() *string {
	return s.FatherCode
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) GetLineName() *string {
	return s.LineName
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) SetFatherCode(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo {
	s.FatherCode = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) SetGroupCode(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo {
	s.GroupCode = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) SetGroupName(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) SetLineCode(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) SetLineName(v string) *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponseBodyAttributeInfosAttributeInfo) Validate() error {
	return dara.Validate(s)
}
