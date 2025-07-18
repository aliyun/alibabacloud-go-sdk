// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountDescriptionResponseBody
	GetRequestId() *string
}

type ModifyAccountDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 99BBBD5E-B5D8-4FC8-B8BF-FB1A0A38BBA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
