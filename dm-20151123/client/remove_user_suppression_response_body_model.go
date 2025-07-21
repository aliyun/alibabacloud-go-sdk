// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserSuppressionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserSuppressionResponseBody
	GetRequestId() *string
}

type RemoveUserSuppressionResponseBody struct {
	// example:
	//
	// 1A846D66-5EC7-551B-9687-5BF1963DCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserSuppressionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserSuppressionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserSuppressionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserSuppressionResponseBody) SetRequestId(v string) *RemoveUserSuppressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserSuppressionResponseBody) Validate() error {
	return dara.Validate(s)
}
