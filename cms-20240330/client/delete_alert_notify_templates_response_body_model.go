// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertNotifyTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlertNotifyTemplatesResponseBody
	GetRequestId() *string
}

type DeleteAlertNotifyTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlertNotifyTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertNotifyTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertNotifyTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertNotifyTemplatesResponseBody) SetRequestId(v string) *DeleteAlertNotifyTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertNotifyTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
