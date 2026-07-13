// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBdrcServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenBdrcServiceResponseBody
	GetRequestId() *string
}

type OpenBdrcServiceResponseBody struct {
	// The unique identifier of the request.
	//
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenBdrcServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenBdrcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBdrcServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenBdrcServiceResponseBody) SetRequestId(v string) *OpenBdrcServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenBdrcServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
