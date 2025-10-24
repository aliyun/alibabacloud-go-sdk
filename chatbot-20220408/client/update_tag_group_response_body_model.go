// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTagGroupResponseBody
	GetSuccess() *bool
}

type UpdateTagGroupResponseBody struct {
	// example:
	//
	// xxxXxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTagGroupResponseBody) SetRequestId(v string) *UpdateTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTagGroupResponseBody) SetSuccess(v bool) *UpdateTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
