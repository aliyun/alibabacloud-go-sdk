// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFingerPrintTemplateDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetFingerPrintTemplateDescriptionResponseBody
	GetRequestId() *string
}

type SetFingerPrintTemplateDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BBD7DFD1-A5DE-51D9-8FD6-3BF54EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetFingerPrintTemplateDescriptionResponseBody) SetRequestId(v string) *SetFingerPrintTemplateDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
