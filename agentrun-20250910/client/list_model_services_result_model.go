// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelServicesResult
	GetCode() *string
	SetData(v *ListModelServicesOutput) *ListModelServicesResult
	GetData() *ListModelServicesOutput
	SetRequestId(v string) *ListModelServicesResult
	GetRequestId() *string
}

type ListModelServicesResult struct {
	Code      *string                  `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListModelServicesOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelServicesResult) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesResult) GoString() string {
	return s.String()
}

func (s *ListModelServicesResult) GetCode() *string {
	return s.Code
}

func (s *ListModelServicesResult) GetData() *ListModelServicesOutput {
	return s.Data
}

func (s *ListModelServicesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelServicesResult) SetCode(v string) *ListModelServicesResult {
	s.Code = &v
	return s
}

func (s *ListModelServicesResult) SetData(v *ListModelServicesOutput) *ListModelServicesResult {
	s.Data = v
	return s
}

func (s *ListModelServicesResult) SetRequestId(v string) *ListModelServicesResult {
	s.RequestId = &v
	return s
}

func (s *ListModelServicesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
