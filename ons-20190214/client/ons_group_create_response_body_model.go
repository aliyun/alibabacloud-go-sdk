// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsGroupCreateResponseBody
	GetRequestId() *string
}

type OnsGroupCreateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsGroupCreateResponseBody) SetRequestId(v string) *OnsGroupCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsGroupCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
