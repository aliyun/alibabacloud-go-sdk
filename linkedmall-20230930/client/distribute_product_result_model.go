// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributeProductResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DistributeProductResult
	GetCode() *string
	SetMessage(v string) *DistributeProductResult
	GetMessage() *string
	SetRequestId(v string) *DistributeProductResult
	GetRequestId() *string
}

type DistributeProductResult struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DistributeProductResult) String() string {
	return dara.Prettify(s)
}

func (s DistributeProductResult) GoString() string {
	return s.String()
}

func (s *DistributeProductResult) GetCode() *string {
	return s.Code
}

func (s *DistributeProductResult) GetMessage() *string {
	return s.Message
}

func (s *DistributeProductResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DistributeProductResult) SetCode(v string) *DistributeProductResult {
	s.Code = &v
	return s
}

func (s *DistributeProductResult) SetMessage(v string) *DistributeProductResult {
	s.Message = &v
	return s
}

func (s *DistributeProductResult) SetRequestId(v string) *DistributeProductResult {
	s.RequestId = &v
	return s
}

func (s *DistributeProductResult) Validate() error {
	return dara.Validate(s)
}
