// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteModelServiceResult
	GetCode() *string
	SetData(v *ModelService) *DeleteModelServiceResult
	GetData() *ModelService
	SetRequestId(v string) *DeleteModelServiceResult
	GetRequestId() *string
}

type DeleteModelServiceResult struct {
	Code      *string       `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ModelService `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteModelServiceResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceResult) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceResult) GetCode() *string {
	return s.Code
}

func (s *DeleteModelServiceResult) GetData() *ModelService {
	return s.Data
}

func (s *DeleteModelServiceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelServiceResult) SetCode(v string) *DeleteModelServiceResult {
	s.Code = &v
	return s
}

func (s *DeleteModelServiceResult) SetData(v *ModelService) *DeleteModelServiceResult {
	s.Data = v
	return s
}

func (s *DeleteModelServiceResult) SetRequestId(v string) *DeleteModelServiceResult {
	s.RequestId = &v
	return s
}

func (s *DeleteModelServiceResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
