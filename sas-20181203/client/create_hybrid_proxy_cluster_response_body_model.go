// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridProxyClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateHybridProxyClusterResponseBody
	GetRequestId() *string
}

type CreateHybridProxyClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CDCB0BBB-CFB2-5D38-BB49-500E2A21xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHybridProxyClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridProxyClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridProxyClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHybridProxyClusterResponseBody) SetRequestId(v string) *CreateHybridProxyClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridProxyClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
