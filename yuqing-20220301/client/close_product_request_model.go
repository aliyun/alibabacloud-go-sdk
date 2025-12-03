// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductInstance(v *ProductInstance) *CloseProductRequest
	GetProductInstance() *ProductInstance
	SetRequestId(v string) *CloseProductRequest
	GetRequestId() *string
}

type CloseProductRequest struct {
	ProductInstance *ProductInstance `json:"productInstance,omitempty" xml:"productInstance,omitempty"`
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CloseProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseProductRequest) GoString() string {
	return s.String()
}

func (s *CloseProductRequest) GetProductInstance() *ProductInstance {
	return s.ProductInstance
}

func (s *CloseProductRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseProductRequest) SetProductInstance(v *ProductInstance) *CloseProductRequest {
	s.ProductInstance = v
	return s
}

func (s *CloseProductRequest) SetRequestId(v string) *CloseProductRequest {
	s.RequestId = &v
	return s
}

func (s *CloseProductRequest) Validate() error {
	if s.ProductInstance != nil {
		if err := s.ProductInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}
