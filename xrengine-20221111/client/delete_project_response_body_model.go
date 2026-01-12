// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteProjectResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteProjectResponseBody
	GetSuccess() *bool
}

type DeleteProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProjectResponseBody) SetCode(v string) *DeleteProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProjectResponseBody) SetMessage(v string) *DeleteProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
