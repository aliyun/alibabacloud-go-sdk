// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchDomainRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchDomainRemarkResponseBody
	GetRequestId() *string
}

type SaveBatchDomainRemarkResponseBody struct {
	// example:
	//
	// 4189E320-961E-4786-8E15-0000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveBatchDomainRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchDomainRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchDomainRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchDomainRemarkResponseBody) SetRequestId(v string) *SaveBatchDomainRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchDomainRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
