// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDasInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDasInstanceConfigResponseBody
	GetRequestId() *string
}

type ModifyDasInstanceConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C3C247D4-1643-4C5D-87C2-C829543FC626
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDasInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDasInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDasInstanceConfigResponseBody) SetRequestId(v string) *ModifyDasInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDasInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
