// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAdditionalCertificatesFromListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody
	GetJobId() *string
	SetRequestId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody
	GetRequestId() *string
}

type DissociateAdditionalCertificatesFromListenerResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetJobId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetRequestId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
