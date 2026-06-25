// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTemplateResponseBody
	GetRequestId() *string
}

type ModifyTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C5AB05E9-D3CA-582E-9538-8C6C554040F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTemplateResponseBody) SetRequestId(v string) *ModifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
