// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMonitorGroupResponseBody
	GetRequestId() *string
}

type DeleteMonitorGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C7E2CBAA-47FF-569F-AF12-B14B80E25422
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMonitorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMonitorGroupResponseBody) SetRequestId(v string) *DeleteMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
