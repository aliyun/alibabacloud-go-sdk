// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDataLimitResponseBody
	GetRequestId() *string
}

type ModifyDataLimitResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataLimitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataLimitResponseBody) SetRequestId(v string) *ModifyDataLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
