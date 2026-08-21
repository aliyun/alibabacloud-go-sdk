// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePrivateAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeletePrivateAccessApplicationResponseBody
	GetRequestId() *string
}

type BatchDeletePrivateAccessApplicationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 170A9DD6-DECA-5E8F-8B0F-4C3B80C0644A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeletePrivateAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeletePrivateAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeletePrivateAccessApplicationResponseBody) SetRequestId(v string) *BatchDeletePrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeletePrivateAccessApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
