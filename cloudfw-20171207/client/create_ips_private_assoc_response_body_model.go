// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpsPrivateAssocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateIpsPrivateAssocResponseBody
	GetRequestId() *string
}

type CreateIpsPrivateAssocResponseBody struct {
	// example:
	//
	// 99A65AA0-C5B5-5092-BFCF-8111B436****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpsPrivateAssocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpsPrivateAssocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpsPrivateAssocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpsPrivateAssocResponseBody) SetRequestId(v string) *CreateIpsPrivateAssocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpsPrivateAssocResponseBody) Validate() error {
	return dara.Validate(s)
}
