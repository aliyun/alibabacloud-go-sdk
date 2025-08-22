// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSharedAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveSharedAccountsResponseBody
	GetRequestId() *string
}

type RemoveSharedAccountsResponseBody struct {
	// example:
	//
	// 8294F4ED-8DBA-5441-B3F2-61C3C5374990
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveSharedAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSharedAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSharedAccountsResponseBody) SetRequestId(v string) *RemoveSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSharedAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
