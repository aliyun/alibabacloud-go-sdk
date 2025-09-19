// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateVulsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateVulsResponseBody
	GetRequestId() *string
}

type OperateVulsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C2677612-7207-4AEB-BD48-8BA528F86777
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateVulsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateVulsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateVulsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateVulsResponseBody) SetRequestId(v string) *OperateVulsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateVulsResponseBody) Validate() error {
	return dara.Validate(s)
}
