// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFileDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckFileDeploymentResponseBody
	GetRequestId() *string
}

type CheckFileDeploymentResponseBody struct {
	// The request ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 86d5a5ac-0cc0-4c5f-a374-a15713b252ab
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckFileDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckFileDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFileDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckFileDeploymentResponseBody) SetRequestId(v string) *CheckFileDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFileDeploymentResponseBody) Validate() error {
	return dara.Validate(s)
}
