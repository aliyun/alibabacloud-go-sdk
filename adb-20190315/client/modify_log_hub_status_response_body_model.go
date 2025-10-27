// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogHubStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLogHubStatusResponseBody
	GetRequestId() *string
}

type ModifyLogHubStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AA1F421-395B-5BC4-BDDC-762C508A952B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogHubStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogHubStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogHubStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLogHubStatusResponseBody) SetRequestId(v string) *ModifyLogHubStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLogHubStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
