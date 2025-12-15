// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSharedStoragesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AttachSharedStoragesResponseBody
	GetClusterId() *string
	SetRequestId(v string) *AttachSharedStoragesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AttachSharedStoragesResponseBody
	GetSuccess() *string
}

type AttachSharedStoragesResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F9B7BEF8-E42E-5090-9880-55FB7872****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachSharedStoragesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachSharedStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *AttachSharedStoragesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachSharedStoragesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AttachSharedStoragesResponseBody) SetClusterId(v string) *AttachSharedStoragesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *AttachSharedStoragesResponseBody) SetRequestId(v string) *AttachSharedStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachSharedStoragesResponseBody) SetSuccess(v string) *AttachSharedStoragesResponseBody {
	s.Success = &v
	return s
}

func (s *AttachSharedStoragesResponseBody) Validate() error {
	return dara.Validate(s)
}
