// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteResourceGroupResponseBody
	GetSuccess() *bool
}

type DeleteResourceGroupResponseBody struct {
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetSuccess(v bool) *DeleteResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
