// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubnetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSubnetResponseBody
	GetRequestId() *string
	SetSubnetId(v string) *DeleteSubnetResponseBody
	GetSubnetId() *string
}

type DeleteSubnetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubnetId  *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s DeleteSubnetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubnetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubnetResponseBody) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DeleteSubnetResponseBody) SetRequestId(v string) *DeleteSubnetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubnetResponseBody) SetSubnetId(v string) *DeleteSubnetResponseBody {
	s.SubnetId = &v
	return s
}

func (s *DeleteSubnetResponseBody) Validate() error {
	return dara.Validate(s)
}
