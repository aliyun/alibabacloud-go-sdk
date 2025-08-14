// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultCustomTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetDefaultCustomTemplateResponseBody
	GetSuccess() *bool
}

type SetDefaultCustomTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDefaultCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultCustomTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDefaultCustomTemplateResponseBody) SetRequestId(v string) *SetDefaultCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultCustomTemplateResponseBody) SetSuccess(v bool) *SetDefaultCustomTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *SetDefaultCustomTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
