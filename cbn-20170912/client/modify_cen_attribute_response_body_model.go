// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCenAttributeResponseBody
	GetRequestId() *string
}

type ModifyCenAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 13526224-5780-4426-8BDF-BC8B08700F22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCenAttributeResponseBody) SetRequestId(v string) *ModifyCenAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCenAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
