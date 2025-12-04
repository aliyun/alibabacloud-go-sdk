// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRestartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelRestartInstanceResponseBody
	GetRequestId() *string
}

type CancelRestartInstanceResponseBody struct {
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRestartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRestartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRestartInstanceResponseBody) SetRequestId(v string) *CancelRestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRestartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
