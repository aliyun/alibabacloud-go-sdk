// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarmPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWarmPoolResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateWarmPoolResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWarmPoolResponseBody
	GetRequestId() *string
	SetWarmPool(v *PublicWarmPoolConfig) *UpdateWarmPoolResponseBody
	GetWarmPool() *PublicWarmPoolConfig
}

type UpdateWarmPoolResponseBody struct {
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
	WarmPool *PublicWarmPoolConfig `json:"warmPool,omitempty" xml:"warmPool,omitempty"`
}

func (s UpdateWarmPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarmPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarmPoolResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWarmPoolResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWarmPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWarmPoolResponseBody) GetWarmPool() *PublicWarmPoolConfig {
	return s.WarmPool
}

func (s *UpdateWarmPoolResponseBody) SetCode(v string) *UpdateWarmPoolResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWarmPoolResponseBody) SetMessage(v string) *UpdateWarmPoolResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWarmPoolResponseBody) SetRequestId(v string) *UpdateWarmPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarmPoolResponseBody) SetWarmPool(v *PublicWarmPoolConfig) *UpdateWarmPoolResponseBody {
	s.WarmPool = v
	return s
}

func (s *UpdateWarmPoolResponseBody) Validate() error {
	if s.WarmPool != nil {
		if err := s.WarmPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}
