// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartDBClusterResponseBody
	GetRequestId() *string
}

type RestartDBClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBClusterResponseBody) SetRequestId(v string) *RestartDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
