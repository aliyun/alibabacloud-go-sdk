// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGdnStandbyMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateGdnStandbyMemberResponseBody
	GetDBInstanceName() *string
	SetOrderId(v string) *CreateGdnStandbyMemberResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateGdnStandbyMemberResponseBody
	GetRequestId() *string
}

type CreateGdnStandbyMemberResponseBody struct {
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 12345
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGdnStandbyMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGdnStandbyMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGdnStandbyMemberResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateGdnStandbyMemberResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateGdnStandbyMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGdnStandbyMemberResponseBody) SetDBInstanceName(v string) *CreateGdnStandbyMemberResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CreateGdnStandbyMemberResponseBody) SetOrderId(v string) *CreateGdnStandbyMemberResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateGdnStandbyMemberResponseBody) SetRequestId(v string) *CreateGdnStandbyMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGdnStandbyMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
