// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCsrId(v int64) *UploadCsrResponseBody
	GetCsrId() *int64
	SetRequestId(v string) *UploadCsrResponseBody
	GetRequestId() *string
}

type UploadCsrResponseBody struct {
	// The ID of the CSR.
	//
	// example:
	//
	// 2271
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadCsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadCsrResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCsrResponseBody) GetCsrId() *int64 {
	return s.CsrId
}

func (s *UploadCsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadCsrResponseBody) SetCsrId(v int64) *UploadCsrResponseBody {
	s.CsrId = &v
	return s
}

func (s *UploadCsrResponseBody) SetRequestId(v string) *UploadCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCsrResponseBody) Validate() error {
	return dara.Validate(s)
}
