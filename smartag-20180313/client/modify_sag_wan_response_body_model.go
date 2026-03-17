// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagWanResponseBody
	GetRequestId() *string
}

type ModifySagWanResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AFF7E5A6-6897-4FDC-A5A8-1978B5B3E545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagWanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagWanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagWanResponseBody) SetRequestId(v string) *ModifySagWanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagWanResponseBody) Validate() error {
	return dara.Validate(s)
}
