// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteComponentAssetResponseBody
	GetRequestId() *string
}

type DeleteComponentAssetResponseBody struct {
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// 39C38A34-8532-5D44-B88A-7263B435C316
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteComponentAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComponentAssetResponseBody) SetRequestId(v string) *DeleteComponentAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComponentAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
