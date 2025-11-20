// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversaionSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetConversaionSpaceResponseBody
	GetRequestId() *string
	SetSpace(v *GetConversaionSpaceResponseBodySpace) *GetConversaionSpaceResponseBody
	GetSpace() *GetConversaionSpaceResponseBodySpace
	SetVendorRequestId(v string) *GetConversaionSpaceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetConversaionSpaceResponseBody
	GetVendorType() *string
}

type GetConversaionSpaceResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Space     *GetConversaionSpaceResponseBodySpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetConversaionSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConversaionSpaceResponseBody) GetSpace() *GetConversaionSpaceResponseBodySpace {
	return s.Space
}

func (s *GetConversaionSpaceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetConversaionSpaceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetConversaionSpaceResponseBody) SetRequestId(v string) *GetConversaionSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConversaionSpaceResponseBody) SetSpace(v *GetConversaionSpaceResponseBodySpace) *GetConversaionSpaceResponseBody {
	s.Space = v
	return s
}

func (s *GetConversaionSpaceResponseBody) SetVendorRequestId(v string) *GetConversaionSpaceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetConversaionSpaceResponseBody) SetVendorType(v string) *GetConversaionSpaceResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetConversaionSpaceResponseBody) Validate() error {
	if s.Space != nil {
		if err := s.Space.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConversaionSpaceResponseBodySpace struct {
	// example:
	//
	// ding1234xxxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 798xxxxx
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetConversaionSpaceResponseBodySpace) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceResponseBodySpace) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceResponseBodySpace) GetCorpId() *string {
	return s.CorpId
}

func (s *GetConversaionSpaceResponseBodySpace) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetConversaionSpaceResponseBodySpace) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetConversaionSpaceResponseBodySpace) GetSpaceId() *string {
	return s.SpaceId
}

func (s *GetConversaionSpaceResponseBodySpace) SetCorpId(v string) *GetConversaionSpaceResponseBodySpace {
	s.CorpId = &v
	return s
}

func (s *GetConversaionSpaceResponseBodySpace) SetCreateTime(v string) *GetConversaionSpaceResponseBodySpace {
	s.CreateTime = &v
	return s
}

func (s *GetConversaionSpaceResponseBodySpace) SetModifiedTime(v string) *GetConversaionSpaceResponseBodySpace {
	s.ModifiedTime = &v
	return s
}

func (s *GetConversaionSpaceResponseBodySpace) SetSpaceId(v string) *GetConversaionSpaceResponseBodySpace {
	s.SpaceId = &v
	return s
}

func (s *GetConversaionSpaceResponseBodySpace) Validate() error {
	return dara.Validate(s)
}
