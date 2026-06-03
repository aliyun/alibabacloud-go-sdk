// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySdkVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySdkVersionResponseBody
	GetCode() *string
	SetData(v string) *QuerySdkVersionResponseBody
	GetData() *string
	SetRequestId(v string) *QuerySdkVersionResponseBody
	GetRequestId() *string
}

type QuerySdkVersionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySdkVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySdkVersionResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySdkVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySdkVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *QuerySdkVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySdkVersionResponseBody) SetCode(v string) *QuerySdkVersionResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySdkVersionResponseBody) SetData(v string) *QuerySdkVersionResponseBody {
	s.Data = &v
	return s
}

func (s *QuerySdkVersionResponseBody) SetRequestId(v string) *QuerySdkVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySdkVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
