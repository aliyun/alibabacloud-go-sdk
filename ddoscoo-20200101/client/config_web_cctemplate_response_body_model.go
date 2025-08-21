// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebCCTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigWebCCTemplateResponseBody
	GetRequestId() *string
}

type ConfigWebCCTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigWebCCTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebCCTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigWebCCTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigWebCCTemplateResponseBody) SetRequestId(v string) *ConfigWebCCTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigWebCCTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
