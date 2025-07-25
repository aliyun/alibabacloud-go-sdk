// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePdnsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumePdnsServiceResponseBody
	GetRequestId() *string
}

type ResumePdnsServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumePdnsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumePdnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumePdnsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumePdnsServiceResponseBody) SetRequestId(v string) *ResumePdnsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumePdnsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
