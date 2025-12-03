// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenProductRequest
	GetClientToken() *string
	SetProductInstance(v *ProductInstance) *OpenProductRequest
	GetProductInstance() *ProductInstance
	SetRequestId(v string) *OpenProductRequest
	GetRequestId() *string
}

type OpenProductRequest struct {
	ClientToken     *string          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ProductInstance *ProductInstance `json:"productInstance,omitempty" xml:"productInstance,omitempty"`
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OpenProductRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenProductRequest) GoString() string {
	return s.String()
}

func (s *OpenProductRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenProductRequest) GetProductInstance() *ProductInstance {
	return s.ProductInstance
}

func (s *OpenProductRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenProductRequest) SetClientToken(v string) *OpenProductRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenProductRequest) SetProductInstance(v *ProductInstance) *OpenProductRequest {
	s.ProductInstance = v
	return s
}

func (s *OpenProductRequest) SetRequestId(v string) *OpenProductRequest {
	s.RequestId = &v
	return s
}

func (s *OpenProductRequest) Validate() error {
	if s.ProductInstance != nil {
		if err := s.ProductInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}
