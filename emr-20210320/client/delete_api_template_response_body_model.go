// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApiTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteApiTemplateResponseBody
	GetSuccess() *bool
}

type DeleteApiTemplateResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deprecated
	//
	// Whether the call was successful: - true: Call succeeded - false: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteApiTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteApiTemplateResponseBody) SetRequestId(v string) *DeleteApiTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiTemplateResponseBody) SetSuccess(v bool) *DeleteApiTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApiTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
