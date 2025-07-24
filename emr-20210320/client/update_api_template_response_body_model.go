// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApiTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApiTemplateResponseBody
	GetSuccess() *bool
}

type UpdateApiTemplateResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deprecated
	//
	// Template ID (to be deprecated).
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateApiTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApiTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApiTemplateResponseBody) SetRequestId(v string) *UpdateApiTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApiTemplateResponseBody) SetSuccess(v bool) *UpdateApiTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApiTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
