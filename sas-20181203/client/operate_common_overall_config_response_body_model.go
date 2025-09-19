// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCommonOverallConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateCommonOverallConfigResponseBody
	GetRequestId() *string
}

type OperateCommonOverallConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateCommonOverallConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateCommonOverallConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OperateCommonOverallConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateCommonOverallConfigResponseBody) SetRequestId(v string) *OperateCommonOverallConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateCommonOverallConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
