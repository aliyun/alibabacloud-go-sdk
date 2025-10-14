// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAddrAttributeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddr(v *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) *DescribeDnsGtmAddrAttributeInfoResponseBody
	GetAddr() *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr
	SetRequestId(v string) *DescribeDnsGtmAddrAttributeInfoResponseBody
	GetRequestId() *string
}

type DescribeDnsGtmAddrAttributeInfoResponseBody struct {
	// The address in the address pool.
	Addr *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBody) GetAddr() *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr {
	return s.Addr
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBody) SetAddr(v *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) *DescribeDnsGtmAddrAttributeInfoResponseBody {
	s.Addr = v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBody) SetRequestId(v string) *DescribeDnsGtmAddrAttributeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBody) Validate() error {
	if s.Addr != nil {
		if err := s.Addr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmAddrAttributeInfoResponseBodyAddr struct {
	Addr []*DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) GetAddr() []*DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr {
	return s.Addr
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) SetAddr(v []*DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr {
	s.Addr = v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddr) Validate() error {
	if s.Addr != nil {
		for _, item := range s.Addr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr struct {
	// The address in the address pool.
	//
	// example:
	//
	// 1.1.1.1
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// The information about the source region of the address.
	AttributeInfo *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) GetAddr() *string {
	return s.Addr
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) GetAttributeInfo() *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo {
	return s.AttributeInfo
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) SetAddr(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr {
	s.Addr = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) SetAttributeInfo(v *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr {
	s.AttributeInfo = v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddr) Validate() error {
	if s.AttributeInfo != nil {
		if err := s.AttributeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo struct {
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
	// DEFAULT
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the source region group.
	//
	// example:
	//
	// Global
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
	// Global
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) GetFatherCode() *string {
	return s.FatherCode
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) GetLineName() *string {
	return s.LineName
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) SetFatherCode(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo {
	s.FatherCode = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) SetGroupCode(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo {
	s.GroupCode = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) SetGroupName(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) SetLineCode(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) SetLineName(v string) *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoResponseBodyAddrAddrAttributeInfo) Validate() error {
	return dara.Validate(s)
}
