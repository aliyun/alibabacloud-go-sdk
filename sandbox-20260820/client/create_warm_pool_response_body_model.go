// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarmPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWarmPoolResponseBody
	GetCode() *string
	SetMessage(v string) *CreateWarmPoolResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWarmPoolResponseBody
	GetRequestId() *string
	SetWarmPool(v *PublicWarmPoolConfig) *CreateWarmPoolResponseBody
	GetWarmPool() *PublicWarmPoolConfig
}

type CreateWarmPoolResponseBody struct {
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

func (s CreateWarmPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWarmPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarmPoolResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWarmPoolResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWarmPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWarmPoolResponseBody) GetWarmPool() *PublicWarmPoolConfig {
	return s.WarmPool
}

func (s *CreateWarmPoolResponseBody) SetCode(v string) *CreateWarmPoolResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWarmPoolResponseBody) SetMessage(v string) *CreateWarmPoolResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWarmPoolResponseBody) SetRequestId(v string) *CreateWarmPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarmPoolResponseBody) SetWarmPool(v *PublicWarmPoolConfig) *CreateWarmPoolResponseBody {
	s.WarmPool = v
	return s
}

func (s *CreateWarmPoolResponseBody) Validate() error {
	if s.WarmPool != nil {
		if err := s.WarmPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}
