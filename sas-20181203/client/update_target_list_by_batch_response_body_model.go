// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTargetListByBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTargetListByBatchResponseBody
	GetRequestId() *string
}

type UpdateTargetListByBatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBBEB173-1F43-505F-A876-C03ECD******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTargetListByBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTargetListByBatchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTargetListByBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTargetListByBatchResponseBody) SetRequestId(v string) *UpdateTargetListByBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTargetListByBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
