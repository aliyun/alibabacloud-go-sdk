// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProxyResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteModelProxyResult
	GetCode() *string
	SetData(v *ModelProxy) *DeleteModelProxyResult
	GetData() *ModelProxy
	SetRequestId(v string) *DeleteModelProxyResult
	GetRequestId() *string
}

type DeleteModelProxyResult struct {
	Code      *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ModelProxy `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteModelProxyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProxyResult) GoString() string {
	return s.String()
}

func (s *DeleteModelProxyResult) GetCode() *string {
	return s.Code
}

func (s *DeleteModelProxyResult) GetData() *ModelProxy {
	return s.Data
}

func (s *DeleteModelProxyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelProxyResult) SetCode(v string) *DeleteModelProxyResult {
	s.Code = &v
	return s
}

func (s *DeleteModelProxyResult) SetData(v *ModelProxy) *DeleteModelProxyResult {
	s.Data = v
	return s
}

func (s *DeleteModelProxyResult) SetRequestId(v string) *DeleteModelProxyResult {
	s.RequestId = &v
	return s
}

func (s *DeleteModelProxyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
