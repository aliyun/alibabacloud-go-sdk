// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsInstanceDeleteResponseBody
	GetRequestId() *string
}

type OnsInstanceDeleteResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsInstanceDeleteResponseBody) SetRequestId(v string) *OnsInstanceDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsInstanceDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
