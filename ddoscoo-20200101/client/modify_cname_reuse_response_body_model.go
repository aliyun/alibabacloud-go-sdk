// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCnameReuseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCnameReuseResponseBody
	GetRequestId() *string
}

type ModifyCnameReuseResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCnameReuseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCnameReuseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCnameReuseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCnameReuseResponseBody) SetRequestId(v string) *ModifyCnameReuseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCnameReuseResponseBody) Validate() error {
	return dara.Validate(s)
}
