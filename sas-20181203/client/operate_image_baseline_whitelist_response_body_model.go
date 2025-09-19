// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateImageBaselineWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateImageBaselineWhitelistResponseBody
	GetRequestId() *string
}

type OperateImageBaselineWhitelistResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 75AD186B-B46A-56CC-BE35-987ADDF6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateImageBaselineWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateImageBaselineWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *OperateImageBaselineWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateImageBaselineWhitelistResponseBody) SetRequestId(v string) *OperateImageBaselineWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateImageBaselineWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
