// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonStatisticPreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCommonStatisticPreviewResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryCommonStatisticPreviewResponseBody
	GetCode() *string
	SetData(v string) *QueryCommonStatisticPreviewResponseBody
	GetData() *string
}

type QueryCommonStatisticPreviewResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryCommonStatisticPreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonStatisticPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommonStatisticPreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCommonStatisticPreviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCommonStatisticPreviewResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryCommonStatisticPreviewResponseBody) SetRequestId(v string) *QueryCommonStatisticPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCommonStatisticPreviewResponseBody) SetCode(v string) *QueryCommonStatisticPreviewResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCommonStatisticPreviewResponseBody) SetData(v string) *QueryCommonStatisticPreviewResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCommonStatisticPreviewResponseBody) Validate() error {
	return dara.Validate(s)
}
