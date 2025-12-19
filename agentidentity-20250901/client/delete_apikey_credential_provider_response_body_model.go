// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAPIKeyCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAPIKeyCredentialProviderResponseBody
	GetRequestId() *string
}

type DeleteAPIKeyCredentialProviderResponseBody struct {
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAPIKeyCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAPIKeyCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAPIKeyCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAPIKeyCredentialProviderResponseBody) SetRequestId(v string) *DeleteAPIKeyCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAPIKeyCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
