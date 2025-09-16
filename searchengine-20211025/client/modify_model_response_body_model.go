// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyModelResponseBody
	GetRequestId() *string
}

type ModifyModelResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ModifyModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyModelResponseBody) SetRequestId(v string) *ModifyModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyModelResponseBody) Validate() error {
	return dara.Validate(s)
}
