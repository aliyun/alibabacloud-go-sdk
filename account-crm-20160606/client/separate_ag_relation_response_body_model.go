// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeparateAgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SeparateAgRelationResponseBody
	GetCode() *string
	SetMessage(v string) *SeparateAgRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *SeparateAgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SeparateAgRelationResponseBody
	GetSuccess() *bool
}

type SeparateAgRelationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SeparateAgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SeparateAgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *SeparateAgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *SeparateAgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SeparateAgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SeparateAgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SeparateAgRelationResponseBody) SetCode(v string) *SeparateAgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *SeparateAgRelationResponseBody) SetMessage(v string) *SeparateAgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *SeparateAgRelationResponseBody) SetRequestId(v string) *SeparateAgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SeparateAgRelationResponseBody) SetSuccess(v bool) *SeparateAgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *SeparateAgRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
