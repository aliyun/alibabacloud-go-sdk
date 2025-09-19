// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateImageVulResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateImageVulResponseBody
	GetRequestId() *string
}

type OperateImageVulResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 486F2228-438A-544A-A533-433F943C15CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateImageVulResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateImageVulResponseBody) GoString() string {
	return s.String()
}

func (s *OperateImageVulResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateImageVulResponseBody) SetRequestId(v string) *OperateImageVulResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateImageVulResponseBody) Validate() error {
	return dara.Validate(s)
}
