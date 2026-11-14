// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarmPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeWarmPoolResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeWarmPoolResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeWarmPoolResponseBody
	GetRequestId() *string
	SetWarmPool(v *PublicWarmPool) *DescribeWarmPoolResponseBody
	GetWarmPool() *PublicWarmPool
}

type DescribeWarmPoolResponseBody struct {
	// The response business code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01234567-89AB-CDEF-0123-456789ABCDEF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The warm pool configuration or current status.
	WarmPool *PublicWarmPool `json:"warmPool,omitempty" xml:"warmPool,omitempty"`
}

func (s DescribeWarmPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarmPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWarmPoolResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeWarmPoolResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWarmPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWarmPoolResponseBody) GetWarmPool() *PublicWarmPool {
	return s.WarmPool
}

func (s *DescribeWarmPoolResponseBody) SetCode(v string) *DescribeWarmPoolResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWarmPoolResponseBody) SetMessage(v string) *DescribeWarmPoolResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWarmPoolResponseBody) SetRequestId(v string) *DescribeWarmPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWarmPoolResponseBody) SetWarmPool(v *PublicWarmPool) *DescribeWarmPoolResponseBody {
	s.WarmPool = v
	return s
}

func (s *DescribeWarmPoolResponseBody) Validate() error {
	if s.WarmPool != nil {
		if err := s.WarmPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}
