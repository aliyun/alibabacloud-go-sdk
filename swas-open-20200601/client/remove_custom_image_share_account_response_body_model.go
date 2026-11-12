// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCustomImageShareAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveCustomImageShareAccountResponseBody
	GetRequestId() *string
}

type RemoveCustomImageShareAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCustomImageShareAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveCustomImageShareAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCustomImageShareAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveCustomImageShareAccountResponseBody) SetRequestId(v string) *RemoveCustomImageShareAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveCustomImageShareAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
