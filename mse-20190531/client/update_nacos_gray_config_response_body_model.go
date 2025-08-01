// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosGrayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateNacosGrayConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateNacosGrayConfigResponseBody
	GetRequestId() *string
}

type UpdateNacosGrayConfigResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNacosGrayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosGrayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacosGrayConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateNacosGrayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNacosGrayConfigResponseBody) SetData(v bool) *UpdateNacosGrayConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNacosGrayConfigResponseBody) SetRequestId(v string) *UpdateNacosGrayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacosGrayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
