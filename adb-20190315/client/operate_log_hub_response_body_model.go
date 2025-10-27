// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateLogHubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateLogHubResponseBody
	GetRequestId() *string
}

type OperateLogHubResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 71B41FF9-1275-5F75-973D-8BC3C60236E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateLogHubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateLogHubResponseBody) GoString() string {
	return s.String()
}

func (s *OperateLogHubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateLogHubResponseBody) SetRequestId(v string) *OperateLogHubResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateLogHubResponseBody) Validate() error {
	return dara.Validate(s)
}
