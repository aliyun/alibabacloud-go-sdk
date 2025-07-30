// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPcaCertToCasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadPcaCertToCasResponseBody
	GetRequestId() *string
}

type UploadPcaCertToCasResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadPcaCertToCasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadPcaCertToCasResponseBody) GoString() string {
	return s.String()
}

func (s *UploadPcaCertToCasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadPcaCertToCasResponseBody) SetRequestId(v string) *UploadPcaCertToCasResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadPcaCertToCasResponseBody) Validate() error {
	return dara.Validate(s)
}
