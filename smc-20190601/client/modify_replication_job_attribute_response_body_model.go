// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReplicationJobAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyReplicationJobAttributeResponseBody
	GetRequestId() *string
}

type ModifyReplicationJobAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1C488B66-B819-4D14-8711-C4EAAA13AC01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReplicationJobAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReplicationJobAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReplicationJobAttributeResponseBody) SetRequestId(v string) *ModifyReplicationJobAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReplicationJobAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
