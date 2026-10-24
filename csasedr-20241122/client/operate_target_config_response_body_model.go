// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateTargetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateTargetConfigResponseBody
	GetRequestId() *string
}

type OperateTargetConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E34B1849-64AA-5F2D-85A1-77F3B8E6BC90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateTargetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OperateTargetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateTargetConfigResponseBody) SetRequestId(v string) *OperateTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateTargetConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
