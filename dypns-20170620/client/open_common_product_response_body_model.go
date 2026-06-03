// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCommonProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenCommonProductResponseBody
	GetRequestId() *string
	SetCode(v string) *OpenCommonProductResponseBody
	GetCode() *string
	SetData(v string) *OpenCommonProductResponseBody
	GetData() *string
}

type OpenCommonProductResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OpenCommonProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenCommonProductResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCommonProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenCommonProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenCommonProductResponseBody) GetData() *string {
	return s.Data
}

func (s *OpenCommonProductResponseBody) SetRequestId(v string) *OpenCommonProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCommonProductResponseBody) SetCode(v string) *OpenCommonProductResponseBody {
	s.Code = &v
	return s
}

func (s *OpenCommonProductResponseBody) SetData(v string) *OpenCommonProductResponseBody {
	s.Data = &v
	return s
}

func (s *OpenCommonProductResponseBody) Validate() error {
	return dara.Validate(s)
}
