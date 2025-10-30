// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountFactoryBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v string) *CreateAccountFactoryBaselineResponseBody
	GetBaselineId() *string
	SetRequestId(v string) *CreateAccountFactoryBaselineResponseBody
	GetRequestId() *string
}

type CreateAccountFactoryBaselineResponseBody struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1e6ixtiwupap8m****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5592E2E-0FC4-557C-B989-DF229B5EBE13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountFactoryBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineResponseBody) GetBaselineId() *string {
	return s.BaselineId
}

func (s *CreateAccountFactoryBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountFactoryBaselineResponseBody) SetBaselineId(v string) *CreateAccountFactoryBaselineResponseBody {
	s.BaselineId = &v
	return s
}

func (s *CreateAccountFactoryBaselineResponseBody) SetRequestId(v string) *CreateAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountFactoryBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}
