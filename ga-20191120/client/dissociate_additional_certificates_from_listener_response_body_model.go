// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAdditionalCertificatesFromListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody
	GetRequestId() *string
}

type DissociateAdditionalCertificatesFromListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetRequestId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
