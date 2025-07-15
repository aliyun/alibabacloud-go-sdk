// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVpcAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetVpcAccessResponseBody
	GetRequestId() *string
	SetVpcAccessId(v string) *SetVpcAccessResponseBody
	GetVpcAccessId() *string
}

type SetVpcAccessResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VPC access authorization.
	//
	// example:
	//
	// 4c68e061860f441ab72af7404137440e
	VpcAccessId *string `json:"VpcAccessId,omitempty" xml:"VpcAccessId,omitempty"`
}

func (s SetVpcAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVpcAccessResponseBody) GoString() string {
	return s.String()
}

func (s *SetVpcAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVpcAccessResponseBody) GetVpcAccessId() *string {
	return s.VpcAccessId
}

func (s *SetVpcAccessResponseBody) SetRequestId(v string) *SetVpcAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVpcAccessResponseBody) SetVpcAccessId(v string) *SetVpcAccessResponseBody {
	s.VpcAccessId = &v
	return s
}

func (s *SetVpcAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
