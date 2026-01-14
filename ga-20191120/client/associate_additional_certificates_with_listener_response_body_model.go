// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAdditionalCertificatesWithListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody
	GetListenerId() *string
	SetRequestId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody
	GetRequestId() *string
}

type AssociateAdditionalCertificatesWithListenerResponseBody struct {
	// The listener ID.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetRequestId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
