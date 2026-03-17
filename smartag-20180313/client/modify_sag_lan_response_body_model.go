// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagLanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagLanResponseBody
	GetRequestId() *string
}

type ModifySagLanResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1120228A-E5E1-4E9C-B56D-96887E1A2B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagLanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagLanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagLanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagLanResponseBody) SetRequestId(v string) *ModifySagLanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagLanResponseBody) Validate() error {
	return dara.Validate(s)
}
