// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *AddIntegrationResponseBody
	GetData() *string
	SetRequestId(v string) *AddIntegrationResponseBody
	GetRequestId() *string
}

type AddIntegrationResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponseBody) GetData() *string {
	return s.Data
}

func (s *AddIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIntegrationResponseBody) SetData(v string) *AddIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *AddIntegrationResponseBody) SetRequestId(v string) *AddIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}
