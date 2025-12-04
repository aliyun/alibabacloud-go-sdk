// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartInstanceResponseBody
	GetRequestId() *string
}

type RestartInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
