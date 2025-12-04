// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOSSStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOSSStorageResponseBody
	GetRequestId() *string
}

type CreateOSSStorageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1F488A93-83FD-540F-9B67-0333AF64E6A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOSSStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOSSStorageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOSSStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOSSStorageResponseBody) SetRequestId(v string) *CreateOSSStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOSSStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
