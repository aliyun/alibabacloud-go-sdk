// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTemplateResourcesResponseBody
	GetRequestId() *string
}

type ModifyTemplateResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CF708F2F-FFB0-54D4-B1E0-B84A7CEBFB60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTemplateResourcesResponseBody) SetRequestId(v string) *ModifyTemplateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTemplateResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
