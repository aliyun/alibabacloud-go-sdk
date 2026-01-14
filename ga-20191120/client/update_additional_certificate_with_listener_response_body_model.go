// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdditionalCertificateWithListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListenerId(v string) *UpdateAdditionalCertificateWithListenerResponseBody
	GetListenerId() *string
	SetRequestId(v string) *UpdateAdditionalCertificateWithListenerResponseBody
	GetRequestId() *string
}

type UpdateAdditionalCertificateWithListenerResponseBody struct {
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAdditionalCertificateWithListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdditionalCertificateWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdditionalCertificateWithListenerResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateAdditionalCertificateWithListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAdditionalCertificateWithListenerResponseBody) SetListenerId(v string) *UpdateAdditionalCertificateWithListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponseBody) SetRequestId(v string) *UpdateAdditionalCertificateWithListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
