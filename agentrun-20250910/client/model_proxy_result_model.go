// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelProxyResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelProxyResult
	GetCode() *string
	SetData(v *ModelProxy) *ModelProxyResult
	GetData() *ModelProxy
	SetRequestId(v string) *ModelProxyResult
	GetRequestId() *string
}

type ModelProxyResult struct {
	Code      *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ModelProxy `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ModelProxyResult) String() string {
	return dara.Prettify(s)
}

func (s ModelProxyResult) GoString() string {
	return s.String()
}

func (s *ModelProxyResult) GetCode() *string {
	return s.Code
}

func (s *ModelProxyResult) GetData() *ModelProxy {
	return s.Data
}

func (s *ModelProxyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelProxyResult) SetCode(v string) *ModelProxyResult {
	s.Code = &v
	return s
}

func (s *ModelProxyResult) SetData(v *ModelProxy) *ModelProxyResult {
	s.Data = v
	return s
}

func (s *ModelProxyResult) SetRequestId(v string) *ModelProxyResult {
	s.RequestId = &v
	return s
}

func (s *ModelProxyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
