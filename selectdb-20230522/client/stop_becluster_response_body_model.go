// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBEClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopBEClusterResponseBody
	GetRequestId() *string
}

type StopBEClusterResponseBody struct {
	// example:
	//
	// BC854513-E85E-54F3-9842-B9CCD3308CDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopBEClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopBEClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopBEClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopBEClusterResponseBody) SetRequestId(v string) *StopBEClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopBEClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
