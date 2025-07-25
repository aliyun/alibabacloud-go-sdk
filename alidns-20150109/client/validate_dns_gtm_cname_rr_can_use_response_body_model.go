// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDnsGtmCnameRrCanUseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateDnsGtmCnameRrCanUseResponseBody
	GetRequestId() *string
}

type ValidateDnsGtmCnameRrCanUseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateDnsGtmCnameRrCanUseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateDnsGtmCnameRrCanUseResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateDnsGtmCnameRrCanUseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateDnsGtmCnameRrCanUseResponseBody) SetRequestId(v string) *ValidateDnsGtmCnameRrCanUseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseResponseBody) Validate() error {
	return dara.Validate(s)
}
