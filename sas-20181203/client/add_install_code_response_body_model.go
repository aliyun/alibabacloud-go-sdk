// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstallCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddInstallCodeResponseBody
	GetRequestId() *string
}

type AddInstallCodeResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddInstallCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddInstallCodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddInstallCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddInstallCodeResponseBody) SetRequestId(v string) *AddInstallCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddInstallCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
