// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateResourceFileResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateResourceFileResponseBody
	GetRequestId() *string
}

type CreateResourceFileResponseBody struct {
	// The ID of the file that is created.
	//
	// example:
	//
	// 1000001
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceFileResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateResourceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceFileResponseBody) SetData(v int64) *CreateResourceFileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateResourceFileResponseBody) SetRequestId(v string) *CreateResourceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceFileResponseBody) Validate() error {
	return dara.Validate(s)
}
