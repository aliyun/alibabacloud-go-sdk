// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindHybridProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindHybridProxyResponseBody
	GetRequestId() *string
}

type BindHybridProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11C96623-E106-59C9-866D-A6C82911459F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindHybridProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindHybridProxyResponseBody) GoString() string {
	return s.String()
}

func (s *BindHybridProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindHybridProxyResponseBody) SetRequestId(v string) *BindHybridProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindHybridProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
