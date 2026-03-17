// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagPortRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagPortRoleResponseBody
	GetRequestId() *string
}

type ModifySagPortRoleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3200E8A3-563F-4FFC-8BDB-0F1263FA69E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagPortRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagPortRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagPortRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagPortRoleResponseBody) SetRequestId(v string) *ModifySagPortRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagPortRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
