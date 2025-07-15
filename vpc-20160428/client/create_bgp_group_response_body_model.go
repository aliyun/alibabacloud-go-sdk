// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBgpGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpGroupId(v string) *CreateBgpGroupResponseBody
	GetBgpGroupId() *string
	SetRequestId(v string) *CreateBgpGroupResponseBody
	GetRequestId() *string
}

type CreateBgpGroupResponseBody struct {
	// The ID of the BGP group.
	//
	// example:
	//
	// bgpg-bp1k25cyp26cllath****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C1221A1F-2ACD-4592-8F27-474E02883159
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBgpGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBgpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBgpGroupResponseBody) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *CreateBgpGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBgpGroupResponseBody) SetBgpGroupId(v string) *CreateBgpGroupResponseBody {
	s.BgpGroupId = &v
	return s
}

func (s *CreateBgpGroupResponseBody) SetRequestId(v string) *CreateBgpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBgpGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
