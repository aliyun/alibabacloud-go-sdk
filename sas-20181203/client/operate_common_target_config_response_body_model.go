// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCommonTargetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateCommonTargetConfigResponseBody
	GetRequestId() *string
}

type OperateCommonTargetConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 95D45C44-4F53-5ED2-8E12-7D134564B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateCommonTargetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateCommonTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OperateCommonTargetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateCommonTargetConfigResponseBody) SetRequestId(v string) *OperateCommonTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateCommonTargetConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
