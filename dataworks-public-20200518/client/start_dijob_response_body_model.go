// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartDIJobResponseBody
	GetRequestId() *string
}

type StartDIJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 999431B2-6013-577F-B684-36F7433C753B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDIJobResponseBody) SetRequestId(v string) *StartDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}
