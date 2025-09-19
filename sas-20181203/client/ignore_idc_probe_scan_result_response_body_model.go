// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreIdcProbeScanResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *IgnoreIdcProbeScanResultResponseBody
	GetRequestId() *string
}

type IgnoreIdcProbeScanResultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E30282D8-AB2D-5EB1-998B-2DDFA948D49D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreIdcProbeScanResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IgnoreIdcProbeScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreIdcProbeScanResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IgnoreIdcProbeScanResultResponseBody) SetRequestId(v string) *IgnoreIdcProbeScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *IgnoreIdcProbeScanResultResponseBody) Validate() error {
	return dara.Validate(s)
}
