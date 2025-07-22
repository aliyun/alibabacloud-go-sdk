// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubnetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSubnetResponseBody
	GetRequestId() *string
	SetSubnetId(v string) *CreateSubnetResponseBody
	GetSubnetId() *string
}

type CreateSubnetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubnetId  *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s CreateSubnetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubnetResponseBody) GetSubnetId() *string {
	return s.SubnetId
}

func (s *CreateSubnetResponseBody) SetRequestId(v string) *CreateSubnetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubnetResponseBody) SetSubnetId(v string) *CreateSubnetResponseBody {
	s.SubnetId = &v
	return s
}

func (s *CreateSubnetResponseBody) Validate() error {
	return dara.Validate(s)
}
