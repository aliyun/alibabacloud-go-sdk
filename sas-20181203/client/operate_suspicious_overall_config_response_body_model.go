// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSuspiciousOverallConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateSuspiciousOverallConfigResponseBody
	GetRequestId() *string
}

type OperateSuspiciousOverallConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C286491D-4A2F-589A-B63B-D2AD3DA9BD71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateSuspiciousOverallConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateSuspiciousOverallConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousOverallConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateSuspiciousOverallConfigResponseBody) SetRequestId(v string) *OperateSuspiciousOverallConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateSuspiciousOverallConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
