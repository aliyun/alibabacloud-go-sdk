// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoamClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RoamClientUserResponseBody
	GetRequestId() *string
}

type RoamClientUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3200E8A3-563F-4FFC-8BDB-0F1263FA69E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RoamClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RoamClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *RoamClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RoamClientUserResponseBody) SetRequestId(v string) *RoamClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RoamClientUserResponseBody) Validate() error {
	return dara.Validate(s)
}
