// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsGroupDeleteResponseBody
	GetRequestId() *string
}

type OnsGroupDeleteResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsGroupDeleteResponseBody) SetRequestId(v string) *OnsGroupDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsGroupDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
