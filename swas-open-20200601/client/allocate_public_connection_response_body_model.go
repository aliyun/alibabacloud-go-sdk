// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicConnection(v string) *AllocatePublicConnectionResponseBody
	GetPublicConnection() *string
	SetRequestId(v string) *AllocatePublicConnectionResponseBody
	GetRequestId() *string
}

type AllocatePublicConnectionResponseBody struct {
	// The public endpoint.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****.mysql.rds.aliyuncs.com
	PublicConnection *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionResponseBody) GetPublicConnection() *string {
	return s.PublicConnection
}

func (s *AllocatePublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocatePublicConnectionResponseBody) SetPublicConnection(v string) *AllocatePublicConnectionResponseBody {
	s.PublicConnection = &v
	return s
}

func (s *AllocatePublicConnectionResponseBody) SetRequestId(v string) *AllocatePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocatePublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
