// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResubmitConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyResubmitConfigResponseBody
	GetRequestId() *string
}

type ModifyResubmitConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResubmitConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResubmitConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResubmitConfigResponseBody) SetRequestId(v string) *ModifyResubmitConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResubmitConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
