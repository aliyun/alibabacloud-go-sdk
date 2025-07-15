// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDeployApplicationGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ContinueDeployApplicationGroupResponseBody
	GetRequestId() *string
}

type ContinueDeployApplicationGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8AF4800A-A316-589A-90C4-313B1FEEB084
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinueDeployApplicationGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeployApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployApplicationGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinueDeployApplicationGroupResponseBody) SetRequestId(v string) *ContinueDeployApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueDeployApplicationGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
