// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHybridProxyResponseBody
	GetRequestId() *string
}

type DeleteHybridProxyResponseBody struct {
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHybridProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridProxyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHybridProxyResponseBody) SetRequestId(v string) *DeleteHybridProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
