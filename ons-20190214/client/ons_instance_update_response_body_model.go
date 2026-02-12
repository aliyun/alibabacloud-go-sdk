// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsInstanceUpdateResponseBody
	GetRequestId() *string
}

type OnsInstanceUpdateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsInstanceUpdateResponseBody) SetRequestId(v string) *OnsInstanceUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsInstanceUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
