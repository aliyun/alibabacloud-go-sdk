// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAIServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CreateAIServiceResponseBody
	GetServiceId() *string
}

type CreateAIServiceResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateAIServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAIServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAIServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateAIServiceResponseBody) SetRequestId(v string) *CreateAIServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIServiceResponseBody) SetServiceId(v string) *CreateAIServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateAIServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
