// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpfilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpFilterId(v string) *AddIpfilterResponseBody
	GetIpFilterId() *string
	SetRequestId(v string) *AddIpfilterResponseBody
	GetRequestId() *string
}

type AddIpfilterResponseBody struct {
	// ID corresponding to the IP
	//
	// example:
	//
	// 10795
	IpFilterId *string `json:"IpFilterId,omitempty" xml:"IpFilterId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 0E9282E8-DC08-5445-8FB0-B9F0CA28B249
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpfilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIpfilterResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpfilterResponseBody) GetIpFilterId() *string {
	return s.IpFilterId
}

func (s *AddIpfilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIpfilterResponseBody) SetIpFilterId(v string) *AddIpfilterResponseBody {
	s.IpFilterId = &v
	return s
}

func (s *AddIpfilterResponseBody) SetRequestId(v string) *AddIpfilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIpfilterResponseBody) Validate() error {
	return dara.Validate(s)
}
