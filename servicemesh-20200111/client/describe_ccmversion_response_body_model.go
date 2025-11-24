// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCCMVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCCMVersions(v map[string]*CCMVersionsValue) *DescribeCCMVersionResponseBody
	GetCCMVersions() map[string]*CCMVersionsValue
	SetRequestId(v string) *DescribeCCMVersionResponseBody
	GetRequestId() *string
}

type DescribeCCMVersionResponseBody struct {
	// The versions of the CCM component in all clusters on the data plane.
	CCMVersions map[string]*CCMVersionsValue `json:"CCMVersions,omitempty" xml:"CCMVersions,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCCMVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCCMVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCCMVersionResponseBody) GetCCMVersions() map[string]*CCMVersionsValue {
	return s.CCMVersions
}

func (s *DescribeCCMVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCCMVersionResponseBody) SetCCMVersions(v map[string]*CCMVersionsValue) *DescribeCCMVersionResponseBody {
	s.CCMVersions = v
	return s
}

func (s *DescribeCCMVersionResponseBody) SetRequestId(v string) *DescribeCCMVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCCMVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
