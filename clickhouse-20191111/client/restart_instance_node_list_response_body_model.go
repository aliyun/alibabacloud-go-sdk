// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceNodeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartInstanceNodeListResponseBody
	GetRequestId() *string
}

type RestartInstanceNodeListResponseBody struct {
	// example:
	//
	// 36E3AB2E-E0DA-5C7B-8E29-89EE44926515
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartInstanceNodeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceNodeListResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceNodeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstanceNodeListResponseBody) SetRequestId(v string) *RestartInstanceNodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceNodeListResponseBody) Validate() error {
	return dara.Validate(s)
}
