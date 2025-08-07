// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterSSLResponseBody
	GetRequestId() *string
}

type ModifyDBClusterSSLResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A94B1755-6D8B-4E27-BF3C-8562BC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterSSLResponseBody) SetRequestId(v string) *ModifyDBClusterSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
