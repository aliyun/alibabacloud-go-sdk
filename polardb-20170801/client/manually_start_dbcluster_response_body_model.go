// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManuallyStartDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ManuallyStartDBClusterResponseBody
	GetRequestId() *string
}

type ManuallyStartDBClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 73A85BAF-1039-4CDE-A83F-1A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManuallyStartDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManuallyStartDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ManuallyStartDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManuallyStartDBClusterResponseBody) SetRequestId(v string) *ManuallyStartDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManuallyStartDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
