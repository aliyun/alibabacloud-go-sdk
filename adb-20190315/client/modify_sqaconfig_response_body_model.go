// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQAConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySQAConfigResponseBody
	GetRequestId() *string
}

type ModifySQAConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 494486CE-6F49-574E-B304-29127EA12E36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySQAConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySQAConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySQAConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySQAConfigResponseBody) SetRequestId(v string) *ModifySQAConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySQAConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
