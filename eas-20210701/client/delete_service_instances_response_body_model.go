// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteServiceInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceInstancesResponseBody
	GetRequestId() *string
}

type DeleteServiceInstancesResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Instance(s) [foo-rdsb,foo-rdsa]  for service [foo] in region [cn-shanghai] was deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceInstancesResponseBody) SetMessage(v string) *DeleteServiceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
