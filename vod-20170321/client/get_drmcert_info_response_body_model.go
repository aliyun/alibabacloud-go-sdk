// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMCertInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDRMCertInfo(v string) *GetDRMCertInfoResponseBody
	GetDRMCertInfo() *string
	SetRequestId(v string) *GetDRMCertInfoResponseBody
	GetRequestId() *string
}

type GetDRMCertInfoResponseBody struct {
	DRMCertInfo *string `json:"DRMCertInfo,omitempty" xml:"DRMCertInfo,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDRMCertInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDRMCertInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDRMCertInfoResponseBody) GetDRMCertInfo() *string {
	return s.DRMCertInfo
}

func (s *GetDRMCertInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDRMCertInfoResponseBody) SetDRMCertInfo(v string) *GetDRMCertInfoResponseBody {
	s.DRMCertInfo = &v
	return s
}

func (s *GetDRMCertInfoResponseBody) SetRequestId(v string) *GetDRMCertInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDRMCertInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
