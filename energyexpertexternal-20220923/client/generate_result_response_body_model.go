// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *GenerateResultResponseBody
	GetData() *bool
	SetRequestId(v string) *GenerateResultResponseBody
	GetRequestId() *string
}

type GenerateResultResponseBody struct {
	// The returned data. `true` indicates that the request is successful, `false` indicates that the request fails.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateResultResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateResultResponseBody) GetData() *bool {
	return s.Data
}

func (s *GenerateResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateResultResponseBody) SetData(v bool) *GenerateResultResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateResultResponseBody) SetRequestId(v string) *GenerateResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateResultResponseBody) Validate() error {
	return dara.Validate(s)
}
