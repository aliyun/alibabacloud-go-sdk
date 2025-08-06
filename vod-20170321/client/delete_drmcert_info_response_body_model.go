// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDRMCertInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDRMCertInfoResponseBody
	GetRequestId() *string
}

type DeleteDRMCertInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDRMCertInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDRMCertInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDRMCertInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDRMCertInfoResponseBody) SetRequestId(v string) *DeleteDRMCertInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDRMCertInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
