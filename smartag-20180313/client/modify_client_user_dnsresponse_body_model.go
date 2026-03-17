// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientUserDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClientUserDNSResponseBody
	GetRequestId() *string
}

type ModifyClientUserDNSResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BFE2D0C0-B69F-422D-A8A3-928AD511B471
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClientUserDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientUserDNSResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClientUserDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClientUserDNSResponseBody) SetRequestId(v string) *ModifyClientUserDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClientUserDNSResponseBody) Validate() error {
	return dara.Validate(s)
}
