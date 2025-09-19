// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterCnnfStatusUserConfirmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterCnnfStatusUserConfirmResponseBody
	GetRequestId() *string
}

type ModifyClusterCnnfStatusUserConfirmResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4FFBEDBD-FA63-5213-9103-306519EE4857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterCnnfStatusUserConfirmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterCnnfStatusUserConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterCnnfStatusUserConfirmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterCnnfStatusUserConfirmResponseBody) SetRequestId(v string) *ModifyClusterCnnfStatusUserConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterCnnfStatusUserConfirmResponseBody) Validate() error {
	return dara.Validate(s)
}
