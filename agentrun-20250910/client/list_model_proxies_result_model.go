// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProxiesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelProxiesResult
	GetCode() *string
	SetData(v *ListModelProxiesOutput) *ListModelProxiesResult
	GetData() *ListModelProxiesOutput
	SetRequestId(v string) *ListModelProxiesResult
	GetRequestId() *string
}

type ListModelProxiesResult struct {
	Code      *string                 `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListModelProxiesOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelProxiesResult) String() string {
	return dara.Prettify(s)
}

func (s ListModelProxiesResult) GoString() string {
	return s.String()
}

func (s *ListModelProxiesResult) GetCode() *string {
	return s.Code
}

func (s *ListModelProxiesResult) GetData() *ListModelProxiesOutput {
	return s.Data
}

func (s *ListModelProxiesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelProxiesResult) SetCode(v string) *ListModelProxiesResult {
	s.Code = &v
	return s
}

func (s *ListModelProxiesResult) SetData(v *ListModelProxiesOutput) *ListModelProxiesResult {
	s.Data = v
	return s
}

func (s *ListModelProxiesResult) SetRequestId(v string) *ListModelProxiesResult {
	s.RequestId = &v
	return s
}

func (s *ListModelProxiesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
