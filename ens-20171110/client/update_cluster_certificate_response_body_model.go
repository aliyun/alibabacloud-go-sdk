// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateClusterCertificateResponseBody
	GetClusterId() *string
	SetMessage(v string) *UpdateClusterCertificateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateClusterCertificateResponseBody
	GetRequestId() *string
}

type UpdateClusterCertificateResponseBody struct {
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClusterCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterCertificateResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterCertificateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateClusterCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterCertificateResponseBody) SetClusterId(v string) *UpdateClusterCertificateResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterCertificateResponseBody) SetMessage(v string) *UpdateClusterCertificateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClusterCertificateResponseBody) SetRequestId(v string) *UpdateClusterCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
