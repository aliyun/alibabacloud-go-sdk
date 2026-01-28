// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iARMSQueryDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ARMSQueryDataSetResponseBody
	GetData() *string
	SetRequestId(v string) *ARMSQueryDataSetResponseBody
	GetRequestId() *string
}

type ARMSQueryDataSetResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ARMSQueryDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ARMSQueryDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetResponseBody) GetData() *string {
	return s.Data
}

func (s *ARMSQueryDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ARMSQueryDataSetResponseBody) SetData(v string) *ARMSQueryDataSetResponseBody {
	s.Data = &v
	return s
}

func (s *ARMSQueryDataSetResponseBody) SetRequestId(v string) *ARMSQueryDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ARMSQueryDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
