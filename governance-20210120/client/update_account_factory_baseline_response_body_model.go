// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountFactoryBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAccountFactoryBaselineResponseBody
	GetRequestId() *string
}

type UpdateAccountFactoryBaselineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C18A891D-7B04-51A1-AAC6-201727A361CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountFactoryBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccountFactoryBaselineResponseBody) SetRequestId(v string) *UpdateAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountFactoryBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}
