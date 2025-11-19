// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelServiceResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelServiceResult
	GetCode() *string
	SetData(v *ModelService) *ModelServiceResult
	GetData() *ModelService
	SetRequestId(v string) *ModelServiceResult
	GetRequestId() *string
}

type ModelServiceResult struct {
	Code      *string       `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ModelService `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ModelServiceResult) String() string {
	return dara.Prettify(s)
}

func (s ModelServiceResult) GoString() string {
	return s.String()
}

func (s *ModelServiceResult) GetCode() *string {
	return s.Code
}

func (s *ModelServiceResult) GetData() *ModelService {
	return s.Data
}

func (s *ModelServiceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelServiceResult) SetCode(v string) *ModelServiceResult {
	s.Code = &v
	return s
}

func (s *ModelServiceResult) SetData(v *ModelService) *ModelServiceResult {
	s.Data = v
	return s
}

func (s *ModelServiceResult) SetRequestId(v string) *ModelServiceResult {
	s.RequestId = &v
	return s
}

func (s *ModelServiceResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
