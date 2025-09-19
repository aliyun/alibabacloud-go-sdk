// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindHybridProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnBindHybridProxyResponseBody
	GetRequestId() *string
}

type UnBindHybridProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB483977
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnBindHybridProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnBindHybridProxyResponseBody) GoString() string {
	return s.String()
}

func (s *UnBindHybridProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnBindHybridProxyResponseBody) SetRequestId(v string) *UnBindHybridProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnBindHybridProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
