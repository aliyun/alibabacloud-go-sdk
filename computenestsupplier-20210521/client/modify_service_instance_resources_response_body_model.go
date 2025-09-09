// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceInstanceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyServiceInstanceResourcesResponseBody
	GetRequestId() *string
}

type ModifyServiceInstanceResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServiceInstanceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyServiceInstanceResourcesResponseBody) SetRequestId(v string) *ModifyServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyServiceInstanceResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
