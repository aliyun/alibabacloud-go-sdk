// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSuspiciousTargetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateSuspiciousTargetConfigResponseBody
	GetRequestId() *string
}

type OperateSuspiciousTargetConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// ABCD-PSD2-5256-1DSA-4222-JHBN
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateSuspiciousTargetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateSuspiciousTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousTargetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateSuspiciousTargetConfigResponseBody) SetRequestId(v string) *OperateSuspiciousTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateSuspiciousTargetConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
