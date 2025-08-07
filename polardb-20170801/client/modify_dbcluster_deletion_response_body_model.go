// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDeletionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterDeletionResponseBody
	GetRequestId() *string
}

type ModifyDBClusterDeletionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 24C80BD8-C710-4138-893A-D2AFED4FC13D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterDeletionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDeletionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterDeletionResponseBody) SetRequestId(v string) *ModifyDBClusterDeletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterDeletionResponseBody) Validate() error {
	return dara.Validate(s)
}
