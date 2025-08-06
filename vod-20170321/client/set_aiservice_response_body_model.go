// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAIServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedTypes(v string) *SetAIServiceResponseBody
	GetFailedTypes() *string
	SetRequestId(v string) *SetAIServiceResponseBody
	GetRequestId() *string
	SetSuccessTypes(v string) *SetAIServiceResponseBody
	GetSuccessTypes() *string
}

type SetAIServiceResponseBody struct {
	FailedTypes  *string `json:"FailedTypes,omitempty" xml:"FailedTypes,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessTypes *string `json:"SuccessTypes,omitempty" xml:"SuccessTypes,omitempty"`
}

func (s SetAIServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAIServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SetAIServiceResponseBody) GetFailedTypes() *string {
	return s.FailedTypes
}

func (s *SetAIServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAIServiceResponseBody) GetSuccessTypes() *string {
	return s.SuccessTypes
}

func (s *SetAIServiceResponseBody) SetFailedTypes(v string) *SetAIServiceResponseBody {
	s.FailedTypes = &v
	return s
}

func (s *SetAIServiceResponseBody) SetRequestId(v string) *SetAIServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAIServiceResponseBody) SetSuccessTypes(v string) *SetAIServiceResponseBody {
	s.SuccessTypes = &v
	return s
}

func (s *SetAIServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
