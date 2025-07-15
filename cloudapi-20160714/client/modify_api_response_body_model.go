// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApiResponseBody
	GetRequestId() *string
}

type ModifyApiResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApiResponseBody) SetRequestId(v string) *ModifyApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiResponseBody) Validate() error {
	return dara.Validate(s)
}
